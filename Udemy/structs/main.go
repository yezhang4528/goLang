package main

import "fmt"

type contactInfo struct {
    email   string
    zipCode int
}

type person struct {
    firstName   string
    lastName    string
    contactInfo 
}

func main() {
    /* // Declare and init variable alex with type person
    
    alex := person {
        firstName: "Alex",
        lastName: "Anderson",
    }

    // Second way of declare and update struct value
    var alex person

    alex.firstName = "Alex"
    alex.lastName = "Anderson"
    fmt.Println(alex)
    fmt.Printf("%+v\n", alex)    // Print struct value with field name
    */

    jim := person {
        firstName: "Jim",
        lastName: "Party",
        contactInfo:  contactInfo {
            email: "jimParty@gmail.com",
            zipCode: 94000,
        },
    }

//    jim.updateFirstName("Jimmy")
//    fmt.Println("In main...")

/*  // Traditional way of declaring and using pointer
    jimPtr := &jim
    jimPtr.updateFirstName("Jimmy")
*/

    jim.updateFirstName("Jimmy")        // Go shortway way of using pointer
    jim.print()
}

/* // Pass by value
func (p person) updateFirstName(newFirstName string) {
    p.firstName = newFirstName
    fmt.Println("In function updateFirstName")
    p.print()
}
*/

// Pass by reference
func (pPtr *person) updateFirstName(newFirstName string) {
    (*pPtr).firstName = newFirstName
}

func (p person) print() {
    fmt.Printf("%+v\n", p)
}
