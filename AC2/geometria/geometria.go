package geometria

type Retangulo struct {
	Base float64
	Altura float64
}

func Area(r Retangulo) float64 {
	return r.Base * r.Altura
}

func Perimetro(r Retangulo) float64 {
	return 2 * (r.Base + r.Altura)
}
