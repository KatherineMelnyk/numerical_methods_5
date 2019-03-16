package main

func sequenceOfX(N int) []float64 {
	sequence := make([]float64, N)
	h := (B - A) / float64(N)
	for i := 0; i < N; i++ {
		sequence[i] = A + h*float64(i)
	}
	return sequence
}

func sequenceOfY(x []float64) []float64 {
	sequence := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		sequence[i] = Func(x[i])
	}
	return sequence
}

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

func Amatrix(N int) [][]float64 {
	h := (B - A) / float64(N)
	a := matrix(N-1, N-1)
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1; j++ {
			if i == j {
				a[i][j] = (2 * h) / 3
			} else if i+1 == j || j+1 == i {
				a[i][j] = h / 6
			} else {
				a[i][j] = 0
			}
		}
	}
	return a
}

func Hmatrix(N int) [][]float64 {
	h := (B - A) / float64(N)
	H := matrix(N+1, N-1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < N-1; j++ {
			if i == j || i+2 == j {
				H[i][j] = 1 / h
			} else if i+1 == j {
				H[i][j] = (-1) * (2 / h)
			} else {
				H[i][j] = 0
			}
		}
	}
	return H
}

func P(N int) [][]float64 {
	P := matrix(N, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				P[i][j] = 1
			}
		}
	}
	return P
}

//func cefM(N int) {
//	S := sequenceOfX(N)
//	Y := sequenceOfY(S)
//	MatrixP := P(N)
//	MatrixA := Amatrix(N)
//	MatrixH := Hmatrix(N)
//
//	vectorP := FromMattoVec(MatrixP)
//	vectorA := FromMattoVec(MatrixA)
//	vectorH := FromMattoVec(MatrixH)
//
//	F := mat.NewDense(len(Y), 1, Y)
//	P := mat.NewDense(len(MatrixP), len(MatrixP[0]), vectorP)
//	A := mat.NewDense(len(MatrixA), len(MatrixA), vectorA)
//	H := mat.NewDense(len(MatrixH), len(MatrixH), vectorH)
//	TransponseH := H.T()
//	InverseP := mat.NewDense(len(MatrixP), len(MatrixP), nil)
//	InverseP.Inverse(P)
//
//	Mul := mat.NewDense(N-1, N-1, nil)
//	Mul.Product(H, InverseP, TransponseH)
//	Add := mat.NewDense(N-1, N-1, nil)
//	Add.Add(Mul, A)
//	addF := mat.NewDense()
//}
