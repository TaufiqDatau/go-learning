package main

import "fmt"

func main(){
  var goMap = map[string][]int{}
  goMap["Enigma"] = make([]int, 0,10)
  goMap["Indocyber"] = make([]int, 0,10)
  goMap["MCO"] = make([]int, 0,10)

  fmt.Println(goMap)
  newMap := goMap
  goMap["MCO"] = append(goMap["MCO"], 4)

  fmt.Println(newMap)

  newMap["Indocyber"] = append(newMap["Indocyber"], 100)
  fmt.Println(goMap)
}
