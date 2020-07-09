package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// func (d Direction) String() string {
// 	return [...]string{"North", "East", "South", "West"}[d]
// }


type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo":Math{2,3},
}

var m2 = map[string]*Math{
	"foo":&Math{2,3,},
}

func main() {
	// fmt.Println(South)
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)

	m2["foo"].x = 4
	fmt.Println(m2["foo"].x)
	fmt.Printf("%#v", m["foo"])
}