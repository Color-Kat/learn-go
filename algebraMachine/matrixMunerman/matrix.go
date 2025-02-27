package matrixMunerman

// MatrixOperations предоставляет методы для операций над матрицами.
type MatrixOperations struct {
	row, col, k int // Вспомогательные переменные для итераций
}

// Add выполняет поэлементное сложение матриц, используя переданную функцию.
func (mo *MatrixOperations) Add(size int, addFunc func()) {
	for mo.row = 0; mo.row < size; mo.row++ {
		for mo.col = 0; mo.col < size; mo.col++ {
			addFunc() // Вызов функции для сложения элементов
		}
	}
}

// Multiply выполняет умножение матриц, используя переданные функции для инициализации и умножения.
func (mo *MatrixOperations) Multiply(size int, initFunc func(), multiplyFunc func()) {
	for mo.row = 0; mo.row < size; mo.row++ {
		for mo.k = 0; mo.k < size; mo.k++ {
			initFunc() // Инициализация элемента результата
			for mo.col = 0; mo.col < size; mo.col++ {
				multiplyFunc() // Вычисление элемента результата
			}
		}
	}
}

// AddInt выполняет поэлементное сложение для целочисленных матриц.
func (mo *MatrixOperations) AddInt(matrixA, matrixB, resultMatrix [][]int, size int) {
	mo.Add(size, func() {
		resultMatrix[mo.row][mo.col] = matrixA[mo.row][mo.col] + matrixB[mo.row][mo.col]
	})
}

// MultiplyInt выполняет умножение для целочисленных матриц.
func (mo *MatrixOperations) MultiplyInt(matrixA, matrixB, resultMatrix [][]int, size int) {
	mo.Multiply(size, func() {
		resultMatrix[mo.row][mo.k] = 0 // Инициализация элемента результата
	}, func() {
		resultMatrix[mo.row][mo.k] += matrixA[mo.row][mo.col] * matrixB[mo.col][mo.k]
	})
}

// AddBool выполняет поэлементное сложение для булевых матриц (логическое ИЛИ).
func (mo *MatrixOperations) AddBool(matrixA, matrixB, resultMatrix [][]bool, size int) {
	mo.Add(size, func() {
		resultMatrix[mo.row][mo.col] = matrixA[mo.row][mo.col] || matrixB[mo.row][mo.col]
	})
}

// MultiplyBool выполняет умножение для булевых матриц (логическое И и ИЛИ).
func (mo *MatrixOperations) MultiplyBool(matrixA, matrixB, resultMatrix [][]bool, size int) {
	mo.Multiply(size, func() {
		resultMatrix[mo.row][mo.k] = false // Инициализация элемента результата
	}, func() {
		resultMatrix[mo.row][mo.k] = resultMatrix[mo.row][mo.k] || (matrixA[mo.row][mo.col] && matrixB[mo.col][mo.k])
	})
}
