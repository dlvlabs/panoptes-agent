package agent

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/config"
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "dlvlabs.net/panoptes-agent/internal/block"
  "dlvlabs.net/panoptes-agent/internal/disk"
  "dlvlabs.net/panoptes-agent/internal/validator"
  "dlvlabs.net/panoptes-agent/utils/scheduler"
)

func NewAgent(cfg *config.Config) *Agent {
  ctx, cancel := context.WithCancel(context.Background())
  return &Agent{
    cfg:     cfg,
    ctx:     ctx,
    cancel:  cancel,
    minutes: cfg.Agent.DataSendInterval,
  }
}

func (a *Agent) Start() error {
  if a.cfg.Feature.BlockHeight {
    blockScheduler, err := scheduler.NewMonitorScheduler(a.ctx, a.minutes)
    if err != nil {
      return err
    }
    rpcClient, err := rpc.NewRPCClient(&a.ctx, a.cfg.Agent.RpcURL)
    if err != nil {
      return err
    }

    a.blockMonitor = block.NewBlockMonitor(rpcClient)
    if err := a.blockMonitor.Start(a.ctx, blockScheduler); err != nil {
      return err
    }

    log.Println("Block height monitoring started")
  }

  if a.cfg.Feature.DiskSpace {
    diskScheduler, err := scheduler.NewMonitorScheduler(a.ctx, a.minutes)
    if err != nil {
      return err
    }
    a.diskMonitor = disk.NewDiskMonitor(a.cfg.DiskSpaceConfig.Paths)
    if err := a.diskMonitor.Start(a.ctx, diskScheduler); err != nil {
      return err
    }
    log.Println("Disk space monitoring started")
  }

  if a.cfg.Feature.ValidatorMassage {
    validatorScheduler, err := scheduler.NewMonitorScheduler(a.ctx, a.minutes)
    if err != nil {
      return err
    }
    rpcClient, err := rpc.NewRPCClient(&a.ctx, a.cfg.Agent.RpcURL)
    if err != nil {
      return err
    }
    a.validator = validator.NewValidatorMonitor(rpcClient, a.cfg.ValidatorMassageConfig.AccAddress)
    if err := a.validator.Start(a.ctx, validatorScheduler); err != nil {
      return err
    }
    log.Println("Validator monitoring started")
  }

  return nil
}

func (a *Agent) Stop() {
  if a.blockMonitor != nil {
    a.blockMonitor.Stop()
  }

  if a.diskMonitor != nil {
    a.diskMonitor.Stop()
  }

  if a.validator != nil {
    a.validator.Stop()
  }

  a.cancel()
}
