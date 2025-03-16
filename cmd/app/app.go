package app

import (
  "os"
  "os/signal"
  "syscall"

  "dlvlabs.net/panoptes-agent/config"
  "dlvlabs.net/panoptes-agent/internal/agent"
)

func Execute() error {
  cfg, err := config.LoadConfig("config/config.toml")
  if err != nil {
    return err
  }

  monitor := agent.NewAgent(cfg)
  if err := monitor.Start(); err != nil {
    return err
  }
  defer monitor.Stop()

  sigCh := make(chan os.Signal, 1)
  signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
  <-sigCh

  return nil
}
