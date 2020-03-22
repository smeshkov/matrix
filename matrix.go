package matrix

// Matrix interface defines the default contract for a matrix.
type Matrix interface {

	// MultiplyReal performs multiplication of this Matrix by a real number "r" and returns result,
	// without changing properties of this Matrix (no side effects).
	MultiplyReal(r int) Matrix

	// MultiplyComplex performs multiplication of this Matrix by a complex number "c" and returns result,
	// without changing properties of this Matrix (no side effects).
	MultiplyComplex(c float64) Matrix

	// Add performs addition of the m Matrix to this Matrix and returns result,
	// without changing properties of this Matrix (no side effects).
	Add(m Matrix) Matrix

	// Multiply performs multiplication of this Matrix by given "m" Matrix and returns result,
	// without changing properties of this Matrix (no side effects).
	Multiply(m Matrix) Matrix

	// Transpose returns transposed version of this Matrix, 
	// without changing properties of this Matrix (no side effects).
	Transpose() Matrix
}

// M - is a default implementation of matrix.
type M struct {}