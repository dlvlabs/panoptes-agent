package agent

import (
  "context"

  "dlvlabs.net/panoptes-agent/config"
  "dlvlabs.net/panoptes-agent/internal/block"
  "dlvlabs.net/panoptes-agent/internal/disk"
  "dlvlabs.net/panoptes-agent/internal/validator"
)

type Agent struct {
  cfg          *config.Config
  ctx          context.Context
  cancel       context.CancelFunc
  minutes      int
  blockMonitor *block.BlockMonitor
  diskMonitor  *disk.DiskMonitor
  validator    *validator.ValidatorMonitor
}
