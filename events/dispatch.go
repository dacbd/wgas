package events

import (
	"fmt"
	"github.com/google/go-github/v62/github"
	"github.com/rs/zerolog/log"
)

func HandleWorkflowDispatch(event github.Workflow) error {
	log.Info().Msg(fmt.Sprintf("WorkflowDispatch: %s", event.GetName()))
	return nil
}
