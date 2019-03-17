package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func Mult(count int, x []float64, y []float64) ([][]float64, []float64) {

	matrix := matrix(count+1, count+1)
	vector := make([]float64, count+1)

	for i := 0; i <= count; i++ {
		for j := 0; j <= count; j++ {
			for k := 0; k < len(x); k++ {
				matrix[i][j] += math.Pow(x[k], float64(j)) * math.Pow(x[k], float64(i))
			}
		}
		for t := 0; t < len(x); t++ {
			vector[i] += y[t] * math.Pow(x[t], float64(i))
		}
	}
	return matrix, vector
}

func PolinomApr(v []float64) func(float64) float64 {
	return func(value float64) float64 {
		res := 0.
		for i := 0; i < len(v); i++ {
			res += v[i] * math.Pow(value, float64(i))
		}
		return res
	}
}

func rho(x, y []float64, f func(float64) float64) float64 {
	n0 := 12
	m0 := 5
	sum := 0.
	for i := 0; i < len(x); i++ {
		sum += math.Pow(f(x[i])-y[i], 2)
	}
	return math.Sqrt(sum / float64(n0-m0))
}

func seqSigma(x, y []float64, from, to int) {
	for i := from; i < to; i++ {
		Ki, Li := Mult(i, x, y)
		ci := SolutionSystem(FromMattoVec(Ki), Li)
		Tablei := PolinomApr(ci)
		fmt.Printf("Rho_m (m=%v) : %f\n", i, rho(x, y, Tablei))
	}
}

func graphicSec() {
	ImageFunc := plotter.NewFunction(Func)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100

	x := sequenceOfX(25)
	y := sequenceOfY(x)

	seqSigma(x, y, 3, 13)

	K, L := Mult(13, x, y)
	c := SolutionSystem(FromMattoVec(K), L)
	Table := PolinomApr(c)
	AprFun3 := plotter.NewFunction(Table)
	AprFun3.Color = color.RGBA{R: 50, G: 88, B: 123, A: 111}
	AprFun3.Width = vg.Inch / 20
	AprFun3.Samples = 100

	pl2, _ := plot.New()
	pl2.X.Min, pl2.X.Max = 0, B
	pl2.Y.Min, pl2.Y.Max = -1, 1

	pl2.Add(ImageFunc)
	pl2.Add(AprFun3)

	pl2.Title.Text = "Approximation"
	pl2.Title.Font.Size = vg.Inch
	pl2.Legend.Font.Size = vg.Inch / 2
	pl2.Legend.XOffs = -vg.Inch
	pl2.Legend.YOffs = vg.Inch / 2
	pl2.Legend.Add("Function", ImageFunc)
	pl2.Legend.Add("Aproximation by Table", AprFun3)

	if err := pl2.Save(14*vg.Inch, 14*vg.Inch, "Task2.png"); err != nil {
		panic(err.Error())
	}

	f3 := Diverse(Func, Table)
	er3 := math.Sqrt(RectangleMethod(MulOf2(f3, f3), 0, B, 5000))
	fmt.Printf("Error of aproximation func by Table %v\n", er3)
}
