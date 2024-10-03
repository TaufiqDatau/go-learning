package main

import "fmt"

type Address struct{
  city,province,nation string
}

func main(){
  Address1 := Address{"Nganjuk", "Jawa Timur", "Indonesia"}
  Address2 := Address1

  Address2.province = "Jawa Barat"
  fmt.Println(Address1, "Value address 1")
  fmt.Println(Address2, "Value Address 2 ")

  //untuk membuat pass by reference kita perlu mengassign pointer

  address3 := &Address1
  address3.city = "Surabaya"
  fmt.Println(Address1, "Value address 1 terupdate")
  fmt.Println(address3, "Merupakan nilai pointer ke address 1")
  fmt.Println(*address3)

  address4 := &*address3
  address4.province = "Jawa Barat"
  address4.city = "Subang"
  fmt.Println(address4, "Mengisi pointer dengan pointer")

  fmt.Println(Address1, "Data yang seharusnya berubah karena diganti dengan manipulasi address 4")
}
