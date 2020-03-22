package matrix

// Matrix interface defines the default contract for a matrix.
type Matrix interface {

	// MultiplyScalar performs multiplication of this Matrix by a scalar "s" and returns result,
	// without changing properties of this Matrix (no side effects).
	MultiplyScalar(s int) Matrix

	// Add performs addition of the m Matrix to this Matrix and returns result,
	// without changing properties of this Matrix (no side effects).
	Add(m Matrix) Matrix

	// Multiply performs multiplication of this Matrix by given "m" Matrix and returns result,
	// without changing properties of this Matrix (no side effects).
	Multiply(m Matrix) Matrix

	// Transpose returns transposed version of this Matrix,
	// without changing properties of this Matrix (no side effects).
	Transpose() Matrix

	// IsIdentity returns true in case if this Matrix is an identity matrix.
	IsIdentity() bool
}

// matrix - is default implementation of a matrix.
type matrix struct {
	bras []Vector
}
