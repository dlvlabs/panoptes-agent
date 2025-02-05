package validator

import (
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

type ValidatorMonitor struct {
  accAddress               string
  voAddress                string
  consAddress              string
  validatorOperatorAddress string

  client      *rpc.RPCClient
  queryClient slashingTypes.QueryClient
  done        chan struct{}
}
