package main

import "fmt"

func main(){
  const radius float32 = 4
  const pi float32 = 3.14
  const circumferenc float32 = 2*radius*pi

  fmt.Println(circumferenc)

  //error
  radius = 2
}
