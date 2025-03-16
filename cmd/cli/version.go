package cli

import (
  "fmt"

  "dlvlabs.net/panoptes-agent/config"
  "github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print Panoptes Agent version information",
  RunE:  runVersion,
}

func init() {
  rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) error {
  fmt.Println("Panoptes Agent version:", config.Version)
  return nil
}
