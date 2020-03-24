package matrix

// Matrix interface defines the default contract for a matrix.
type Matrix interface {

	// Dim returns number of rows (1st value) and columns (2nd value) of this Matrix.
	Dim() (int, int)

	// Entry returns Matrix entry value in the position of given row and col.
	Entry(row, col int) int

	// SetEntry sets given entry value to the position of given row and col.
	SetEntry(row, col, entry int)

	// ForEach iterates through all entries of this Matrix,
	// given function receives row-index (1st argument), column-index (2nd argument)
	// and value of an entry (3rd argument), return true to continue iteration.
	ForEach(func(int, int, int) bool)

	// Equals returns true if this Matrix equals to given Matrix m.
	Equals(m Matrix) bool

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
	// rs - number of rows
	// cs - number of columns
	rs, cs int
	// es - is a raw slice of entries
	es []int
}

func (mx *matrix) index(row, col int) int {
	return mx.cs*row + col
}

func newMatrix(rows, cols int, e []int) *matrix {
	res := &matrix{rs: rows, cs: cols, es: make([]int, len(e))}
	copy(res.es, e)
	return res
}

// DimM creates and returns new instance of an m*n-dimensional Matrix.
func DimM(m, n int) Matrix {
	return newMatrix(m, n, make([]int, m*n))
}

// IdentityM creates and returns new instance of an n*n-dimensional identity Matrix.
func IdentityM(n int) Matrix {
	es := make([]int, n*n)
	for i := range es {
		row := i / n
		col := i % n
		if row == col {
			es[i] = 1
		}
	}
	return newMatrix(n, n, es)
}

// M creates and returns new instance of a Matrix based on the provided slice of bras (row-Vectors).
func M(bras []Vector) Matrix {
	rows := len(bras)
	if rows == 0 {
		return nil
	}
	entries := []int{}
	for _, v := range bras {
		v.ForEach(func(idx, entry int) bool {
			entries = append(entries, entry)
			return true
		})
	}
	return newMatrix(rows, bras[0].Dim(), entries)
}

func (mx *matrix) Dim() (int, int) {
	return mx.rs, mx.cs
}

func (mx *matrix) Entry(row, col int) int {
	return mx.es[mx.index(row, col)]
}

func (mx *matrix) SetEntry(row, col, entry int) {
	mx.es[mx.index(row, col)] = entry
}

func (mx *matrix) ForEach(fn func(int, int, int) bool) {
	for i, e := range mx.es {
		row := i / mx.cs
		col := i % mx.cs
		if !fn(row, col, e) {
			return
		}
	}
}

func (mx *matrix) Equals(m Matrix) bool {
	rows, cols := mx.Dim()
	rows2, cols2 := m.Dim()
	if rows != rows2 || cols != cols2 {
		return false
	}
	var ok bool
	m.ForEach(func(r, c, e int) bool {
		ok = mx.Entry(r, c) == e
		return ok
	})
	return ok
}

func (mx *matrix) MultiplyScalar(s int) Matrix {
	res := newMatrix(mx.rs, mx.cs, mx.es)
	res.ForEach(func(r, c, e int) bool {
		res.SetEntry(r, c, e*s)
		return true
	})
	return res
}

func (mx *matrix) Add(m Matrix) Matrix {
	rows, cols := mx.Dim()
	rows2, cols2 := m.Dim()
	if rows != rows2 || cols != cols2 {
		return nil
	}
	res := newMatrix(rows, cols, mx.es)
	m.ForEach(func(r, c, e int) bool {
		en := res.Entry(r, c)
		res.SetEntry(r, c, en+e)
		return true
	})
	return res
}

func (mx *matrix) Multiply(m Matrix) Matrix {
	rows, cols := mx.Dim()
	rows2, cols2 := m.Dim()
	if cols != rows2 {
		return nil
	}
	res := DimM(rows, cols2)
	res.ForEach(func(r, c, e int) bool {
		var entry int
		for i := 0; i < cols; i++ {
			entry += mx.Entry(r, i) * m.Entry(i, c)
		}
		res.SetEntry(r, c, entry)
		return true
	})
	return res
}

func (mx *matrix) Transpose() Matrix {
	res := newMatrix(mx.cs, mx.rs, make([]int, mx.rs*mx.cs))

	mx.ForEach(func(r, c, e int) bool {
		res.es[r+c*mx.rs] = e
		return true
	})

	return res
}

func (mx *matrix) IsIdentity() bool {
	return false
}
