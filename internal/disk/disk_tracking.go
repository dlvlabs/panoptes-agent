package disk

import (
  "fmt"
  "syscall"
)

func (d *DiskMonitor) getDiskUsage(path string) (*diskUsage, error) {
  var stat syscall.Statfs_t
  err := syscall.Statfs(path, &stat)
  if err != nil {
    return nil, fmt.Errorf("failed to get disk stats for %s: %v", path, err)
  }

  total := stat.Blocks * uint64(stat.Bsize)
  free := stat.Bfree * uint64(stat.Bsize)
  used := total - free
  usagePercent := float64(used) / float64(total) * 100

  totalInodes := stat.Files
  freeInodes := stat.Ffree
  usedInodes := totalInodes - freeInodes
  inodePercent := float64(usedInodes) / float64(totalInodes) * 100

  return &diskUsage{
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

func (d *DiskMonitor) monitorAll() ([]*diskUsage, error) {
  var results []*diskUsage

  for _, path := range d.Paths {
    usage, err := d.getDiskUsage(path)
    if err != nil {
      return nil, err
    }
    results = append(results, usage)
  }

  return results, nil
}
