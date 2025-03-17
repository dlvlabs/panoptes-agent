package schedule

import (
  "context"
  "time"
)

func monitorSchedule(ctx context.Context, minutes int) <-chan time.Time {
  if minutes <= 0 {
    return nil
  }

  // 이거 주입받으면 되는거 아님 ?
  ticker := time.NewTicker(time.Duration(minutes) * time.Minute)
  ch := make(chan time.Time, 1)

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

func NewMonitorSchedule(ctx context.Context, minutes int) <-chan time.Time {
  return monitorSchedule(ctx, minutes)
}
