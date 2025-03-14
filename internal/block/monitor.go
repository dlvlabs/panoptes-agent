package block

import (
  "context"
  "log"
  "time"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
)

func NewBlockMonitor(client *rpc.RPCClient) *BlockMonitor {
  return &BlockMonitor{
    client: client,
    done:   make(chan struct{}),
  }
}

func (m *BlockMonitor) Start(ctx context.Context, schedule <-chan time.Time) error {
  defer m.client.Close()
  go func() {
    for {
      select {
      case <-schedule:
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
