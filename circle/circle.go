package circle

type Circle struct {
	r int
}

func GetCircle(r int) *Circle {
	return &Circle{
		r: r,
	}
}

func (c *Circle) Square() float32 {
	return 3.1415926 * float32(c.r) * float32(c.r)
}

func (c *Circle) Perimetr() float32 {
	return float32(2) * float32(c.r) * 3.1415926
}
