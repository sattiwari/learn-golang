package main

import(
 "fmt"
 "runtime"
)

func main() {
  fmt.Println("number of CPU(s) %d", runtime.NumCPU())
  fmt.Println("number of threads %d", runtime.GOMAXPROCS(0))
}