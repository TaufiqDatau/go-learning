package main

import "fmt"

type NoKTP string

func main(){
  var ktpKu NoKTP = "1111111"
  var contoh string = "222222"

  var contohKTP NoKTP = NoKTP(contoh)

  contohKTP.printNoKTP()
  fmt.Println(ktpKu)
}

func (no NoKTP) printNoKTP(){
  fmt.Println(no)
}
