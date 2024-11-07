package main

import (
	"fmt"
	"homework/square"
	"homework/rectangle"
	"homework/circle"
	"homework/parallelepiped"
	"homework/ball"
	"homework/cube"
)

type Figure interface {
	Name() string
}

type Squarer interface {
	Figure
	Square() float32
}

type Perimetrer interface {
	Figure
	Perimetr() float32
}

type Volumer interface {
	Figure
	Volume() float32
}

func filterSquare(figures []Figure) []Squarer {
	relevants := []Squarer{}
	for _, v := range figures {
		if res, ok := v.(Squarer); ok {
			relevants = append(relevants, res)
		}
	}
	return relevants
}

func filterPerimetr(figures []Figure) []Perimetrer {
	relevants := []Perimetrer{}
	for _, v := range figures {
		if res, ok := v.(Perimetrer); ok {
			relevants = append(relevants, res)
		}
	}
	return relevants
}

func filterVolume(figures []Figure) []Volumer {
	relevants := []Volumer{}
	for _, v := range figures {
		if res, ok := v.(Volumer); ok {
			relevants = append(relevants, res)
		}
	}
	return relevants
}

func printSquare(figures ...Squarer) {
	for _, v := range figures {
		fmt.Printf("Это %s с площадью %0.2f\n", v.Name(), v.Square())
	}
}

func printPerimetr(figures ...Perimetrer) {
	for _, v := range figures {
		fmt.Printf("Это %s с периметром %0.2f\n", v.Name(), v.Perimetr())
	}
}

func printVolume(figures ...Volumer) {
	for _, v := range figures {
		fmt.Printf("Это %s с объемом %0.2f\n", v.Name(), v.Volume())
	}	
}

func main() {
	figures := []Figure{
		square.NewSquare(5),
		rectangle.NewRectangle(2,8),
		circle.NewCircle(6),
		parallelepiped.NewParallelepiped(3,6,9),
		ball.NewBall(2),
		cube.NewCube(8),
	}

	printSquare(filterSquare(figures)...)
	fmt.Println()
	printPerimetr(filterPerimetr(figures)...)
	fmt.Println()
	printVolume(filterVolume(figures)...)
}
