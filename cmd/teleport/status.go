package main

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/rexposadas/teleport/pkg/client"
)

func StatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status [job_id]",
		Short: "get status of a job",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := client.NewClient()
			defer c.Close()

			jobID := args[0]

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			r, err := c.GetStatus(ctx, jobID)
			if err != nil {

				fmt.Printf("failed to get status: %v", err)
			}
			fmt.Printf("job status: %s", r.GetStatus())
		},
	}
}
