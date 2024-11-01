package square

type Square struct {
	a int
}

func GetSquare(a int) *Square {
	return &Square{
		a: a,
	}
}

func (s *Square) Square() float32 {
	return float32(s.a) * float32(s.a)
}

func (s *Square) Perimetr() float32 {
	return 4.0 * float32(s.a)
}
