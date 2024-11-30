package schedule

import (
  "context"
  "time"
)

func MonitorSchedule(ctx context.Context, minutes int) <-chan time.Time {
  if minutes <= 0 {
    return nil
  }

  interval := time.Duration(minutes) * time.Millisecond
  ticker := time.NewTicker(interval)
  ch := make(chan time.Time)

  go func() {
    defer ticker.Stop()
    defer close(ch)

    select {
    case ch <- time.Now():
    case <-ctx.Done():
      return
    }

    for {
      select {
      case t := <-ticker.C:
        select {
        case ch <- t:
        case <-ctx.Done():
          return
        }
      case <-ctx.Done():
        return
      }
    }
  }()

  return ch
}
