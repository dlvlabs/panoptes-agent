package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "panoptes",
	Short:                 "Panoptes Agent - Cosmos node monitoring agent",
	DisableAutoGenTag:     true,
	DisableFlagParsing:    false,
	DisableFlagsInUseLine: true,
	DisableSuggestions:    true,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func ExecuteCLI() error {
	return rootCmd.Execute()
}
