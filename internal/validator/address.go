package validator

import (
  "context"
  "crypto/sha256"
  "encoding/base64"
  "encoding/json"
  "fmt"
  "strings"

  "github.com/cometbft/cometbft/rpc/client/http"
  sdk "github.com/cosmos/cosmos-sdk/types"
  stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func ConvertToConsAddress(c *http.HTTP, ctx context.Context, address string) (string, error) {

  valAddr, err := ConvertToValoperAddress(address)
  if err != nil {
    return "", fmt.Errorf("failed to convert to valoper address: %w", err)
  }

  req := stakingTypes.QueryValidatorRequest{
    ValidatorAddr: valAddr,
  }

  data, err := req.Marshal()
  if err != nil {
    return "", fmt.Errorf("failed to marshal request: %w", err)
  }

  result, err := c.ABCIQuery(ctx, "/cosmos.staking.v1beta1.Query/Validator", data)
  if err != nil {
    return "", fmt.Errorf("failed to query validator: %w", err)
  }

  if result.Response.Value == nil {
    return "", fmt.Errorf("no validator data returned")
  }

  var response stakingTypes.QueryValidatorResponse
  if err := response.Unmarshal(result.Response.Value); err != nil {
    return "", fmt.Errorf("failed to unmarshal response: %w", err)
  }

  if response.Validator.ConsensusPubkey == nil {
    return "", fmt.Errorf("consensus pubkey is nil")
  }

  pubkeyBase64 := base64.StdEncoding.EncodeToString(response.Validator.ConsensusPubkey.Value[2:])
  pubKeyBytes, err := base64.StdEncoding.DecodeString(pubkeyBase64)
  if err != nil {
    panic(fmt.Errorf("failed to decode Base64 key: %w", err))
  }
  hash := sha256.Sum256(pubKeyBytes)

  valConsBytes := hash[:20]

  valConsAddress, err := sdk.Bech32ifyAddressBytes("cosmosvalcons", valConsBytes)
  if err != nil {
    panic(fmt.Errorf("failed to convert to Bech32: %w", err))
  }

  return valConsAddress, nil
}

type Ed25519PubKey struct {
  Type string `json:"@type"`
  Key  string `json:"key"`
}

func ConvertPubKeyToConsAddress(pubkeyJSON string) (string, error) {
  var pubKey Ed25519PubKey
  if err := json.Unmarshal([]byte(pubkeyJSON), &pubKey); err != nil {
    return "", fmt.Errorf("failed to unmarshal pubkey JSON: %w", err)
  }

  keyBytes, err := base64.StdEncoding.DecodeString(pubKey.Key)
  if err != nil {
    return "", fmt.Errorf("failed to decode base64 key: %w", err)
  }

  consAddr := sdk.ConsAddress(keyBytes)

  bech32ConsAddr, err := sdk.Bech32ifyAddressBytes("cosmosvalcons", consAddr)
  if err != nil {
    return "", fmt.Errorf("failed to convert to consensus address: %w", err)
  }

  return bech32ConsAddr, nil
}

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

func GetPrefix(address string) string {
  if address == "" {
    return ""
  }
  parts := strings.Split(address, "1")
  if len(parts) < 2 {
    return ""
  }
  return parts[0]
}
