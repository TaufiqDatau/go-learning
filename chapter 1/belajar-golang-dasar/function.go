package main

import "fmt"

func printHello(){
  fmt.Println("Hello Gaes")
}

//Parameters
func printHelloName(name string){
  fmt.Println("Hello",name)
}

//Function with return value
func getAllChar(name string)[]string{
  var runeArr []string
  for _,char := range name {
    runeArr = append(runeArr, string(char))
  }

  return runeArr
}

//Variadic function 
func totalSum(ints ...int) int{
  result :=0
  for _,value := range ints {
    result += value
  }

  return result
}

func main(){
  printHello() // "Hello Gaes"
  printHelloName("Aqira") // "Hello Aqira"
  runeArr := getAllChar("Aqira")
  fmt.Println(runeArr)
  fmt.Println(totalSum(1,2,3,4,5,6,7,8))

  //If we already declared as array we can do this to a variadic function

  var listOfNumbers []int
  listOfNumbers = append(listOfNumbers, 10)
  listOfNumbers = append(listOfNumbers, 10)
  listOfNumbers = append(listOfNumbers, 10)
  listOfNumbers = append(listOfNumbers, 10)
  listOfNumbers = append(listOfNumbers, 10)

  fmt.Println(listOfNumbers)
  fmt.Println(totalSum(listOfNumbers...))
}
