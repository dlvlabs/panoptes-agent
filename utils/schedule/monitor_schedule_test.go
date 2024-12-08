package schedule

import (
  "context"
  "sync"
  "testing"
  "time"
)

func TestMonitorSchedule(t *testing.T) {
  tests := []struct {
    name        string
    minutes     int
    checkTimes  int
    shouldStop  bool
    expectError bool
  }{
    {
      name:       "basic_schedule_test",
      minutes:    1,
      checkTimes: 2,
      shouldStop: false,
    },
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
      defer cancel()

      testMinutes := 1
      if tt.minutes > 0 {
        testMinutes = 1
      } else {
        testMinutes = tt.minutes
      }

      schedule := MonitorSchedule(ctx, testMinutes)
      if tt.expectError {
        if schedule != nil {
          t.Error("expected nil schedule for invalid interval")
        }
        return
      }

      testDuration := 1 * time.Minute
      for i := 0; i < tt.checkTimes; i++ {
        select {
        case _, ok := <-schedule:
          if !ok {
            t.Fatal("channel closed unexpectedly")
          }
        case <-time.After(testDuration):
          t.Errorf("trigger %d not received in time", i+1)
          return
        }
      }

      if tt.shouldStop {
        cancel()
        select {
        case _, ok := <-schedule:
          if ok {
            t.Error("channel should be closed after context cancellation")
          }
        case <-time.After(testDuration):
          t.Error("channel not closed after context cancellation")
        }
      }
    })
  }
}

func TestMonitorScheduleConcurrency(t *testing.T) {
  ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
  defer cancel()

  schedule := MonitorSchedule(ctx, 1)

  const goroutines = 5
  var wg sync.WaitGroup
  wg.Add(goroutines)

  for i := 0; i < goroutines; i++ {
    go func() {
      defer wg.Done()
      for {
        select {
        case _, ok := <-schedule:
          if !ok {
            return
          }
        case <-ctx.Done():
          return
        }
      }
    }()
  }

  <-ctx.Done()

  done := make(chan struct{})
  go func() {
    wg.Wait()
    close(done)
  }()

  select {
  case <-done:

  case <-time.After(100 * time.Millisecond):
    t.Error("goroutines did not finish in time")
  }
}
