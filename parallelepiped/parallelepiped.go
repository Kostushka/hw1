package parallelepiped

type Parallelepiped struct {
	a int
	b int
	c int
}

func NewParallelepiped(a,b,c int) *Parallelepiped {
	return &Parallelepiped{
		a: a,
		b: b,
		c: c,
	}
}

func (p *Parallelepiped) Square() float32 {
	return 2.0 *(float32(p.a)*float32(p.b) + float32(p.b)*float32(p.c) + float32(p.a)*float32(p.c))
}

func (p *Parallelepiped) Perimetr() float32 {
	return 4.0 * (float32(p.a)+float32(p.b)+float32(p.c))
}

func (p *Parallelepiped) Volume() float32 {
	return float32(p.a) * float32(p.b) * float32(p.c)
}

func (_ *Parallelepiped) Name() string {
	return "параллелепипед"
}
