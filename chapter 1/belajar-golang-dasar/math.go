package main

import "fmt"

func main(){
  var a = 10
  var b = 10
  var d = 5
  var e =2
  var c = a + b - d*e

  fmt.Println(c)

  a +=10
  fmt.Println(a)
  a-=5
  fmt.Println(a)
  a*=5
  fmt.Println(a)
  
  a/=5
  fmt.Println(a)

  a%=4
  fmt.Println(a)

  a++
  fmt.Println(a)

  a--
  fmt.Println(a)

  fmt.Println(10>11)
  fmt.Println("abc" == "abc")
  fmt.Println("abc" > "a")
  
}
