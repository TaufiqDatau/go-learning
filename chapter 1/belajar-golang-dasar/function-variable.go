package main

import "fmt"

type stringParam func(string)
type filterFunction func(val int32, condition int8)bool

func sayHello(name string){
  fmt.Println("Hello", name)
}

func filterSlice(numberArr []int32, nilaiMinimum int8, filterCondition filterFunction) []int32{
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

  var filter filterFunction = func(val int32, condition int8) bool {
    return val >= int32(condition)
  }

  
  nilaiSiswa = filterSlice(nilaiSiswa, 80, filter)

  fmt.Println(nilaiSiswa, "Menggunakan anonymous function dengan filter ")
}
