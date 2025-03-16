package main

import (
  "log"
  "os"

  "dlvlabs.net/panoptes-agent/cmd/app"
  "dlvlabs.net/panoptes-agent/cmd/cli"
)

func main() {
  if len(os.Args) > 1 {
    if err := cli.ExecuteCLI(); err != nil {
      log.Fatal(err)
    }
    return
  }

  if err := app.Execute(); err != nil {
    log.Fatal(err)
  }
}
