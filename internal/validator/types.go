package validator

import (
  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
)

type ValidatorMonitor struct {
  accAddress               string
  consAddress              string
  validatorOperatorAddress string

  client *rpc.RPCClient
  done   chan struct{}
}
