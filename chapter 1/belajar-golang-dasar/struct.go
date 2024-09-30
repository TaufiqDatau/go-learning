package main

import (
	"fmt"
	"time"
)

type animal interface{
  sound() string;
  numberOfLeg() int;
}

type Cat struct{
  name string;
  age int8;
  breed string;
}

func (c Cat) getBirthYear(){
  currentYear := time.Now().Year()
  fmt.Println(c.name, "is born in", currentYear - int(c.age))
}

func getTheAnimaSound(a animal){
  fmt.Println("This animal have this sound", a.sound())
}

func (c Cat) sound() string{
  return "Meow"
}

func (c Cat) numberOfLeg() int{
  return 4
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

  getTheAnimaSound(tommy)
}
