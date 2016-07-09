package main
import "fmt"

// Trying to create a struct which contains another struct as it's constituent.
// Something like a parent-child/nested relationship

type Base struct {
    v1 int
}

type Sub1 struct {
    b1, b2 Base
    v2 int
}

type Sub2 struct {
    Base // Anonymous
    v3 int
}

func main() {
    s1 := Sub1{Base{1}, Base{33}, 2}
    fmt.Println(s1)
}
