package events

import (
	"fmt"
	"github.com/google/go-github/v62/github"
	"github.com/rs/zerolog/log"
)

func HandleWorkflowJob(event *github.WorkflowJob) error {
	log.Info().Msg(fmt.Sprintf("WorkflowJob, status: %s", *event.Status))
	return nil
}
