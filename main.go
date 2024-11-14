package main

import (
  "log"

  "dlvlabs.net/panoptes-agent/config"
)

func main() {
  cfg, err := config.LoadConfig("config/config.toml")

  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }

  log.Printf("Version: %s", config.Version)
  log.Printf("Node Name: %s", cfg.Agent.Name)
  log.Printf("Disk space feature Enabled: %v", cfg.Feature.DiskSpace)
}
