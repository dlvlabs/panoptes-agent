package config

import (
  "fmt"
  "github.com/BurntSushi/toml"
  "os"
)

const Version = "0.0.1"

type Config struct {
  Agent             AgentConfig       `toml:"agent"`
  Feature           FeatureConfig     `toml:"feature"`
  BlockHeightConfig BlockHeightConfig `toml:"block-height"`
  DiskSpaceConfig   DiskSpaceConfig   `toml:"disk-space"`
  VoteConfig        VoteConfig        `toml:"vote"`
}

type AgentConfig struct {
  Name             string `toml:"name"`
  DataSendInterval int    `toml:"data_send_interval"`
  MainSystemUrl    string `toml:"main_system_url"`
  RpcURL           string `toml:"rpc_url"`
}

type BlockHeightConfig struct {
}

type DiskSpaceConfig struct {
  Paths []string `toml:"paths"`
}
type VoteConfig struct {
  // TODO: 삭제 후 일반 Address -> consAddress로 변환하는 기능 추가
  ConsAddress string `toml:"cons_address"`
}

type FeatureConfig struct {
  BlockHeight bool `toml:"block_height"`
  DiskSpace   bool `toml:"disk_space"`
  Vote        bool `toml:"vote"`
  IBCTransfer bool `toml:"ibc_transfer"`
}

func (c *Config) ValidateAgent() error {
  if c.Agent.Name == "" {
    return fmt.Errorf("agent name is required")
  }
  if c.Agent.MainSystemUrl == "" {
    return fmt.Errorf("main system url is required")
  }
  if c.Agent.DataSendInterval <= 0 {
    return fmt.Errorf("data send interval is more then 0")
  }
  return nil
}

func (c *Config) ValidateBlockHeightFeature() error {

  return nil
}
func (c *Config) ValidateDiskSpaceFeature() error {
  fmt.Println(c.DiskSpaceConfig.Paths)
  if len(c.DiskSpaceConfig.Paths) == 0 {
    return fmt.Errorf("to use the disk space feature, paths is required")
  }
  return nil
}
func (c *Config) ValidateVotingFeature() error {

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

  if err := config.ValidateAgent(); err != nil {
    return nil, fmt.Errorf("invalid config: %w", err)
  }

  if config.Feature.BlockHeight {
    if err := config.ValidateBlockHeightFeature(); err != nil {
      return nil, fmt.Errorf("invalid config: %w", err)
    }
  }
  if config.Feature.DiskSpace {
    if err := config.ValidateDiskSpaceFeature(); err != nil {
      return nil, fmt.Errorf("invalid config: %w", err)
    }
  }
  if config.Feature.Vote {
    if err := config.ValidateVotingFeature(); err != nil {
      return nil, fmt.Errorf("invalid config: %w", err)
    }
  }
  return config, nil
}
