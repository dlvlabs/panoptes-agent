package validator

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "dlvlabs.net/panoptes-agent/utils/scheduler"
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

func (v *ValidatorMonitor) Start(ctx context.Context, schedule scheduler.Scheduler) error {
  scheduleCh := schedule.Execute()
  go func() {
    for {
      select {
      case <-scheduleCh:
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
