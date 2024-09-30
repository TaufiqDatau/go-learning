package main

import (
	"fmt"
	"time"
)



type Cat struct{
  name string;
  age int8;
  breed string;
}

func (c Cat) getBirthYear(){
  currentYear := time.Now().Year()
  fmt.Println(c.name, "is born in", currentYear - int(c.age))
}

func main(){
  tommy := Cat{
    name: "Tommy",
    age: 4,
    breed: "Persians",
  }
  
  fmt.Println(tommy)
  fmt.Println(tommy.name)
  fmt.Println(tommy.age)
  fmt.Println(tommy.breed)

  tommy.getBirthYear()
}
