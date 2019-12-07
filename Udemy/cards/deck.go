package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "strings"
    "time"
)

// Define a new type deck,
// which is a slice of strings

type deck []string

func newDeck() deck {
    cards := deck{}

    cardSuits := []string {"Spades", "Diamond", "Heart", "Club"}
    cardValues := []string {"Ace", "Two", "Three", "Four"}

    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value+" of "+ suit)
        }
    }
    return cards
}

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func deal(d deck, handsNum int) (deck, deck) {
    return d[:handsNum], d[handsNum:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0660)
}

func getDeckFromFile(filename string) deck {
    deckByte, err := ioutil.ReadFile(filename)

    if err != nil {
        // Error reading file, log error and exit
        fmt.Println("Error: ", err)
        os.Exit(1)
    }

    s := strings.Split(string(deckByte), ",")
    return deck(s)
}

func (d deck) shuffle() {
    // Generate real random number, get seed based on current Unix time in nanosecond (UInt64)
    src := rand.NewSource(time.Now().UnixNano())  // returns Source generated with the seed
    r := rand.New(src)       // returns *Rand
    for i := range d {
        newIndex := r.Intn(len(d) - 1)
        d[i], d[newIndex] = d[newIndex], d[i]
    }
}
