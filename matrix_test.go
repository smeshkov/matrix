package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Matrix_DimM(t *testing.T) {
	m := DimM(2, 3)
	assert.NotNil(t, m)
}

func Test_Matrix_Dim(t *testing.T) {
	rows, cols := DimM(2, 3).Dim()
	assert.Equal(t, 2, rows)
	assert.Equal(t, 3, cols)
}

func Test_Matrix_IdentityM(t *testing.T) {
	m := IdentityM(2)
	assert.NotNil(t, m)

	rows, cols := m.Dim()
	assert.Equal(t, 2, rows)
	assert.Equal(t, 2, cols)
}

func Test_Matrix_Entry(t *testing.T) {
	m := IdentityM(2)
	assert.Equal(t, 1, m.Entry(0, 0))
	assert.Equal(t, 0, m.Entry(0, 1))
	assert.Equal(t, 0, m.Entry(1, 0))
	assert.Equal(t, 1, m.Entry(1, 1))
}

func Test_Matrix_IdentityM_2(t *testing.T) {
	m := IdentityM(3)
	assert.NotNil(t, m)

	rows, cols := m.Dim()
	assert.Equal(t, 3, rows)
	assert.Equal(t, 3, cols)

	assert.Equal(t, 1, m.Entry(0, 0))
	assert.Equal(t, 0, m.Entry(0, 1))
	assert.Equal(t, 0, m.Entry(0, 2))
	assert.Equal(t, 0, m.Entry(1, 0))
	assert.Equal(t, 1, m.Entry(1, 1))
	assert.Equal(t, 0, m.Entry(1, 2))
	assert.Equal(t, 0, m.Entry(2, 0))
	assert.Equal(t, 0, m.Entry(2, 1))
	assert.Equal(t, 1, m.Entry(2, 2))
}

func Test_Matrix_M(t *testing.T) {
	m := M([]Vector{
		V([]int{1, 0}),
		V([]int{0, 1}),
	})
	assert.NotNil(t, m)

	rows, cols := m.Dim()
	assert.Equal(t, 2, rows)
	assert.Equal(t, 2, cols)

	assert.Equal(t, 1, m.Entry(0, 0))
	assert.Equal(t, 0, m.Entry(0, 1))
	assert.Equal(t, 0, m.Entry(1, 0))
	assert.Equal(t, 1, m.Entry(1, 1))
}

func Test_Matrix_M_2(t *testing.T) {
	m := M([]Vector{})
	assert.Nil(t, m)
}

func Test_Matrix_SetEntry(t *testing.T) {
	m := IdentityM(2)
	m.SetEntry(1, 1, 2)
	assert.Equal(t, 2, m.Entry(1, 1))
}

func Test_Matrix_ForEach(t *testing.T) {
	var count, sum int
	IdentityM(2).
		ForEach(func(row, col, entry int) bool {
			count++
			sum += entry
			return true
		})
	assert.Equal(t, 4, count)
	assert.Equal(t, 2, sum)
}

func Test_Matrix_ForEach_2(t *testing.T) {
	var count, sum int
	IdentityM(2).
		ForEach(func(row, col, entry int) bool {
			count++
			sum += entry
			return false
		})
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, sum)
}

func Test_Matrix_ForEach_3(t *testing.T) {
	IdentityM(2).
		ForEach(func(row, col, entry int) bool {
			if row == 0 && col == 0 {
				assert.Equal(t, 1, entry)
			} else if row == 0 && col == 1 {
				assert.Equal(t, 0, entry)
			} else if row == 1 && col == 0 {
				assert.Equal(t, 0, entry)
			} else {
				assert.Equal(t, 1, entry)
			}
			return true
		})
}

