package main

import "fmt"

// Point ...
type Point struct {
	X, Y int
}

// Circle ...
type Circle struct {
	Point
	Radius int
}

// Wheel ...
type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w = Wheel{Circle{Point{8, 8}, 5}, 20}

	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, //Nota: a vírgula no final e necessária aqui
	}

	fmt.Printf("%#v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)
}
