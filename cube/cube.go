package cube

type Cube struct {
	a int
}

func GetCube(a int) *Cube {
	return &Cube{
		a: a,
	}
}

func (c *Cube) Square() float32 {
	return 6.0 * float32(c.a) * float32(c.a)
}

func (c *Cube) Perimetr() float32 {
	return 12.0 * float32(c.a)
}

func (c *Cube) Volume() float32 {
	return float32(c.a) * float32(c.a) * float32(c.a)
}

func (_ *Cube) Name() string {
	return "куб"
}
