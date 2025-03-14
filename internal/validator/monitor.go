package validator

import (
  "context"
  "log"
  "time"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
)

func NewValidatorMonitor(client *rpc.RPCClient, accAddress string) *ValidatorMonitor {
  consAddress, err := ConvertToConsAddress(client.GetClient(), context.Background(), accAddress)
  if err != nil {
    log.Fatal(err)
  }

  valAddress, err := ConvertToValoperAddress(accAddress)
  if err != nil {
    log.Fatal(err)
  }

  return &ValidatorMonitor{
    accAddress:               accAddress,
    validatorOperatorAddress: valAddress,
    consAddress:              consAddress,
    client:                   client,
    done:                     make(chan struct{}),
  }
}

func (v *ValidatorMonitor) Start(ctx context.Context, schedule <-chan time.Time) error {
  go func() {
    for {
      select {
      case <-schedule:
        if err := v.GetValidatorMissedBlocks(ctx); err != nil {
          log.Printf("Error getting vote status: %v", err)
        }
        if err := v.GetValidatorStatus(ctx); err != nil {
          log.Printf("Error getting validator status: %v", err)
        }
        log.Println("Validator monitoring started")
      case <-ctx.Done():
        return
      case <-v.done:
        return
      }
    }
  }()
  return nil
}

func (v *ValidatorMonitor) Stop() {
  close(v.done)
}
