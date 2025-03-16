package cli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"dlvlabs.net/panoptes-agent/config"
	"dlvlabs.net/panoptes-agent/internal/agent"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run the Panoptes Agent",
	RunE:  runStart,
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func runStart(cmd *cobra.Command, args []string) error {
	cfg, err := config.LoadConfig("config/config.toml")
	if err != nil {
		return fmt.Errorf("configuration file load failed: %v", err)
	}

	monitor := agent.NewAgent(cfg)
	if err := monitor.Start(); err != nil {
		return fmt.Errorf("monitoring start failed: %v", err)
	}
	defer monitor.Stop()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	return nil
}
