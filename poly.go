//package
//
//import "math"
//
//type Poly []float64
//
//func (p Poly) Val(x float64) float64 {
//	var ans float64
//	for i, ci := range p {
//		// ci == p[i]
//		ans += math.Pow(x, float64(i)) * ci
//	}
//	return ans
//}
//
//func Mul(p1, p2 Poly) Poly {
//	var ans Poly
//
//	// 1+2*x+3*x^2
//	// 4+5*x+6*x^2
//	// 1*4 + (1*5+2*4)*x + (1*6+2*5+3*4)*x^2 +...
//
//	for i, x := range p1 {
//		for j, y := range p2 {
//			for n := i + j; len(ans) <= n; {
//				ans = append(ans, 0)
//			}
//
//			ans[i+j] += x * y
//		}
//	}
//
//	return ans
//}
//
//func (p Poly) Int() Poly {
//	ans := make(Poly, len(p)+1)
//
//	for i, c := range p {
//		ans[i+1] = c / float64(i+1)
//	}
//
//	return ans
//}
