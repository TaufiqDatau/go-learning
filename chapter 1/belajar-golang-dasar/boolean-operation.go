package main

import "fmt"

func main(){
  var nilaiAkhir = 90
  var nilaiAkhir2 = 30

  fmt.Println(string(decideGrade(int16(nilaiAkhir))))
  fmt.Println(string(decideGrade(int16(nilaiAkhir2))))
  
}

func decideGrade(score int16) rune{
  if score > 85 {
    return 'A'
  }

  if score > 75 {
    return 'B'
  }

  if score > 65 {
    return 'C'
  }

  return 'D'
}
