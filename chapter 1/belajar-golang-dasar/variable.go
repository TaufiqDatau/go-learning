package main

import "fmt"

func main(){
  var name string

  name = "Taufiqurrahman"
  fmt.Println(name)

  name = "Taufiqurrahman Hafiidh Datau"
  fmt.Println(name)

  var newName = "Taufiq"
  fmt.Println(newName)

  carType := "Maclaren"
  fmt.Println(carType)

  var (
    firstName = "Taufiqurrahman"
    middleName = "Hafiidh"
    lastName = "Datau"
  )

  fmt.Println(firstName, middleName, lastName)
}