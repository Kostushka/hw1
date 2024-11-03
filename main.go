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

type Squarer interface {
	Square() float32
	Name() string
}

type Perimetrer interface {
	Perimetr() float32
	Name() string
}

type Volumer interface {
	Volume() float32
	Name() string
}

func getSquare(figures ...Squarer) {
	for _, v := range figures {
		fmt.Printf("Это %s с площадью %0.2f\n", v.Name(), v.Square())
	}
}

func getPerimetr(figures ...Perimetrer) {
	for _, v := range figures {
		fmt.Printf("Это %s с периметром %0.2f\n", v.Name(), v.Perimetr())
	}
}

func getVolume(figures ...Volumer) {
	for _, v := range figures {
		fmt.Printf("Это %s с объемом %0.2f\n", v.Name(), v.Volume())
	}	
}

func main() {
	square := square.GetSquare(5)
	rectangle := rectangle.GetRectangle(2,8)
	circle := circle.GetCircle(6)
	parallelepiped := parallelepiped.GetParallelepiped(3,6,9)
	ball := ball.GetBall(2)
	cube := cube.GetCube(8)
	getSquare(square, rectangle, circle, parallelepiped, ball, cube)
	fmt.Println()
	getPerimetr(square, rectangle, circle, parallelepiped, cube)
	fmt.Println()
	getVolume(parallelepiped, ball, cube)
}
