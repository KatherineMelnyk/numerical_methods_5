package main

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

func sequenceOfX(N int) []float64 {
	sequence := make([]float64, N+1)
	h := (math.Pi - 0.) / float64(N)
	for i := 0; i < len(sequence); i++ {
		sequence[i] = 0 + h*float64(i)
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

func Amatrix(N int) [][]float64 {
	h := (B - 0.) / float64(N)
	a := matrix(N-1, N-1)
	for i := 0; i < N-1; i++ {
		for j := 0; j < N-1; j++ {
			if i == j {
				if i == 0 {
					a[i][j+1] = 1 / h
				} else if i == N-2 {
					a[i][j-1] = 1 / h
				} else {
					a[i][j-1] = 1 / h
					a[i][j+1] = 1 / h
				}
				a[i][j] = (2 * h) / 3
			}
		}
	}
	return a
}

func Hmatrix(N int) [][]float64 {
	h := (B - 0) / float64(N)
	H := matrix(N+1, N-1)
	for i := 0; i < N-1; i++ {
		for j := 0; j < N+1; j++ {
			if i+1 == j {
				H[i][j] = (-1) * (2 / h)
				H[i][j-1] = 1 / h
				H[i][j+1] = 1 / h
			}
		}
	}
	return H
}

func P(N int) [][]float64 {
	P := matrix(N+1, N+1)
	for i := 0; i < N+1; i++ {
		for j := 0; j < N+1; j++ {
			if i == j {
				P[i][j] = 1
			}
		}
	}
	return P
}

func m(N int) []float64 {

	X := sequenceOfX(N)
	Y := sequenceOfY(X)
	MatA := Amatrix(N)
	MatH := Hmatrix(N)
	MAtP := P(N)

	vecH := FromMattoVec(MatH)
	vecA := FromMattoVec(MatA)
	vecP := FromMattoVec(MAtP)

	F := mat.NewDense(len(Y), 1, Y)
	H := mat.NewDense(len(MatH), len(MatH[0]), vecH)
	P := mat.NewDense(len(MAtP), len(MAtP[0]), vecP)
	A := mat.NewDense(len(MatA), len(MatA[0]), vecA)

	Right := mat.NewDense(N-1, 1, nil)
	Right.Product(H, F)

	InvP := mat.NewDense(len(MAtP), len(MAtP), nil)
	InvP.Inverse(P)

	Mul := mat.NewDense(N-1, N-1, nil)
	Mul.Product(H, InvP, H.T())

	Left := mat.NewDense(N-1, N-1, nil)
	Left.Add(Mul, A)

	var Res mat.Dense
	Res.Solve(Left, Right)
	m := make([]float64, N-1)
	for i := 0; i < len(m); i++ {
		m[i] = Res.RawRowView(i)[0]
	}
	return m
}

func mu(m []float64) []float64 {
	N := len(m)

	X := sequenceOfX(N)
	Y := sequenceOfY(X)
	MatH := Hmatrix(N)
	MAtP := P(N)

	vecH := FromMattoVec(MatH)
	vecP := FromMattoVec(MAtP)

	F := mat.NewDense(len(Y), 1, Y)
	M := mat.NewDense(len(m), 1, m)
	H := mat.NewDense(len(MatH), len(MatH[0]), vecH)
	P := mat.NewDense(len(MAtP), len(MAtP[0]), vecP)

	InvP := mat.NewDense(len(MAtP), len(MAtP), nil)
	InvP.Inverse(P)

	Mul := mat.NewDense(len(Y), 1, nil)
	Mul.Product(InvP, H.T(), M)

	var Res mat.Dense
	Res.Sub(F, Mul)

	r := make([]float64, len(Y))
	for i := 0; i < len(r); i++ {
		r[i] = Res.RawRowView(i)[0]
	}
	return m
}

func seqS(m, mu []float64, N int) []func(float64) float64 {
	h := B / float64(N)
	X := sequenceOfX(N)
	S := make([]func(float64) float64, N)
	for i := 0; i < N-1; i++ {
		S[i] = func(x float64) float64 {
			return m[i]*math.Pow(X[i+1]-x, 3)/6*h +
				m[i+1]*math.Pow(x-X[i], 3)/6*h +
				(mu[i]-m[i]*math.Pow(h, 2)/6)*(3-X[i+1])/h +
				(mu[i+1] - m[i+1]*math.Pow(h, 2)/6)
		}
	}
	return S
}
