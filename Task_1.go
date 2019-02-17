package main

import (
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const A = -math.Pi
const B = math.Pi
const w = 0.5

func Comp3(f, g, k func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) * g(x) * k(x)
	}
}

func Int(f func(float64) float64, a, b float64, n int) float64 {
	var sum float64
	for i := 0; i <= n; i++ {
		rectangle := F(a + (b-a)*float64(i)/float64(n))
		sum += rectangle
	}
	return sum
}

func F(x float64) float64 {
	t := (2*x - (B + A)) / (B - A)
	k := math.Pow(t, 2)
	return math.Cos(w * k)
}

func rho(x float64) float64 {
	g := math.Pow(x, 2)
	return math.Sqrt(1 - g)
}

func App(count, n int) func(float64) float64 {
	return func(value float64) float64 {
		C := make([]float64, count)
		for i := 0; i < count; i++ {
			C[i] = Int(Comp3(
				func(x float64) float64 {
					t := (2*x - float64(B+A)) / float64(B-A)
					return math.Cos(float64(i) * math.Acos(t))
				},
				F, rho), A, B, n)
		}
		res := 0.
		for i := 0; i < len(C); i++ {
			res += C[i] * math.Pow(value, float64(i))
		}
		return res
	}
}

func main() {
	f1 := App(4, 1000)
	fPol1 := plotter.NewFunction(f1)
	fPol1.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	fPol1.Width = vg.Inch / 20
	fPol1.Samples = 200
	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A-0.5, B+0.5
	pl.Y.Min, pl.Y.Max = -2, 2
	pl.Add(fPol1)
	pl.Title.Text = "Interpolation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Pol", fPol1)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())

	}
}
