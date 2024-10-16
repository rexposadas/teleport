package server

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	pb "github.com/rexposadas/teleport/api"
	"log"
)

type Server struct {
	pb.UnimplementedTeleportServiceServer

	jobmanager *JobManager
}

func NewServer() *Server {
	return &Server{
		jobmanager: NewJobManager(),
	}
}

func (s *Server) Start(_ context.Context, in *pb.StartRequest) (*pb.StartResponse, error) {
	job, err := s.jobmanager.StartJob(in.Command, in.Args)
	if err != nil {
		return nil, err
	}

	go func() {
		err := job.Cmd.Wait()
		if err != nil {
			log.Printf("Process %d failed: %v", job.Cmd.Process.Pid, err)
		} else {
			job.SetStatus(pb.Status_STATUS_EXITED)
			log.Printf("Process %d finished successfully", job.Cmd.Process.Pid)
		}
	}()

	return &pb.StartResponse{JobId: job.ID.String()}, nil
}

func (s *Server) GetStatus(_ context.Context, in *pb.GetStatusRequest) (*pb.GetStatusResponse, error) {
	id, err := uuid.Parse(in.JobId)
	if err != nil {
		return nil, fmt.Errorf("invalid job id %s", in.JobId)
	}

	status, err := s.jobmanager.JobStatus(id)
	if err != nil {
		return nil, err
	}

	return &pb.GetStatusResponse{Status: status}, nil
}
