package main

import "fmt"

func main(){
  MainMenu()
}

func MainMenu() {
    fmt.Print("Welcome to Phoenix. Please choose the menu option you want to select:\n" +
        "\t 1. Book A Room\n" +
        "\t 2. See Room\n")

    fmt.Print("Input your option : ")

    var a int8

    _, err := fmt.Scan(&a)

    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    switch a {
    case 1:
        fmt.Println("You selected: Book A Room")
    case 2:
        fmt.Println("You selected: See Room")
    default:
        fmt.Println("Invalid option selected")
    }
}
