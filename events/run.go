package events

import (
	"fmt"
	"github.com/google/go-github/v62/github"
	"github.com/rs/zerolog/log"
)

func HandleWorkflowRun(event *github.WorkflowRun) error {
	log.Info().Msg(fmt.Sprintf("WorkflowRun, status: %s", *event.Status))
	return nil
}
