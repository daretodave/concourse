package commands

import (
	"fmt"

	"github.com/concourse/concourse/fly/commands/internal/flaghelpers"
	"github.com/concourse/concourse/fly/rc"
)

type PauseJobCommand struct {
	Job  flaghelpers.JobFlag `short:"j" long:"job" required:"true" value-name:"PIPELINE/JOB" description:"Name of a job to pause"`
	Team string              `long:"team" description:"Name of the team to which the job belongs, if different from the target default"`
}

func (command *PauseJobCommand) Execute(args []string) error {
	pipelineName, jobName := command.Job.PipelineName, command.Job.JobName
	target, err := rc.LoadTarget(Fly.Target, Fly.Verbose)
	if err != nil {
		return err
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	team := GetTeam(target, command.Team)
	found, err := team.PauseJob(pipelineName, jobName)
	if err != nil {
		return err
	}

	if !found {
		return fmt.Errorf("%s/%s not found on team %s\n", pipelineName, jobName, team.Name())
	}

	fmt.Printf("paused '%s'\n", jobName)

	return nil
}
