package main

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/config"
  rpc "dlvlabs.net/panoptes-agent/infrastructure/rpc"
  blockheight "dlvlabs.net/panoptes-agent/internal/blockheight"
)

func main() {
  cfg, err := config.LoadConfig("config/config.toml")
  ctx := context.Background()
  if err != nil {
    log.Fatalf("Failed to load config: %v", err)
  }
  if cfg.Feature.BlockHeight {
    c, err := rpc.NewClient(&ctx, cfg.BlockHeightConfig.RpcURL)
    if err != nil {
      log.Fatalf("Failed to create rpc client: %v", err)
    }
    blockheight.GetBlockHeight(c, ctx)
  }
}
