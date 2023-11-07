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
    // Define an array of 5 items
    items := [5]string{"apple", "banana", "orange", "grape", "pear"}

    // Iterate over the array and print each item with its index
    for index, item := range items {
        fmt.Printf("This is %s and its index in the array is %d\n", item, index)
    }
    
}
