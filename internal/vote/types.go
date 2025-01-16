package vote

import (
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

type VoteMonitor struct {
  client      *rpc.RPCClient
  queryClient slashingTypes.QueryClient
  done        chan struct{}
}
