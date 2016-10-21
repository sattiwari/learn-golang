package main

import(
 "fmt"
 "runtime"
 "time"
 "strconv"
)

func showNumber(num int) {
  tstamp := strconv.FormatInt(time.Now().UnixNano(), 10)
  fmt.Println(num, tstamp)
  // time.Sleep(time.Millisecond * 10000)
}

func main() {
  // runtime.GOMAXPROCS(8)	
  iterations := 10
  for i := 0; i <= iterations; i++ {
	go showNumber(i)
  }
  runtime.Gosched()
  fmt.Println("hello")
}