package server

import (
	"fmt"
	"github.com/google/uuid"
	"os/exec"

	pb "github.com/rexposadas/teleport/api"
)

// Job represent an individual job.
type Job struct {
	ID uuid.UUID

	// to get PID cmd.Process.Pid
	Cmd *exec.Cmd

	status pb.Status
}

func NewJob(cmd string, args []string) *Job {
	j := &Job{
		ID:  uuid.New(),
		Cmd: exec.Command(cmd, args...),
	}

	return j
}

func (j *Job) Start() error {
	if err := j.Cmd.Start(); err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}
	j.SetStatus(pb.Status_STATUS_RUNNING)

	return nil
}

func (j *Job) SetStatus(status pb.Status) {
	j.status = status
}
