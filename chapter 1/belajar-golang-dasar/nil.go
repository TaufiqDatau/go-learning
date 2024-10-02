package main

import "fmt"

func NewMap(name string) map[string]string {
  if name == ""{
    return nil
  }else{
    return map[string]string{
      "name": name,
    }
  }
}


func main(){
  //data := NewMap("")
  data := NewMap("TORIKARA")
  if data == nil {
    fmt.Println("Datanya kosong");
  }else{
    fmt.Println(data)
  }

}
