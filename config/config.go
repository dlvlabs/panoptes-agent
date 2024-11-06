package config

import (
  "fmt"
  "github.com/BurntSushi/toml"
  "os"
)

const Version = "0.0.1"

type Config struct {
  Project    ProjectConfig    `toml:"project"`
  Module     ModuleConfig     `toml:"module"`
  MainSystem MainSystemConfig `toml:"main-system"`
}

type ProjectConfig struct {
  Name   string `toml:"name"`
  RpcURL string `toml:"rpc_url"`
  WsURL  string `toml:"ws_url"`
}

type ModuleConfig struct {
  HealthCheck bool `toml:"health_check"`
  DiskSpace   bool `toml:"disk_space"`
  Voting      bool `toml:"voting"`
  IBCTransfer bool `toml:"ibc_transfer"`
}

type MainSystemConfig struct {
  ApiURL string `toml:"api_url"`
}

func (c *Config) Validate() error {
  if c.Project.Name == "" {
    return fmt.Errorf("project name is required")
  }
  if c.Project.RpcURL == "" {
    return fmt.Errorf("rpc url is required")
  }
  if c.Project.WsURL == "" {
    return fmt.Errorf("ws url is required")
  }
  if c.MainSystem.ApiURL == "" {
    return fmt.Errorf("main system api url is required")
  }

  return nil
}

func LoadConfig(path string) (*Config, error) {
  config := &Config{}

  if _, err := os.Stat(path); os.IsNotExist(err) {
    return nil, fmt.Errorf("config file does not exist: %s", path)
  }

  if _, err := toml.DecodeFile(path, config); err != nil {
    return nil, fmt.Errorf("failed to decode config file: %w", err)
  }

  if err := config.Validate(); err != nil {
    return nil, fmt.Errorf("invalid config: %w", err)
  }

  return config, nil
}
