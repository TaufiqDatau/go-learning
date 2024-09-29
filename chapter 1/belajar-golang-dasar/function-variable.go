package main

import "fmt"

type stringParam func(string)

func sayHello(name string){
  fmt.Println("Hello", name)
}

func filterSlice(numberArr []int32, nilaiMinimum int8, filterCondition func(int32, int8)bool) []int32{
  result := make([]int32, 0, 10)
  for _,value := range numberArr {
    if filterCondition(value, nilaiMinimum){
      result = append(result, value)
    }
  }
  return result
} 

func main(){
  var containFunction stringParam = sayHello
  containFunction("Jiya")
  nilaiSiswa := []int32{85,43,78,41,80,42,65}

  filter := func(val int32, condition int8)bool{
    if val >= int32(condition){
      return true
    }
    return false
  }

  nilaiSiswa = filterSlice(nilaiSiswa, 80, filter)

  fmt.Println(nilaiSiswa, "Menggunakan anonymous function dengan filter ")
}
