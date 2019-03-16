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

func Mult(count int, x []float64, y []float64) ([][]float64, []float64) {

	var matrix [][]float64
	for i := 0; i <= count; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j <= count; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}
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
