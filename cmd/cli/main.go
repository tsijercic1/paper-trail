package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "paper-trail",
		Short: "Paper Trail is a CLI application",
		Long:  `Paper Trail is a CLI application to manage and track your documents.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Paper Trail CLI!")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Paper Trail",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Paper Trail CLI v0.1.0")
		},
	}

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
