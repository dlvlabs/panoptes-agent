package agent

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/config"
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "dlvlabs.net/panoptes-agent/internal/block"
  "dlvlabs.net/panoptes-agent/internal/disk"
  "dlvlabs.net/panoptes-agent/utils/schedule"
)

type Agent struct {
  cfg     *config.Config
  ctx     context.Context
  cancel  context.CancelFunc
  minutes int
}

func NewAgent(cfg *config.Config) *Agent {
  ctx, cancel := context.WithCancel(context.Background())
  return &Agent{
    cfg:     cfg,
    ctx:     ctx,
    cancel:  cancel,
    minutes: cfg.Agent.DataSendInterval,
  }
}

func (m *Agent) Start() error {
  if m.cfg.Feature.BlockHeight {
    blockSchedule := schedule.NewMonitorSchedule(m.ctx, m.minutes)
    rpcClient, err := rpc.NewRPCClient(&m.ctx, m.cfg.BlockHeightConfig.RpcURL)
    if err != nil {
      return err
    }

    blockMonitor := block.NewBlockMonitor(rpcClient)
    if err := blockMonitor.Start(m.ctx, blockSchedule); err != nil {
      return err
    }
    log.Println("Block height monitoring started")
  }

  if m.cfg.Feature.DiskSpace {
    diskSchedule := schedule.NewMonitorSchedule(m.ctx, m.minutes)
    diskMonitor := disk.NewDiskMonitor(m.cfg.DiskSpaceConfig.Paths)
    if err := diskMonitor.Start(m.ctx, diskSchedule); err != nil {
      return err
    }
    log.Println("Disk space monitoring started")
  }

  return nil
}

func (m *Agent) Stop() {
  m.cancel()
}
