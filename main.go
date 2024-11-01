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
}

type Perimetrer interface {
	Perimetr() float32
}

type Volumer interface {
	Volume() float32
}

func getSquare(figures ...Squarer) {
	for _, v := range figures {
		switch res := v.(type) {
			case *square.Square:
				fmt.Printf("Это квадрат с площадью %0.2f\n", res.Square())
			case *rectangle.Rectangle:
				fmt.Printf("Это прямоугольник с площадью %0.2f\n", res.Square())
			case *circle.Circle:
				fmt.Printf("Это круг с площадью %0.2f\n", res.Square())
			case *parallelepiped.Parallelepiped:
				fmt.Printf("Это параллелепипед с площадью %0.2f\n", res.Square())
			case *ball.Ball:
				fmt.Printf("Это шар с площадью %0.2f\n", res.Square())	
			case *cube.Cube:
				fmt.Printf("Это куб с площадью %0.2f\n", res.Square())
		}
	}
}

func getPerimetr(figures ...Perimetrer) {
	for _, v := range figures {
		switch res := v.(type) {
			case *square.Square:
				fmt.Printf("Это квадрат с периметром %0.2f\n", res.Perimetr())
			case *rectangle.Rectangle:
				fmt.Printf("Это прямоугольник с периметром %0.2f\n", res.Perimetr())
			case *circle.Circle:
				fmt.Printf("Это круг с периметром %0.2f\n", res.Perimetr())
			case *parallelepiped.Parallelepiped:
				fmt.Printf("Это параллелепипед с периметром %0.2f\n", res.Perimetr())	
			case *cube.Cube:
				fmt.Printf("Это куб с периметром %0.2f\n", res.Perimetr())
		}
	}
}

func getVolume(figures ...Volumer) {
	for _, v := range figures {
		switch res := v.(type) {
			case *parallelepiped.Parallelepiped:
				fmt.Printf("Это параллелепипед с объемом %0.2f\n", res.Volume())
			case *ball.Ball:
				fmt.Printf("Это шар с объемом %0.2f\n", res.Volume())
			case *cube.Cube:
				fmt.Printf("Это куб с объемом %0.2f\n", res.Volume())
		}
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
