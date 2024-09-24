package main

import "fmt"

func main(){
  //Array only have a fixed size 
  //After declaration you can't add more items inside the array
  //you can only override the existing value

  var names [3]string
  names[0] = "Taufiqurrahman"
  names[1] = "Hafiidh"
  names[2] = "Datau"

  fmt.Println(names[0])
  fmt.Println(names[1])
  fmt.Println(names[2])

  var grades = []string  {
    "A",
    "B",
    "C",
    "D",
  }

  fmt.Println(grades)

  //len berguna untuk menghitung panjang dari array nya
  fmt.Println(len(grades))
  

  //mendapatkand data array[index]
  var grade = grades[1]
  fmt.Println(grade)

  //mengubah data array[index] = value
  grades[3] = "X"
}
