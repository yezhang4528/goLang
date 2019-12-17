package main

import "fmt"

type Shape interface {
    getArea() float64
}

type Triangle struct {
    h float64
    b float64
}

type Square struct {
    sideLen float64
}

func (tri Triangle) getArea() float64 {
    return tri.h * tri.b * 0.5
}

func (sq Square) getArea() float64 {
    return sq.sideLen * sq.sideLen
}

func printArea(sh Shape) {
    fmt.Println(sh.getArea())
}

func main() {
    tTri := Triangle{1, 3}

    printArea(tTri)

    tSh := Square{5}

    printArea(tSh)
}
