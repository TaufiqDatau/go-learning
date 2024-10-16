package main

import (
	"errors"
	"fmt"
)


type vectorInt []int16

func main(){
  var listValue vectorInt = vectorInt{}

  sumValue,err := sumVector(listValue)

  if err != nil {
    fmt.Println(err.Error())
    return
  }

  fmt.Println("The total sum of the vector is %s", sumValue)

}

func sumVector(vec vectorInt) (int16, error){
  var result int16

  if len(vec) == 0 {
    return 0, errors.New("Length is zero")
  }

  for _,value := range vec{
    result +=  value
  }

  return result, nil

} 
