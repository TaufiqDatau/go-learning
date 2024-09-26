package main

import (
	"fmt"
)

func main(){

  counter :=1
  for counter<=10{
    fmt.Println("Ini iterasi ke ", counter)
    counter++
  }

  lsStudent := []string{"Marc","Cathy","Ridel","Nicole","Raya"}

  for _,name := range lsStudent {
    fmt.Println(name)
  }

  gameByGenre := map[string][]string{
    "slice of life":[]string{"Until Then", "A Space For The Unbound"},
    "fps":[]string{"Warhammer","Counter Strike", "Valorant", "Overwatch"},
    "Open World":[]string{"Pal World", "Grand Theft Auto"},
  }

  for key,value := range gameByGenre{
    fmt.Print(key, " :")
    for _,gameName := range value{
      fmt.Println(" ",gameName)
    }
  }
}
