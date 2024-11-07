package config

import (
  "testing"

  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/require"
)

func TestVersion(t *testing.T) {
  assert.Equal(t, "0.0.1", Version)
}

func TestLoadConfig(t *testing.T) {
  t.Run("load valid config", func(t *testing.T) {
    config, err := LoadConfig("testdata/valid_config.toml")
    require.NoError(t, err)

    assert.Equal(t, "cosmos-hub", config.Agent.Name)
    assert.Equal(t, true, config.Feature.BlockHeight)
    assert.Equal(t, true, config.Feature.DiskSpace)
    assert.Equal(t, true, config.Feature.Voting)
    assert.Equal(t, false, config.Feature.IBCTransfer)
    assert.Equal(t, "http://localhost:8080", config.Agent.MainSystemUrl)
  })

  t.Run("load invalid config", func(t *testing.T) {
    _, err := LoadConfig("testdata/invalid_config.toml")
    require.Error(t, err)
    assert.Contains(t, err.Error(), "invalid config")
  })

  t.Run("config file does not exist", func(t *testing.T) {
    _, err := LoadConfig("config/testdata/non_existent.toml")
    require.Error(t, err)
    assert.Contains(t, err.Error(), "config file does not exist")
  })
}

func TestValidateAgent(t *testing.T) {
  t.Run("validate config", func(t *testing.T) {
    config := &Config{
      Agent: AgentConfig{
        Name:             "cosmos-hub",
        MainSystemUrl:    "http://localhost:8080",
        DataSendInterval: 5,
      },
    }
    err := config.ValidateAgent()
    require.NoError(t, err)
  })
  t.Run("invalidate config", func(t *testing.T) {
    config := &Config{
      Agent: AgentConfig{
        Name: "",
      },
      Feature: FeatureConfig{
        BlockHeight: true,
        DiskSpace:   true,
        Voting:      true,
        IBCTransfer: false,
      },
    }
    err := config.ValidateAgent()
    require.Error(t, err)
    assert.Contains(t, err.Error(), "agent name is required")
  })
}
