package validator

import (
  "context"
  "fmt"
  "log"
  "time"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "dlvlabs.net/panoptes-agent/utils/convert"
  "github.com/cometbft/cometbft/rpc/client/http"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func NewValidatorMonitor(client *rpc.RPCClient, accAddress string) *ValidatorMonitor {
  consAddress, err := convert.ConvertToConsAddress(client.GetClient(), context.Background(), accAddress)
  if err != nil {
    log.Fatal(err)
  }

  valAddress, err := convert.ConvertToValoperAddress(accAddress)
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
func getValidatorStatus(c *http.HTTP, ctx context.Context, consAddress string) error {
  req := slashingTypes.QuerySigningInfoRequest{
    ConsAddress: consAddress,
  }
  data, err := req.Marshal()
  if err != nil {
    return fmt.Errorf("failed to marshal request: %w", err)
  }

  result, err := c.ABCIQuery(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfo", data)
  if err != nil {
    return fmt.Errorf("failed to query signing info: %w", err)
  }

  var response slashingTypes.QuerySigningInfoResponse
  if err := response.Unmarshal(result.Response.Value); err != nil {
    return fmt.Errorf("failed to unmarshal response: %w", err)
  }
  fmt.Println(response.ValSigningInfo.MissedBlocksCounter)

  return nil
}

func (v *ValidatorMonitor) Start(ctx context.Context, schedule <-chan time.Time, accAddress string) error {
  go func() {
    for {
      select {
      case <-schedule:
        if err := getValidatorStatus(v.client.GetClient(), ctx, accAddress); err != nil {
          log.Printf("Error getting vote status: %v", err)
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
