package rpc

import "github.com/cometbft/cometbft/rpc/client/http"

type Client interface {
  GetClient() *http.HTTP
  Close() error
}

type RPCClient struct {
  client *http.HTTP
}
