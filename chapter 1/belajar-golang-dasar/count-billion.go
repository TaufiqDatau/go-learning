package main

import (
  "fmt"
  "time"
)

func main(){
  start := time.Now()
  sum := 0
  
  for i:= 0 ; i <= 1_000_000_000 ; i++{
    sum += i

  }

  elapsed := time.Since(start)

  fmt.Printf("Time taken to count to 1 billion: %v\n", elapsed)
  fmt.Printf("Total sum off all number to 1 billion: %v\n", sum)
}
