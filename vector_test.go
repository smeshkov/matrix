package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Vector_DimV(t *testing.T) {
	assert.NotNil(t, DimV(2))
}

func Test_Vector_V(t *testing.T) {
	assert.NotNil(t, V([]int{1, 0}))
}

func Test_Vector_Dim(t *testing.T) {
	v := DimV(2)
	assert.Equal(t, 2, v.Dim())
	assert.Equal(t, v.Dim(), V([]int{1, 0}).Dim())
}

func Test_Vector_ForeEach(t *testing.T) {
	var count int
	var sum int
	V([]int{1, 2}).
		ForEach(func(i int, s int) bool {
			count++
			sum += s
			return true
		})
	assert.Equal(t, 2, count)
	assert.Equal(t, 3, sum)
}

func Test_Vector_ForeEach_2(t *testing.T) {
	var count int
	var sum int
	V([]int{1, 2}).
		ForEach(func(i int, s int) bool {
			count++
			sum += s
			return false
		})
	assert.Equal(t, 1, count)
	assert.Equal(t, 1, sum)
}

func Test_Vector_Equals(t *testing.T) {
	assert.True(t, V([]int{1, 0}).Equals(V([]int{1, 0})))
	assert.False(t, V([]int{1, 0}).Equals(V([]int{0, 1})))
	assert.False(t, V([]int{1, 0}).Equals(V([]int{1, 0, 0})))
}

func Test_Vector_MultiplyScalar(t *testing.T) {
	V([]int{1, 0, 3}).
		MultiplyScalar(2).
		ForEach(func(i int, s int) bool {
			switch i {
			case 0:
				assert.Equal(t, 2, s)
			case 1:
				assert.Equal(t, 0, s)
			case 2:
				assert.Equal(t, 6, s)
			}
			return true
		})
}

func Test_Vector_Add(t *testing.T) {
	ok := V([]int{1, 0, 3}).
		Add(V([]int{2, 5, -4})).
		Equals(V([]int{3, 5, -1}))
	assert.True(t, ok)
}

func Test_Vector_Multiply(t *testing.T) {
	assert.Equal(t, 1, V([]int{1, 0}).Multiply(V([]int{1, 0})))
	assert.Equal(t, 1, V([]int{0, 1}).Multiply(V([]int{0, 1})))
	assert.Equal(t, 0, V([]int{1, 0}).Multiply(V([]int{0, 1})))
}

func Test_Vector_Len(t *testing.T) {
	assert.Equal(t, 1.0, V([]int{1, 0}).Len())
	assert.Equal(t, 1.0, V([]int{0, 1}).Len())
	assert.Equal(t, 3.0, V([]int{3, 0}).Len())
}

func Test_Vector_IsUnit(t *testing.T) {
	assert.True(t, V([]int{1, 0}).IsUnit())
	assert.True(t, V([]int{0, 1}).IsUnit())
	assert.False(t, V([]int{0, 2}).IsUnit())
	assert.False(t, V([]int{2, 0}).IsUnit())
	assert.False(t, V([]int{1, 1}).IsUnit())
}

func Test_Vector_IsOrthogonal(t *testing.T) {
	assert.True(t, V([]int{1, 0}).IsOrthogonal(V([]int{0, 1})))
	assert.True(t, V([]int{0, 1}).IsOrthogonal(V([]int{1, 0})))
	assert.False(t, V([]int{1, 0}).IsOrthogonal(V([]int{1, 0})))
	assert.False(t, V([]int{0, 1}).IsOrthogonal(V([]int{0, 1})))
}
