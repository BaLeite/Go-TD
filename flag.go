package main

import "flag"
import "fmt"

const VERSION = "1.0"

func main() {
  version := flag.Bool("version", false, "Demandez la version")
  flag.Parse()
  if *version {
    fmt.Println(VERSION)
    os.Exit(0)
  }
}
  
