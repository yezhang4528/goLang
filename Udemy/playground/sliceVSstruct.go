package main

import "fmt"

type person struct {
    firstName string
    lastName string
}

func main() {
    mySlice := []string{"Hi", "There", "How", "Are", "You"}

    updateSlice(mySlice)
    fmt.Println(mySlice)

    jim := person {
        firstName: "Jim",
        lastName: "Party",
    }

    fmt.Printf("Original person: %+v\n", jim)

    fmt.Println("Call by value")
    jim.updateStructByValue()
    fmt.Println(jim)

    fmt.Println("Call by reference")
    jim.updateStructByRef()
    fmt.Println(jim)
}

func updateSlice(s []string) {
    s[0] = "Bye"
}

func (p person) updateStructByValue() {
    p.firstName = "Jimmy"
}

func (p *person) updateStructByRef() {
    (*p).firstName = "Jimmy"
}
