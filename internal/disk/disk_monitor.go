package disk

import (
  "context"
  "fmt"
  "log"
  "syscall"
  "time"
)

type DiskMonitor struct {
  Paths []string
}

type DiskUsage struct {
  Path         string
  TotalGB      float64
  UsedGB       float64
  FreeGB       float64
  UsagePercent float64

  // inode 정보 추가
  TotalInodes  uint64
  UsedInodes   uint64
  FreeInodes   uint64
  InodePercent float64
}

func NewDiskMonitor(paths []string) *DiskMonitor {
  return &DiskMonitor{
    Paths: paths,
  }
}

func bytesToGB(bytes uint64) float64 {
  return float64(bytes) / 1024 / 1024 / 1024
}

func (d *DiskMonitor) GetDiskUsage(path string) (*DiskUsage, error) {
  var stat syscall.Statfs_t
  err := syscall.Statfs(path, &stat)
  if err != nil {
    return nil, fmt.Errorf("failed to get disk stats for %s: %v", path, err)
  }

  total := stat.Blocks * uint64(stat.Bsize)
  free := stat.Bfree * uint64(stat.Bsize)
  used := total - free
  usagePercent := float64(used) / float64(total) * 100

  // inode 계산
  totalInodes := stat.Files
  freeInodes := stat.Ffree
  usedInodes := totalInodes - freeInodes
  inodePercent := float64(usedInodes) / float64(totalInodes) * 100

  return &DiskUsage{
    Path:         path,
    TotalGB:      bytesToGB(total),
    UsedGB:       bytesToGB(used),
    FreeGB:       bytesToGB(free),
    UsagePercent: usagePercent,
    TotalInodes:  totalInodes,
    UsedInodes:   usedInodes,
    FreeInodes:   freeInodes,
    InodePercent: inodePercent,
  }, nil
}

func (d *DiskMonitor) MonitorAll() ([]*DiskUsage, error) {
  var results []*DiskUsage

  for _, path := range d.Paths {
    usage, err := d.GetDiskUsage(path)
    if err != nil {
      return nil, err
    }
    results = append(results, usage)
  }

  return results, nil
}

func (d *DiskMonitor) Start(ctx context.Context, schedule <-chan time.Time) error {
  go func() {
    for {
      select {
      case <-schedule:
        usages, err := d.MonitorAll()
        if err != nil {
          log.Printf("Error monitoring disk usage: %v", err)
          continue
        }

        for _, usage := range usages {
          log.Printf("Disk usage for %s:\n"+
            "Space: %.2f%% (Used: %.2f GB, Free: %.2f GB, Total: %.2f GB)\n"+
            "Inodes: %.2f%% (Used: %d, Free: %d, Total: %d)",
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
      }
    }
  }()

  return nil
}
