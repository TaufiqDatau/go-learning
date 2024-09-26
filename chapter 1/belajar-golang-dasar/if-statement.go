package main

import "fmt"

func main(){
  var age = 70 

  if checkAge(age) {
    fmt.Println("you're allowed in the ride")
  }else{
    if age > 60 {
      fmt.Println("For health reason you're not allowed to enter the ride")
    }else{
      fmt.Println("Please wait another ", 13 - age, " years to enter this ride")
    } 
  }

}

func checkAge(age int) bool{

  if age > 60 {
    return false 
  } else if (age > 15){
    return true 
  }

  return false
}
