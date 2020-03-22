package matrix

import "math"

// Vector interface defines the default contract for a vector.
type Vector interface {

	// Dim returns number of dimensions of this Vector.
	Dim() int

	// ForEach iterates through all entries of this Vector,
	// given iterate function receives index of an entry as the first argument
	// and value of an entry as the second,
	// return true to continue iteration.
	ForEach(func(int, int) bool)

	// Equals returns true if this Vector equals to given Vector v.
	Equals(v Vector) bool

	// MultiplyScalar performs multiplication of this Vector by a scalar "s" and returns result,
	// without changing properties of this Vector (no side effects).
	MultiplyScalar(s int) Vector

	// Add performs addition of this Vector with the "v" Vector and returns result,
	// without changing properties of this Vector (no side effects).
	Add(v Vector) Vector

	// Multiply performs multiplication of this Vector by given "v" Vector and returns result,
	// without changing properties of this Vector (no side effects).
	Multiply(v Vector) int

	// Len returns length of this vector.
	Len() float64

	// IsUnit returns true if this vector is a unit vector,
	// i.e. its bracket is 1.
	IsUnit() bool

	// IsOrthogonal returns true if this vector is a orthogonal
	// (perpendicular) to given Vector "v".
	IsOrthogonal(v Vector) bool
}

// vector - is a default implementation of a vector.
type vector struct {
	v []int
}

func newVector(v []int) *vector {
	res := &vector{v: make([]int, len(v))}
	copy(res.v, v)
	return res
}

// DimV creates and returns new instance of an n-dimensional Vector.
func DimV(n int) Vector {
	return newVector(make([]int, n))
}

// V creates and returns new instance of a Vector,
// with the coordinates based of given slice of scalars.
func V(v []int) Vector {
	return newVector(v)
}

func (v *vector) Dim() int {
	return len(v.v)
}

func (v *vector) ForEach(fn func(int, int) bool) {
	for i, s := range v.v {
		if !fn(i, s) {
			return
		}
	}
}

func (v *vector) Equals(vec Vector) bool {
	if v.Dim() != vec.Dim() {
		return false
	}
	ok := true
	vec.ForEach(func(i int, s int) bool {
		ok = v.v[i] == s
		return ok
	})
	return ok
}

func (v *vector) MultiplyScalar(s int) Vector {
	res := newVector(v.v)
	res.ForEach(func(i int, n int) bool {
		res.v[i] = n * s
		return true
	})
	return res
}

func (v *vector) Add(vec Vector) Vector {
	res := newVector(v.v)
	vec.ForEach(func(i int, n int) bool {
		res.v[i] += n
		return true
	})
	return res
}

func (v *vector) Multiply(vec Vector) int {
	var x int
	vec.ForEach(func(i int, n int) bool {
		x += v.v[i] * n
		return true
	})
	return x
}

func (v *vector) Len() float64 {
	return math.Sqrt(float64(v.Multiply(v)))
}

func (v *vector) IsUnit() bool {
	return v.Multiply(v) == 1
}

func (v *vector) IsOrthogonal(vec Vector) bool {
	return v.Multiply(vec) == 0
}
