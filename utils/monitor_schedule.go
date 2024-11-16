package utils

import (
  "context"
  "time"
)

func MonitorSchedule(ctx context.Context, minutes int) <-chan time.Time {
  ticker := time.NewTicker(time.Duration(minutes) * time.Minute)
  ch := make(chan time.Time)

  go func() {
    defer ticker.Stop()
    defer close(ch)

    ch <- time.Now()

    for {
      select {
      case <-ctx.Done():
        return
      case t := <-ticker.C:
        ch <- t
      }
    }
  }()

  return ch
}
