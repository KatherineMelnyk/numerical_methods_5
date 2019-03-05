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

//const A = 1
//const B = 4
//const w = 0.5

func Comp2(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) * g(x)
	}
}

func Comp3(f, g, k func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) * g(x) * k(x)
	}
}

func Int(f func(float64) float64, a, b float64, n int) float64 {
	var sum float64
	h := b - a/float64(n)
	for i := 0; i <= n; i++ {
		rectangle := h * F(a+float64(i)/h)
		sum += rectangle
	}
	return sum
}

func F(x float64) float64 {
	t := (2*x - (B + A)) / (B - A)
	//return 2.*math.Pow(t, 2) + math.Sin(w*t)
	k := math.Pow(math.Pi*t, 2)
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
			if i == 0 {
				C[i] = Int(Comp2(F, rho), A, B, n)
			} else if i%2 == 0 {
				C[i] = Int(Comp3(
					func(x float64) float64 {
						t := (2*x - float64(B+A)) / float64(B-A)
						return math.Cos(float64(i) * math.Acos(t))
					}, F, rho), A, B, n)
			}
		}
		res := 0.
		for i := 0; i < len(C); i++ {
			res += C[i] * math.Pow(value, float64(i))
		}
		return res
	}
}

//func App2(count, n int) func(float64) float64 {
//	return func(value float64) float64 {
//		C := make([]float64, count)
//		for i := 0; i < count; i++ {
//			C[i] = Int(Comp2(F, func(x float64) float64 {
//				return math.Pow(math.E, float64(i)*x)
//			}), A, B, n)
//		}
//		res := 0.
//		for i := 0; i < len(C); i++ {
//			res += C[i] * math.Pow(value, float64(i))
//		}
//		return res
//	}
//}

//func TTT(value float64) float64 {
//	res := 0.
//	for i := 0; i < 4; i++ {
//		res += 31.9 * math.Pow(value, float64(i))
//	}
//	return res
//}

func main() {
	fPol3 := plotter.NewFunction(F)
	fPol3.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	fPol3.Width = vg.Inch / 20
	fPol3.Samples = 200
	f1 := App(4, 5000)
	fPol1 := plotter.NewFunction(f1)
	fPol1.Color = color.RGBA{R: 223, G: 78, B: 90, A: 100}
	fPol1.Width = vg.Inch / 20
	fPol1.Samples = 200
	//f2 := App2(4, 5000)
	//fPol2 := plotter.NewFunction(f2)
	//fPol2.Color = color.RGBA{R: 23, G: 108, B: 200, A: 150}
	//fPol2.Width = vg.Inch / 20
	//fPol2.Samples = 200
	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A+1, B+4
	pl.Y.Min, pl.Y.Max = -5, 35
	pl.Add(fPol1)
	pl.Add(fPol3)
	//pl.Add(fPol2)
	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Pol1", fPol1)
	//pl.Legend.Add("Pol2", fPol2)
	pl.Legend.Add("Pol3", fPol3)
	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())

	}
}
