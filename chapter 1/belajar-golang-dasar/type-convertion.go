package main

import "fmt"

func main(){
  var x int32 = 32768
  var y int64 = int64(x)
  var z int16 = int16(y)

  fmt.Println("this is the value of x :",x)
  fmt.Println("this is the value of y :",y)
  fmt.Println("this is the value of z :",z)

}