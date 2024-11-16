package main

import (
  "log"
  "os"
  "os/signal"
  "syscall"

  "dlvlabs.net/panoptes-agent/config"
  "dlvlabs.net/panoptes-agent/internal/agent"
)

func main() {
  cfg, err := config.LoadConfig("config/config.toml")
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }

  monitor := agent.NewAgent(cfg)
  if err := monitor.Start(); err != nil {
    log.Fatalf("Failed to start monitoring: %v", err)
  }
  defer monitor.Stop()

  sigCh := make(chan os.Signal, 1)
  signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
  <-sigCh
}
