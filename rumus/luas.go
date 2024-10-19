package rumus

type Hitung interface {
	Luas() float64
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return l.JariJari * l.JariJari * (22.0 / 7.0)
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

type PersegiPanjang struct {
	Panjang float64
	Lebar   float64
}

func (pp PersegiPanjang) Luas() float64 {
	return pp.Panjang * pp.Lebar
}

type Segitiga struct {
	Alas   float64
	Tinggi float64
}

func (s3 Segitiga) Luas() float64 {
	return s3.Alas * s3.Tinggi / 2
}

type Ketupat struct {
	Diagonal1 float64
	Diagonal2 float64
}

func (k Ketupat) Luas() float64 {
	return k.Diagonal1 * k.Diagonal2 / 2
}
