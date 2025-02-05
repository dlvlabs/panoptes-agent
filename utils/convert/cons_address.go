package convert

import (
  "context"
  "crypto/sha256"
  "encoding/base64"
  "encoding/json"
  "fmt"

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
  // fmt.Printf("Querying validator with address: %s\n", req.ValidatorAddr)

  data, err := req.Marshal()
  if err != nil {
    return "", fmt.Errorf("failed to marshal request: %w", err)
  }
  // fmt.Printf("Request data (hex): %x\n", data)

  // fmt.Printf("Sending ABCI query to: %s\n", "/cosmos.staking.v1beta1.Query/Validator")
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
  fmt.Printf("Pubkey key value: %s\n", pubkeyBase64)
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

  fmt.Printf("Validator consensus address: %s\n", valConsAddress)
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
  // fmt.Printf("Generated consensus address bytes: %x\n", consAddr.Bytes())

  bech32ConsAddr, err := sdk.Bech32ifyAddressBytes("cosmosvalcons", consAddr)
  if err != nil {
    return "", fmt.Errorf("failed to convert to consensus address: %w", err)
  }
  // fmt.Printf("Final bech32 consensus address: %s\n", bech32ConsAddr)

  return bech32ConsAddr, nil
}
