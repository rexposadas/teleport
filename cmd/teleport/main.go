package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "teleport",
		Short: "Teleport CLI application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Teleport CLI!")
		},
	}

	rootCmd.AddCommand(StartCmd())
	rootCmd.AddCommand(StatusCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
