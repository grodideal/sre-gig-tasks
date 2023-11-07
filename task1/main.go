package main

import "fmt"

func main() {
    // Define the array "Menu"
    var Menu []string
    
    // Append items to the Menu array
    Menu = append(Menu, "hamburger")
    Menu = append(Menu, "salad")
    
    // Iterate over the list and print for each item Food: <Food name>
    for _, food := range Menu {
        fmt.Println("Food:", food)
    }
}