func Test_Matrix_Equals(t *testing.T) {
	assert.True(t, DimM(2, 2).Equals(DimM(2, 2)))
	assert.True(t, DimM(3, 1).Equals(DimM(3, 1)))
	assert.False(t, DimM(2, 2).Equals(DimM(3, 2)))
	assert.False(t, DimM(2, 2).Equals(DimM(2, 3)))

	assert.True(t, IdentityM(2).Equals(IdentityM(2)))
	assert.True(t, IdentityM(3).Equals(IdentityM(3)))
	assert.True(t, IdentityM(4).Equals(IdentityM(4)))
	assert.False(t, IdentityM(4).Equals(IdentityM(3)))
	assert.False(t, IdentityM(3).Equals(IdentityM(4)))

	assert.True(t, M([]Vector{V([]int{1, 0}), V([]int{0, 1})}).Equals(IdentityM(2)))
	assert.True(t, IdentityM(2).Equals(M([]Vector{V([]int{1, 0}), V([]int{0, 1})})))
	assert.True(t, M([]Vector{V([]int{1, 2, 3}), V([]int{-15, 32, 0})}).
		Equals(M([]Vector{V([]int{1, 2, 3}), V([]int{-15, 32, 0})})))
}

func Test_Matrix_MultiplyScalar(t *testing.T) {
	assert.True(t, IdentityM(2).
		MultiplyScalar(3).
		Equals(M([]Vector{V([]int{3, 0}), V([]int{0, 3})})))

	assert.True(t, M([]Vector{V([]int{12, 7}), V([]int{0, -8})}).
		MultiplyScalar(3).
		Equals(M([]Vector{V([]int{36, 21}), V([]int{0, -24})})))

	assert.True(t, M([]Vector{V([]int{12, 7}), V([]int{0, -8})}).
		MultiplyScalar(-3).
		Equals(M([]Vector{V([]int{-36, -21}), V([]int{0, 24})})))
}

func Test_Matrix_Add(t *testing.T) {
	a := M([]Vector{
		V([]int{1, -4, 2}),
		V([]int{2, 3, 0}),
	})
	b := M([]Vector{
		V([]int{1, 2, 0}),
		V([]int{7, -3, 17}),
	})
	c := M([]Vector{
		V([]int{2, -2, 2}),
		V([]int{9, 0, 17}),
	})
	assert.True(t, a.Add(b).Equals(c))
}

func Test_Matrix_Multiply(t *testing.T) {
	a := M([]Vector{
		V([]int{1, -4, 2}),
		V([]int{2, 3, 0}),
	})
	b := M([]Vector{
		V([]int{1, 2}),
		V([]int{7, 5}),
		V([]int{6, 1}),
	})
	c := M([]Vector{
		V([]int{-15, -16}),
		V([]int{23, 19}),
	})
	assert.True(t, a.Multiply(b).Equals(c))
}

func Test_Matrix_Transpose(t *testing.T) {
	a := M([]Vector{
		V([]int{1, -4, 2}),
		V([]int{2, 3, 0}),
	})
	b := M([]Vector{
		V([]int{1, 2}),
		V([]int{-4, 3}),
		V([]int{2, 0}),
	})
	assert.True(t, a.Transpose().Equals(b))
}

func Test_Matrix_Transpose_2(t *testing.T) {
	a := M([]Vector{
		V([]int{1, 2}),
		V([]int{-4, 3}),
		V([]int{2, 0}),
	})
	b := M([]Vector{
		V([]int{1, -4, 2}),
		V([]int{2, 3, 0}),
	})
	assert.True(t, a.Transpose().Equals(b))
}

func Test_Matrix_IsIdentity(t *testing.T) {
	assert.True(t, IdentityM(2).IsIdentity())
	assert.True(t, IdentityM(3).IsIdentity())
	assert.True(t, IdentityM(4).IsIdentity())
	assert.True(t, IdentityM(5).IsIdentity())
	assert.True(t, M([]Vector{
		V([]int{1, 0}),
		V([]int{0, 1}),
	}).IsIdentity())
	assert.True(t, M([]Vector{
		V([]int{1, 0, 0}),
		V([]int{0, 1, 0}),
		V([]int{0, 0, 1}),
	}).IsIdentity())

	assert.False(t, M([]Vector{
		V([]int{0, 1}),
		V([]int{1, 0}),
	}).IsIdentity())
	assert.False(t, M([]Vector{
		V([]int{0, 1, 0}),
		V([]int{0, 1, 0}),
		V([]int{1, 0, 1}),
	}).IsIdentity())
}
