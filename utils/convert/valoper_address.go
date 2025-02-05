package convert

import (
  "fmt"

  sdk "github.com/cosmos/cosmos-sdk/types"
)

func ConvertToValoperAddress(address string) (string, error) {
  prefix := GetPrefix(address) + "valoper"

  accAddr, err := sdk.AccAddressFromBech32(address)
  if err != nil {
    return "", fmt.Errorf("invalid address: %v", err)
  }

  valAddr := sdk.ValAddress(accAddr)
  bech32ValAddr, err := sdk.Bech32ifyAddressBytes(prefix, valAddr)
  if err != nil {
    return "", fmt.Errorf("failed to convert to valoper address: %v", err)
  }

  return bech32ValAddr, nil
}
