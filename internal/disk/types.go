package disk

type DiskMonitor struct {
  Paths []string
  done  chan struct{}
}

type diskUsage struct {
  Path         string
  TotalGB      float64
  UsedGB       float64
  FreeGB       float64
  UsagePercent float64

  TotalInodes  uint64
  UsedInodes   uint64
  FreeInodes   uint64
  InodePercent float64
}
