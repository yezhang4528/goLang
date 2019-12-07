package main

import "fmt"

func main() {
    /* // 1st way of declaring a map
    colors := map[string]string{
        "red": "#ff0000",
        "green": "#7b7534",
        "white": "#ffffff",
    }
    */

    /* // 2nd way of declaring a map
    var colors map[string]string
    //colors["red"] = "#ff00000"      // Got error: panic: assignment to entry in nil map
    */

    /* // 3rd way of declaring a map
    colors := make(map[string]string)
    colors["red"] = "#ff00000"      // This works!
    fmt.Println(colors)

    // Delete an element in map
    delete(colors, "red")
    */
    
    colors := map[string]string{
        "red": "#ff0000",
        "green": "#7b7534",
        "white": "#ffffff",
    }
    colors["yellow"] = "#4ff45e"    // This works!

    printMap(colors)

    updateMap(colors)
    fmt.Println("After updateMap")
    printMap(colors)
}

func printMap(c map[string]string) {
    for color, hex := range c {
        fmt.Println("Hex code for color ", color, " is ", hex)
    }
}

func updateMap(c map[string]string) {
    c["red"] = "#000000"
}
