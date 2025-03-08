package validator

import (
  "context"
  "fmt"
  "log"

  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func (v *ValidatorMonitor) GetValidatorMissedBlocks(ctx context.Context) error {
  req := slashingTypes.QuerySigningInfoRequest{
    ConsAddress: v.consAddress,
  }
  data, err := req.Marshal()
  if err != nil {
    return fmt.Errorf("failed to marshal request: %w", err)
  }

  result, err := v.client.GetClient().ABCIQuery(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfo", data)
  if err != nil {
    return fmt.Errorf("failed to query signing info: %w", err)
  }

  var response slashingTypes.QuerySigningInfoResponse
  if err := response.Unmarshal(result.Response.Value); err != nil {
    return fmt.Errorf("failed to unmarshal response: %w", err)
  }
  log.Printf("Missed blocks: %d", response.ValSigningInfo.GetMissedBlocksCounter())

  return nil
}
