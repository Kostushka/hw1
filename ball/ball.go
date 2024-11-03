package ball

type Ball struct {
	r int
}

func GetBall(r int) *Ball {
	return &Ball{
		r: r,
	}
}

func (b *Ball) Square() float32 {
	return 4.0 * 3.1415926 * float32(b.r)
}

func (b *Ball) Volume() float32 {
	return 4/3 * 3.1415926 * float32(b.r) * float32(b.r) * float32(b.r)
}

func (_ *Ball) Name() string {
	return "шар"
}
