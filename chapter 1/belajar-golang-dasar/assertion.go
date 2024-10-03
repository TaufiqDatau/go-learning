package main

import "fmt"

func random() any{
  return "OK"
}

func main(){
  result := random()
  var resultString string = result.(string) // Type Assertion
  fmt.Println(resultString)

  //better way to do type assertion 
  switch value := result.(type){
  case string :
    fmt.Println("This is a string", value)
  case int :
    fmt.Println("This is an int", value)
  default:
    fmt.Println("unknown")
  }

}
