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

func scalMul(count int, x []float64, y []float64) [][]float64 {
	var matrix [][]float64
	for i := 0; i < count; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < len(x); j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			sum := 0.

			if j != len(x) {
				for t := 0; t < len(x); t++ {
					sum += y[t] * math.Pow(x[t], float64(j))
				}
				matrix[i][j] = sum
			} else {
				for k := 0; k < len(x); k++ {
					sum += math.Pow(x[k], float64(i)) * math.Pow(x[k], float64(j))
				}
				matrix[i][j] = sum
			}

		}
	}
	return matrix
}

//size := 10
//x := X(size)
//y := dataPoint(F, x)
//for i := 0; i < len(x); i++ {
//fmt.Printf("%.3f \n", x[i])
//}
//for i := 0; i < len(x); i++ {
//fmt.Printf("%.3f \n", y[i])
//}
