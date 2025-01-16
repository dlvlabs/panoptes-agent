package vote

import (
  "context"
  "fmt"
  "log"
  "time"

  "dlvlabs.net/panoptes-agent/infrastructure/client/rpc"
  "github.com/cometbft/cometbft/rpc/client/http"
  slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

func NewVoteMonitor(client *rpc.RPCClient) *VoteMonitor {
  return &VoteMonitor{
    client: client,
    done:   make(chan struct{}),
  }
}
func getVoteStatus(c *http.HTTP, ctx context.Context, consAddress string) error {
  req := slashingTypes.QuerySigningInfoRequest{
    ConsAddress: consAddress,
  }
  data, err := req.Marshal()
  if err != nil {
    return fmt.Errorf("failed to marshal request: %w", err)
  }

  result, err := c.ABCIQuery(ctx, "/cosmos.slashing.v1beta1.Query/SigningInfo", data)
  if err != nil {
    return fmt.Errorf("failed to query signing info: %w", err)
  }

  var response slashingTypes.QuerySigningInfoResponse
  if err := response.Unmarshal(result.Response.Value); err != nil {
    return fmt.Errorf("failed to unmarshal response: %w", err)
  }
  fmt.Println(response.ValSigningInfo.MissedBlocksCounter)
  return nil
}

func (v *VoteMonitor) Start(ctx context.Context, schedule <-chan time.Time, consAddress string) error {
  go func() {
    for {
      select {
      case <-schedule:
        if err := getVoteStatus(v.client.GetClient(), ctx, consAddress); err != nil {
          log.Printf("Error getting vote status: %v", err)
        }
        log.Println("Vote monitoring started")
      case <-ctx.Done():
        return
      case <-v.done:
        return
      }
    }
  }()
  return nil
}

func (v *VoteMonitor) Stop() {
  close(v.done)
}
