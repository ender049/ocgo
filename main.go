package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "opencodeui",
		Short: "OpenCodeUI Frontend Server",
		Long: `Static file server with API proxy to opencode backend.

Examples:
  opencodeui start                  # Start server (daemon)
  opencodeui status                # Check server status
  opencodeui stop                  # Stop server
  opencodeui update                # Update tool/frontend
  opencodeui version               # Show versions
`,
	}

	rootCmd.AddCommand(startCmd, stopCmd, statusCmd, versionCmd, updateCmd, serveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
