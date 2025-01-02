package block

import (
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
)

type BlockMonitor struct {
  client *rpc.RPCClient
  done   chan struct{}
}
