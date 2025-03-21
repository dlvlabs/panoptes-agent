package main

import (
  "log"
  "os"

  "dlvlabs.net/panoptes-agent/app"
  "dlvlabs.net/panoptes-agent/cmd"
)

func main() {
  if len(os.Args) > 1 {
    if err := cmd.ExecuteCLI(); err != nil {
      log.Fatal(err)
    }
    return
  }

  if err := app.Execute(); err != nil {
    log.Fatal(err)
  }
}
