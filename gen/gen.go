package gen

import (
	"math"

	perlin "github.com/aquilax/go-perlin"
)

type Perlin struct {
	alpha float64
	beta  float64
	n     int
	seed  int64
}

func LowEdge2(x, y, width, height int) float64 {

	return 1.
}

func LowEdge(x, y, width, height int) float64 {

	shift := 2
	f := func(v, max int) float64 {
		return (math.Tanh(float64(v-shift)) + 1) * (-math.Tanh(float64(v-(max-shift))) + 1) / 4
	}

	// fmt.Println(f(x, width) * f(y, height))
	return f(x, width) * f(y, height)
	// return 1.z
}

func (per Perlin) Generate(width, height int) [][]float64 {

	p := perlin.NewPerlin(per.alpha, per.beta, per.n, per.seed)

	var min float64 = 100000
	bs2 := make([][]float64, height)
	for y := 0; y < height; y++ {
		row2 := make([]float64, width)
		for x := 0; x < width; x++ {
			v := p.Noise2D(float64(x)/10, float64(y)/10)
			row2[x] = v

			if v < min {
				min = v
			}
		}
		bs2[y] = row2
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Shift curve to remove negative values, then muliply by change functions
			bs2[y][x] = LowEdge(x, y, width, height) * (bs2[y][x] - min)
		}
	}

	return bs2
}

func GenerateParam(alpha, beta float64, seed int64, n, width, height int) [][]float64 {
	per := Perlin{
		alpha: alpha,
		beta:  beta,
		n:     n,
		seed:  seed,
	}
	return per.Generate(width, height)
}
