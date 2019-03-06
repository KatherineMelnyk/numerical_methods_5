package main

import (
	"math"
)

func X(size int) []float64 {
	x := make([]float64, size)
	for i := 0; i < size; i++ {
		x[i] = (float64(i) / 10.) * math.Pow(-1, float64(i))
	}
	return x
}

func dataPoint(f func(float64) float64, x []float64) []float64 {
	y := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		y[i] = f(x[i])
	}
	return y
}

func Mult(count int, x []float64, y []float64) ([]float64, []float64) {

	var matrix [][]float64
	for i := 0; i <= count; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j <= count; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}

	var vector, coeficients []float64
	for i := 0; i <= count; i++ {
		vector = append(vector, 0.)
	}
	for t := 0; t <= count*3; t++ {
		coeficients = append(coeficients, 0.)
	}

	for i := 0; i <= count; i++ {
		for j := 0; j <= count; j++ {
			for k := 0; k < len(x); k++ {
				matrix[i][j] += math.Pow(x[k], float64(j)) * math.Pow(x[k], float64(i))
			}
			matrix[i][j] /= 6
		}
		for t := 0; t < len(x); t++ {
			vector[i] += y[t] * math.Pow(x[t], float64(i))
		}
		vector[i] /= 6
	}
	index := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			coeficients[index] = matrix[i][j]
			index++
		}
	}
	return coeficients, vector
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

//x := []float64{0., 0.2, 0.4, 0.6, 0.8, 1.}
//y := []float64{0.2857, 0.165, 0.0468, -0.0721, -0.1938, -0.3171}
//K, L := Mult(1, x, y)
//N := mat.NewDense(2, 2, K)
//M := mat.NewDense(2, 1, L)
//var c mat.Dense
//c.Solve(N, M)
