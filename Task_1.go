package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/gonum/mat"
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

func Diverse(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(x) - g(x)
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

func PolinomCheb(count, n int) func(float64) float64 {
	coef := TestMethod(count, n)

	return func(value float64) float64 {
		res := 0.
		for i, ci := range coef {
			if i%2 == 0 {
				res += ci * math.Cos(float64(i)*math.Acos((value-(B+A)/2)*(2/(B-A))))
			}
		}
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

func matrix_of_scalar_mults(count, n int) ([][]float64, []float64) {
	var matrix [][]float64
	for i := 0; i < count; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < count; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}

	vector := make([]float64, count)
	for l := 0; l < count; l++ {
		vector[l] = RectangleMethod(MulOf2(Func, func(x float64) float64 {
			return math.Pow(math.E, float64(l)*x)
		}), 0, math.Pi, n)
	}

	for k := 0; k < count; k++ {
		for t := 0; t < count; t++ {
			matrix[k][t] = RectangleMethod(MulOf2(func(x float64) float64 {
				return math.Pow(math.E, float64(t)*x)
			},
				func(x float64) float64 {
					return math.Pow(math.E, float64(k)*x)
				}), 0, math.Pi, n)
		}
	}

	return matrix, vector
}

func SolutionSystem(Vec, v []float64) []float64 {
	ScalarProd := mat.NewDense(len(v), len(v), Vec)
	ScalProdF := mat.NewDense(len(v), 1, v)
	var c mat.Dense
	c.Solve(ScalarProd, ScalProdF)
	coefficients := make([]float64, len(v))
	for i := 0; i < len(v); i++ {
		coefficients[i] = c.RawRowView(i)[0]
	}
	return coefficients
}

func Another(coef []float64) func(float64) float64 {
	return func(value float64) float64 {
		res := 0.
		for i := 0; i < len(coef); i++ {
			res += coef[i] * math.Pow(math.E, float64(i)*value)
		}
		return res
	}
}

func main() {
	ImageFunc := plotter.NewFunction(Func)
	ImageFunc.Color = color.RGBA{R: 209, G: 15, B: 15, A: 200}
	ImageFunc.Width = vg.Inch / 20
	ImageFunc.Samples = 100

	aprFun := PolinomCheb(15, 5000)
	AprFun := plotter.NewFunction(aprFun)
	AprFun.Color = color.RGBA{R: 223, G: 78, B: 90, A: 100}
	AprFun.Width = vg.Inch / 20
	AprFun.Samples = 100

	//fmt.Println(aprFun(-0.1), aprFun(1))
	//fmt.Println(Func(-0.1), Func(1))
	//fmt.Print("\n")

	C, v := matrix_of_scalar_mults(15, 7000)
	result := SolutionSystem(FromMattoVec(C), v)
	ExpSys := Another(result)
	AprFun2 := plotter.NewFunction(ExpSys)
	AprFun2.Color = color.RGBA{R: 23, G: 108, B: 200, A: 150}
	AprFun2.Width = vg.Inch / 20
	AprFun2.Samples = 100

	//fmt.Println(ExpSys(-0.1), ExpSys(1))
	//fmt.Println(Func(-0.1), Func(1))
	//fmt.Print("\n")

	pl, _ := plot.New()
	pl.X.Min, pl.X.Max = 0, B
	pl.Y.Min, pl.Y.Max = -1.5, 1

	pl.Add(ImageFunc)
	pl.Add(AprFun)
	pl.Add(AprFun2)

	pl.Title.Text = "Approximation"
	pl.Title.Font.Size = vg.Inch
	pl.Legend.Font.Size = vg.Inch / 2
	pl.Legend.XOffs = -vg.Inch
	pl.Legend.YOffs = vg.Inch / 2
	pl.Legend.Add("Function", ImageFunc)
	pl.Legend.Add("Aproximation by polinom Chebeshov", AprFun)
	pl.Legend.Add("Aproximation by e^ix", AprFun2)

	//graphicSec()

	if err := pl.Save(14*vg.Inch, 14*vg.Inch, "Task1.png"); err != nil {
		panic(err.Error())
	}

	f1 := Diverse(Func, aprFun)
	er1 := math.Sqrt(RectangleMethod(MulOf2(f1, f1), A, B, 5000))
	fmt.Printf("Error of aproximation func by polinom Cheb : %v\n", er1)

	f2 := Diverse(Func, ExpSys)
	er2 := math.Sqrt(RectangleMethod(MulOf2(f2, f2), 0, B, 5000))
	fmt.Printf("Error of aproximation func by system E^ix %v\n", er2)

	print_vector(m(3))
}
