package disk

import (
  "context"
  "log"

  "dlvlabs.net/panoptes-agent/utils/scheduler"
)

func NewDiskMonitor(paths []string) *DiskMonitor {
  return &DiskMonitor{
    Paths: paths,
    done:  make(chan struct{}),
  }
}

func (d *DiskMonitor) Start(ctx context.Context, schedule scheduler.Scheduler) error {
  scheduleCh := schedule.Execute()
  go func() {
    for {
      select {
      case <-scheduleCh:
        usages, err := d.monitorAll()
        if err != nil {
          log.Printf("Error monitoring disk usage: %v", err)
          continue
        }

        for _, usage := range usages {
          log.Printf("Disk usage for %s:\n"+
            "\tSpace: %.2f%% (Used: %.2f GB, Free: %.2f GB, Total: %.2f GB)\n"+
            "\tInodes: %.2f%% (Used: %d, Free: %d, Total: %d)",
            usage.Path,
            usage.UsagePercent,
            usage.UsedGB,
            usage.FreeGB,
            usage.TotalGB,
            usage.InodePercent,
            usage.UsedInodes,
            usage.FreeInodes,
            usage.TotalInodes)
        }

      case <-ctx.Done():
        return
      case <-d.done:
        return
      }
    }
  }()

  return nil
}

func (d *DiskMonitor) Stop() {
  close(d.done)
}
