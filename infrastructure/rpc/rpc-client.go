package rpc

import (
  "context"
  "github.com/cometbft/cometbft/rpc/client/http"
  "log"
)

func NewClient(ctx *context.Context, rpcURL string) (*http.HTTP, error) {

  client, err := http.New(rpcURL, "")
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  if err != nil {

  }
  return client, nil
}
