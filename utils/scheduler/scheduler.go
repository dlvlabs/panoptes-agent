package scheduler

import (
  "context"
  "fmt"
  "time"
)

func (s Scheduler) Execute() <-chan time.Time {

  ch := make(chan time.Time, 1)

  go func() {
    defer s.ticker.Stop()
    defer close(ch)

    select {
    case ch <- time.Now():
    case <-s.ctx.Done():
      return
    }

    for {
      select {
      case t := <-s.ticker.C:
        select {
        case ch <- t:
        case <-s.ctx.Done():
          return
        }
      case <-s.ctx.Done():
        return
      }
    }
  }()

  return ch
}

func NewMonitorScheduler(ctx context.Context, minutes int) (Scheduler, error) {
  if minutes <= 0 {
    return Scheduler{}, fmt.Errorf("minutes must be greater than 0")
  }
  ticker := time.NewTicker(time.Duration(minutes) * time.Minute)
  return Scheduler{ctx: ctx, ticker: ticker}, nil
}
