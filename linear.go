package matrix

// ProbabiltyAmpl calculates probability amplitudes
// for given state vector and the pair of orthonormal basis kets.
// The calculationis based on the following formula:
// v = x1*ket1 + x2*ket2
// [ket1, ket2] = a
// aT * v = [x1, x2]
// where:
// * v - is the given vector;
// * ket1 - is the 1st element in the kets slice;
// * ket2 - is the 2nd element in the kets slice;
// * a - is the matrix based of given kets;
// * aT - is the transpose of matrix "a";
// * x1 - is the probability amplitude for ket1;
// * x2 - is the probability amplitude for ket2.
func ProbabiltyAmpl(v Vector, ket1, ket2 Vector) (x1 int, x2 int) {
	a := DimM(ket1.Dim(), 2)
	ket1.ForEach(func(i, e int) bool {
		a.SetEntry(i, 0, e)
		return true
	})
	ket2.ForEach(func(i, e int) bool {
		a.SetEntry(i, 1, e)
		return true
	})

	result := a.Transpose().Multiply(M([]Vector{v}))

	return result.Entry(0, 0), result.Entry(1, 0)
}
