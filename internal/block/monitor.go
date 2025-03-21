package block

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "dlvlabs.net/panoptes-agent/utils/scheduler"
)

func NewBlockMonitor(client *rpc.RPCClient) *BlockMonitor {
  return &BlockMonitor{
    client: client,
    done:   make(chan struct{}),
  }
}

func (m *BlockMonitor) Start(ctx context.Context, schedule scheduler.Scheduler) error {
  scheduleCh := schedule.Execute()
  defer m.client.Close()
  go func() {
    for {
      select {
      case <-scheduleCh:
        if err := getBlockHeight(m.client.GetClient(), ctx); err != nil {
          log.Printf("Error getting block height: %v", err)
        }
      case <-ctx.Done():
        return
      case <-m.done:
        return

      }
    }
  }()

  return nil
}

func (m *BlockMonitor) Stop() {
  close(m.done)
}
