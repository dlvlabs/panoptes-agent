package block

import (
  "context"
  "fmt"
  "log"

  "github.com/cometbft/cometbft/rpc/client/http"
  coretypes "github.com/cometbft/cometbft/rpc/core/types"
)

func getBlockHeight(c *http.HTTP, ctx context.Context) error {
  abciInfo, err := c.ABCIInfo(ctx)
  if err != nil {
    return fmt.Errorf("failed to get ABCI info: %v", err)
  }

  height := abciInfo.Response.LastBlockHeight
  printABCIInfo(abciInfo)
  log.Printf("Current block height: %d\n", height)

  return nil
}

func printABCIInfo(info *coretypes.ResultABCIInfo) {
  log.Printf("LastBlockHeight: %d\n", info.Response.LastBlockHeight)
}
