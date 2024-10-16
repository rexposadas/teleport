package server

import (
	"fmt"
	"sync"

	"github.com/google/uuid"

	pb "github.com/rexposadas/teleport/api"
)

// JobManager manages the jobs.
type JobManager struct {
	mutex sync.RWMutex
	jobs  map[uuid.UUID]*Job
}

func NewJobManager() *JobManager {

	return &JobManager{jobs: map[uuid.UUID]*Job{}}
}

func (jm *JobManager) StartJob(cmd string, args []string) (*Job, error) {
	j := NewJob(cmd, args)

	if err := j.Start(); err != nil {
		return nil, fmt.Errorf("job start: %w", err)
	}

	jm.mutex.Lock()
	jm.jobs[j.ID] = j
	jm.mutex.Unlock()

	return j, nil
}

func (jm *JobManager) JobStatus(id uuid.UUID) (pb.Status, error) {
	job, err := jm.GetJob(id)
	if err != nil {
		return pb.Status_STATUS_UNKNOWN, fmt.Errorf("job status: %w", err)
	}

	return job.status, nil
}

func (jm *JobManager) GetJob(id uuid.UUID) (*Job, error) {
	jm.mutex.RLock()
	j := jm.jobs[id]
	jm.mutex.RUnlock()

	if j == nil {
		return nil, fmt.Errorf("job not found %s", j.ID)
	}

	return j, nil
}
