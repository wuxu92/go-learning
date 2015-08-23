package main

import (
  "fmt"
)

func main() {
  go values();
  _ = fmt.Println;
  variables()
  constants()
  forThins()
  switchThings()
  arrays()
  slices()
  maps()
  structs()
  methods()
  inters()
  errors()
  channels()
  chanSync()
  selects()
}

func get() {
  fmt.Println("new func")

  variables()
}
