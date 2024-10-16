package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/rexposadas/teleport/pkg/client"
	"github.com/spf13/cobra"
)

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start [command] [args...]",
		Short: "Start a process",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := client.NewClient()
			defer c.Close()

			command := args[0]
			commandArgs := args[1:]

			// Contact the server and print out its response.
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := c.Start(ctx, command, commandArgs)
			if err != nil {
				slog.Error("could not start process: %v", err)
			}

			fmt.Printf("job id: %s", r.JobId)
		},
	}
}
