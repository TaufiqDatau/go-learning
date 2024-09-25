package main

import "fmt"

func main(){
  //Slice adalah potongan dari data array
  //Slice mirip dengan array yang membedakan adalah tipe datanya dinamis
  //artinya dapat menyimpan ukuran dari arraynya

  //detail tipe data slice ada 3 yaitu
  //Pointer penunjuk data pertama di array pada slice
  //length panjang dari slice
  //Capacity adalah kapasitas dari slicenya
  
  //membuat slice dari array
  var nilai = [5]int{1,2,3,4,5}
  
  var slice1 = nilai[0:3]
  fmt.Println(slice1)
  
  var slice2 = nilai[2:]
  fmt.Println(slice2)

  var slice3 = nilai[:4]
  fmt.Println(slice3)

  var slice4 = nilai[:]
  fmt.Println(slice4)

  //mengubah array akan mengubah slice juga
  nilai[3] = 100
  fmt.Println(slice4)
  //hal ini dikarenakan slice mereference array yang sama dengan array

  //untuk mendeklarasikan slice sama seperti array tapi tidak menuliskan lengthnya
  var namaSiswa []string

  fmt.Println(namaSiswa)


  namaSiswa2 := append(namaSiswa,"Aldy") //mencreate address baru untuk namaSiswa dan melepaskan yang lama 

  fmt.Println(namaSiswa2, "Ini adalah data setelah di append")
  fmt.Println(namaSiswa, "ini adalah array sebelum di append")


  //make(type, int length, int capacity)
  var newSlice = make([]string,3, 5) //declaring new slice

  fmt.Println(newSlice)
  
  newSlice = append(newSlice,"data1")
  referenceSlice := newSlice[0:1]
  referenceSlice[0]="test"
  fmt.Println(newSlice, "1")
  newSlice = append(newSlice, "new Data") // capacity = 5
  referenceSlice[0] ="baru test"
  fmt.Println(newSlice, "2")
  newSlice = append(newSlice, "newer data") // capacity = 10 
  referenceSlice[0] = "should in different memory"
  fmt.Println(newSlice,"3")

  fmt.Println(cap(newSlice))
  fmt.Println(len(newSlice))


  newSlice[0] = "data2"
  newSlice[1] = "data3"
  newSlice[2] = "data4"
  fmt.Println(newSlice)

}
