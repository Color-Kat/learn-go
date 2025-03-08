package matrices_test

import (
	"algebraMachine/matrices"
	"testing"
)

func CompareMatrices[T int | bool](m1, m2 matrices.Matrix[T]) bool {
	if m1.Rows != m2.Rows || m1.Cols != m2.Cols {
		return false
	}

	for i := 0; i < m1.Rows; i++ {
		for j := 0; j < m1.Cols; j++ {
			if m1.Data[i][j] != m2.Data[i][j] {
				return false
			}
		}
	}

	return true
}

// ===== Add Int ===== //

var intAddTestCases = []struct {
	m1     [][]int
	m2     [][]int
	expect [][]int
}{
	{
		m1:     [][]int{{1, 1}, {1, 1}},
		m2:     [][]int{{1, 1}, {1, 1}},
		expect: [][]int{{2, 2}, {2, 2}},
	},
	{
		m1:     [][]int{{-5, -5}, {-10, -10}},
		m2:     [][]int{{5, 5}, {10, 10}},
		expect: [][]int{{0, 0}, {0, 0}},
	},
}

func TestIntMatrixOperation_Add(t *testing.T) {
	for ts := range intAddTestCases {
		t.Run("Add Int Matrices", func(t *testing.T) {
			m1 := matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intAddTestCases[ts].m1,
			}

			m2 := matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intAddTestCases[ts].m2,
			}

			expected := &matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intAddTestCases[ts].expect,
			}

			got := (&matrices.IntMatrixOperation{
				MatrixA: m1,
				MatrixB: m2,
			}).Add().(*matrices.Matrix[int])

			if !CompareMatrices(*got, *expected) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})
	}
}

// ===== Multiply Int ===== //

var intMultiplyTestCases = []struct {
	m1     [][]int
	m2     [][]int
	expect [][]int
}{
	{
		m1:     [][]int{{1, 0}, {0, 1}},
		m2:     [][]int{{1, 12}, {173, 0}},
		expect: [][]int{{1, 12}, {173, 0}},
	},
	{
		m1:     [][]int{{1, 2}, {3, 4}},
		m2:     [][]int{{1, 2}, {3, 4}},
		expect: [][]int{{7, 10}, {15, 22}},
	},
}

func TestIntMatrixOperation_Multiply(t *testing.T) {
	for ts := range intMultiplyTestCases {
		t.Run("Multiply Int Matrices", func(t *testing.T) {
			m1 := matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intMultiplyTestCases[ts].m1,
			}

			m2 := matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intMultiplyTestCases[ts].m2,
			}

			expected := &matrices.Matrix[int]{
				Rows: 2,
				Cols: 2,
				Data: intMultiplyTestCases[ts].expect,
			}

			got := (&matrices.IntMatrixOperation{
				MatrixA: m1,
				MatrixB: m2,
			}).Multiply().(*matrices.Matrix[int])

			if !CompareMatrices(*got, *expected) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})
	}
}

// ===== Add Bool ===== //

var boolAddTestCases = []struct {
	m1     [][]bool
	m2     [][]bool
	expect [][]bool
}{
	{
		m1:     [][]bool{{true, true}, {true, true}},
		m2:     [][]bool{{false, false}, {false, false}},
		expect: [][]bool{{true, true}, {true, true}},
	},
	{
		m1:     [][]bool{{true, false}, {true, false}},
		m2:     [][]bool{{true, false}, {true, false}},
		expect: [][]bool{{true, false}, {true, false}},
	},
}

func TestBoolMatrixOperation_Add(t *testing.T) {
	for ts := range boolAddTestCases {
		t.Run("Add Bool Matrices", func(t *testing.T) {
			m1 := matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolAddTestCases[ts].m1,
			}

			m2 := matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolAddTestCases[ts].m2,
			}

			expected := &matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolAddTestCases[ts].expect,
			}

			got := (&matrices.BoolMatrixOperation{
				MatrixA: m1,
				MatrixB: m2,
			}).Add().(*matrices.Matrix[bool])

			if !CompareMatrices(*got, *expected) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})
	}
}

// ===== Multiply Bool ===== //

var boolMultiplyTestCases = []struct {
	m1     [][]bool
	m2     [][]bool
	expect [][]bool
}{
	{
		m1:     [][]bool{{true, false}, {false, true}},
		m2:     [][]bool{{true, true}, {false, false}},
		expect: [][]bool{{true, true}, {false, false}},
	},
	{
		m1:     [][]bool{{true, false}, {true, true}},
		m2:     [][]bool{{true, false}, {true, true}},
		expect: [][]bool{{true, false}, {true, true}},
	},
}

func TestBoolMatrixOperation_Multiply(t *testing.T) {
	for ts := range boolAddTestCases {
		t.Run("Multiply Bool Matrices", func(t *testing.T) {
			m1 := matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolMultiplyTestCases[ts].m1,
			}

			m2 := matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolMultiplyTestCases[ts].m2,
			}

			expected := &matrices.Matrix[bool]{
				Rows: 2,
				Cols: 2,
				Data: boolMultiplyTestCases[ts].expect,
			}

			got := (&matrices.BoolMatrixOperation{
				MatrixA: m1,
				MatrixB: m2,
			}).Multiply().(*matrices.Matrix[bool])

			if !CompareMatrices(*got, *expected) {
				t.Errorf("Expected %v, got %v", expected, got)
			}
		})
	}
}
