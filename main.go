package main

import (
	"fmt"
	"homework/square"
	"homework/rectangle"
	"homework/circle"
	"homework/parallelepiped"
	"homework/ball"
	"homework/cube"
	"time"
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

func filter[T any](figures []Figure) []T {
	relevants := []T{}
	for _, v := range figures {
		if res, ok := v.(T); ok {
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
	ch := make(chan []Figure)
	
	go func() {
		var figures []Figure
		a,b,c := 1,5,7
		for {
			figures = []Figure{
				square.NewSquare(a),
				rectangle.NewRectangle(b,c),
				circle.NewCircle(c),
				parallelepiped.NewParallelepiped(a,b,c),
				ball.NewBall(b),
				cube.NewCube(c),
			}
			time.Sleep(2 * time.Second)
			ch <-figures
			a++
			b *= 2
			c += 7
		}		
	}()

	go func() {
		for {
			figures := <-ch

			fmt.Println("_______________")
			printSquare(filter[Squarer](figures)...)
			fmt.Println()
			printPerimetr(filter[Perimetrer](figures)...)
			fmt.Println()
			printVolume(filter[Volumer](figures)...)
			fmt.Println("_______________")
		}
	}()	
	fmt.Scanf("%s", new(string))
}
