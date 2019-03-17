package main

import "fmt"

func matrix(c, r int) [][]float64 {
	var matrix [][]float64
	for i := 0; i < r; i++ {
		matrix = append(matrix, []float64{})
		for j := 0; j < c; j++ {
			matrix[i] = append(matrix[i], 0.)
		}
	}
	return matrix
}

func FromMattoVec(matrix [][]float64) []float64 {
	Size := len(matrix) * len(matrix[0])
	vector := make([]float64, Size)
	t := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			vector[t] = matrix[i][j]
			t = t + 1
		}
	}
	return vector
}

func print_matrix(matrix [][]float64) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%f ", matrix[i][j])
		}
		fmt.Print("\n")
	}
}

func print_vector(vector []float64) {
	for j := 0; j < len(vector); j++ {
		fmt.Printf("%f ", vector[j])
	}
	fmt.Print("\n")
}
