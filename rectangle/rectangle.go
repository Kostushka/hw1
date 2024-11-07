package rectangle

type Rectangle struct {
	a int
	b int
}

func NewRectangle(a,b int) *Rectangle {
	return &Rectangle{
		a: a,
		b: b,
	}
}

func (r *Rectangle) Square() float32 {
	return float32(r.a) * float32(r.b)
}

func (r *Rectangle) Perimetr() float32 {
	return float32(2) * (float32(r.a) + float32(r.b))
}

func (_ *Rectangle) Name() string {
	return "прямоугольник"
}
