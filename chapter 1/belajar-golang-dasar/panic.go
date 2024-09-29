package main

import (
	"fmt"
)

func logging(){
  fmt.Println("function done")
  message := recover()

  if message != nil {
    fmt.Println(message)
  }
}

func runningApp(error bool, iter int, errorMessage string){
  defer logging()
  if error{
    panic(errorMessage)
  }
  fmt.Println("App is running", iter)
  
  err:=accessDb(true)

  if err != nil {
    fmt.Println(err)
  }
}

func accessDb(error bool) (error){
  defer logging()

  if error{
    panic("Error accessing database")
  }

  fmt.Println("accessing database")
  return nil
}

func main(){
  defer logging()
  runningApp(false, 1, "error happen in 1")
}
