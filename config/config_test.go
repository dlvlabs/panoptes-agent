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

    assert.Equal(t, "cosmos-hub", config.Project.Name)
    assert.Equal(t, "http://localhost:11190", config.Project.RpcURL)
    assert.Equal(t, "http://localhost:11157", config.Project.WsURL)
    assert.Equal(t, true, config.Module.HealthCheck)
    assert.Equal(t, true, config.Module.DiskSpace)
    assert.Equal(t, false, config.Module.Voting)
    assert.Equal(t, false, config.Module.IBCTransfer)
    assert.Equal(t, "http:localhost:8080", config.MainSystem.ApiURL)
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

func TestValidate(t *testing.T) {
  t.Run("validate config", func(t *testing.T) {
    config := &Config{
      Project: ProjectConfig{
        Name:   "cosmos-hub",
        RpcURL: "http://localhost:11190",
        WsURL:  "http://localhost:11157",
      },
      Module: ModuleConfig{
        HealthCheck: true,
        DiskSpace:   true,
        Voting:      false,
        IBCTransfer: false,
      },
      MainSystem: MainSystemConfig{
        ApiURL: "http://localhost:8080",
      },
    }
    err := config.Validate()
    require.NoError(t, err)
  })
  t.Run("invalidate config", func(t *testing.T) {
    config := &Config{
      Project: ProjectConfig{
        Name:   "",
        RpcURL: "",
        WsURL:  "",
      },
      Module: ModuleConfig{
        HealthCheck: true,
        DiskSpace:   true,
        Voting:      false,
        IBCTransfer: false,
      },
      MainSystem: MainSystemConfig{
        ApiURL: "",
      },
    }
    err := config.Validate()

    require.Error(t, err)
    assert.Contains(t, err.Error(), "project name is required")

  })
}
