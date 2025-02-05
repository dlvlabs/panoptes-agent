package convert

import "strings"

func GetPrefix(address string) string {
  if address == "" {
    return ""
  }
  parts := strings.Split(address, "1")
  if len(parts) < 2 {
    return ""
  }
  return parts[0]
}
