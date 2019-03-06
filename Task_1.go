package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const A = -math.Pi
const B = math.Pi
const w = 1

func MulOf2(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) * g(x)
	}
}

func MulOf3(f, g, k func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) * g(x) * k(x)
	}
}

func RectangleMethod(f func(float64) float64, a, b float64, n int) float64 {
	sum := 0.
	weight := (b - a) / float64(n)
	for i := 0; i < n; i++ {
		sum += weight * f(a+float64(i)*weight)
	}
	return sum
}

func Func(x float64) float64 {
	return math.Cos(w * math.Pow(x, 2))
}

func MyFunc(x float64) float64 {
	t := ((B-A)/2)*x + (B+A)/2
	return math.Cos(w * math.Pow(t, 2))
}

func weightCoef(x float64) float64 {
	return 1. / (math.Sqrt(1 - math.Pow(x, 2)))
}

func PolinomCheb(count, n int) func(float64) float64 {
	coef := TestMethod(count, n)

	return func(value float64) float64 {
		fmt.Println(value)
		res := 0.
		for i, ci := range coef {
			fmt.Println(i)
			if i%2 == 0 {
				res += ci * math.Cos(float64(i)*math.Acos((value-(B+A)/2)*(2/(B-A))))
			}
		}
		fmt.Println(res)
		return res
	}
}

func TestMethod(count, n int) []float64 {
	coef := make([]float64, count)

	for i := range coef {
		ii := RectangleMethod(func(t float64) float64 {
			return MyFunc(math.Cos(t)) * math.Cos(float64(i)*t)
		}, 0, math.Pi, n)

		if i == 0 {
			coef[i] = ii / math.Pi
		} else if i%2 == 0 {
			coef[i] = 2 * ii / math.Pi
		}
	}

	return coef
}

func main() {
	a := TestMethod(4, 5000)
	for i := 0; i < len(a); i++ {
		fmt.Printf("%f\n", a[i])
	}

	ImageFunc := plotter.NewFunction(Func)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100
	aprFun := PolinomCheb(4, 5000)

	fmt.Println(aprFun(-0.1), aprFun(1))
	fmt.Println(Func(-0.1), Func(1))

	AprFun := plotter.NewFunction(aprFun)
	AprFun.Color = color.RGBA{R: 223, G: 78, B: 90, A: 100}
	AprFun.Width = vg.Inch / 20
	AprFun.Samples = 100
	//f2 := App2(4, 5000)
	//fPol2 := plotter.NewFunction(f2)
	//fPol2.Color = color.RGBA{R: 23, G: 108, B: 200, A: 150}
	//fPol2.Width = vg.Inch / 20
	//fPol2.Samples = 200
	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = A, B
	pl.Y.Min, pl.Y.Max = -math.Pi, math.Pi
	pl.Add(AprFun)
	pl.Add(ImageFunc)
	//pl.Add(fPol2)
	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	//pl.Legend.Add("Pol1", fPol1)
	//pl.Legend.Add("Pol2", fPol2)
	//pl.Legend.Add("Pol3", fPol3)

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "pol.png"); err != nil {
		panic(err.Error())

	}
	//er := math.Pow(Int(Comp2(Comp1(f, f1), Comp1(f, f1)), A, B, 10000), 1/2)
	//fmt.Printf("%v\n", er)

}
