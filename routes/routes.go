package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/go-github/v62/github"
	"github.com/rs/zerolog/log"

	"github.com/dacbd/wgas/events"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Info().Dur("response_time", time.Since(start)).Str("method", r.Method).Str("path", r.URL.Path).Msg("")
	})
}

func AddRoutes(mux *http.ServeMux) {
	mux.Handle("GET /status", logging(statusCheck()))
	mux.Handle("GET /webhook", logging(handleGithubEvent([]byte("temp"))))
}

func statusCheck() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok\n"))
	})
}

func handleGithubEvent(githubSecret []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payload, err := github.ValidatePayload(r, githubSecret)
		if err != nil {
			log.Warn().Err(err).Msg("failed to validate github payload")
			w.WriteHeader(http.StatusBadRequest)
		}
		githubEvent := github.WebHookType(r)
		event, err := github.ParseWebHook(githubEvent, payload)
		if err != nil {
			log.Err(err).Msg("failed to Parse github payload")
			w.WriteHeader(http.StatusBadRequest)
		}

		switch e := event.(type) {
		case *github.WorkflowJobEvent:
			err := events.HandleWorkflowJob(e.WorkflowJob)
			if err != nil {
				log.Err(err).Msg("failed processing job event")
			}
			w.WriteHeader(http.StatusOK)
		case *github.WorkflowRunEvent:
			err := events.HandleWorkflowRun(e.WorkflowRun)
			if err != nil {
				log.Err(err).Msg("failed processing run event")
			}
			w.WriteHeader(http.StatusOK)
		case *github.WorkflowDispatchEvent:
			log.Info().Msg(fmt.Sprintf("WorkflowDispatchEvent: %s", *e.Workflow))
			w.WriteHeader(http.StatusOK)
		default:
			log.Info().Msg("")
		}
	})
}
