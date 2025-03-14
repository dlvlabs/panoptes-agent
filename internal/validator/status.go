package validator

import (
  "context"
  "fmt"
  "log"

  stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func (v *ValidatorMonitor) GetValidatorStatus(ctx context.Context) error {
  req := stakingTypes.QueryValidatorRequest{
    ValidatorAddr: v.validatorOperatorAddress,
  }
  data, err := req.Marshal()
  if err != nil {
    return fmt.Errorf("failed to marshal request: %w", err)
  }

  result, err := v.client.GetClient().ABCIQuery(ctx, "/cosmos.staking.v1beta1.Query/Validator", data)
  if err != nil {
    return fmt.Errorf("failed to query validator: %w", err)
  }

  var response stakingTypes.QueryValidatorResponse
  if err := response.Unmarshal(result.Response.Value); err != nil {
    return fmt.Errorf("failed to unmarshal response: %w", err)
  }
  log.Printf("Validator moniker: %s\n", response.Validator.GetMoniker())
  log.Println("Validator status:", response.Validator.GetStatus())

  return nil
}
