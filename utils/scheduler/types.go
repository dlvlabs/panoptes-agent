package scheduler

import (
  "context"
  "time"
)

type Scheduler struct {
  ctx    context.Context
  ticker *time.Ticker
}
