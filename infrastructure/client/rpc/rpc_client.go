package rpc

import (
  "context"
  "fmt"
  "github.com/cometbft/cometbft/rpc/client/http"
)

func NewRPCClient(ctx *context.Context, rpcURL string) (*RPCClient, error) {
  client, err := http.New(rpcURL, "")
  if err != nil {
    return nil, fmt.Errorf("failed to create RPC client: %w", err)
  }

  return &RPCClient{
    client: client,
  }, nil
}

func (c *RPCClient) Close() error {
  if c.client != nil {
    err := c.client.Stop()
    if err != nil {
      return fmt.Errorf("failed to close RPC client: %w", err)
    }
  }
  return nil
}

func (c *RPCClient) GetClient() *http.HTTP {
  return c.client
}
