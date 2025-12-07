package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Matrix представляет матрицу
type Matrix struct {
	Rows int
	Cols int
	Data [][]float64
}

// NewMatrix создает новую матрицу
func NewMatrix(rows, cols int) *Matrix {
	data := make([][]float64, rows)
	for i := range data {
		data[i] = make([]float64, cols)
	}
	return &Matrix{Rows: rows, Cols: cols, Data: data}
}

// Print выводит матрицу на экран
func (m *Matrix) Print() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			fmt.Printf("%8.2f", m.Data[i][j])
		}
		fmt.Println()
	}
}

// InputMatrix читает матрицу с консоли
func InputMatrix(rows, cols int, name string) *Matrix {
	scanner := bufio.NewScanner(os.Stdin)
	matrix := NewMatrix(rows, cols)

	fmt.Printf("Ввод матрицы %s (%dx%d):\n", name, rows, cols)
	for i := 0; i < rows; i++ {
		for {
			fmt.Printf("Строка %d (введите %d чисел через пробел): ", i+1, cols)
			if !scanner.Scan() {
				return nil
			}

			line := strings.TrimSpace(scanner.Text())
			parts := strings.Fields(line)

			if len(parts) != cols {
				fmt.Printf("Ошибка: нужно ввести ровно %d чисел\n", cols)
				continue
			}

			valid := true
			for j, part := range parts {
				val, err := strconv.ParseFloat(part, 64)
				if err != nil {
					fmt.Printf("Ошибка: '%s' не является числом\n", part)
					valid = false
					break
				}
				matrix.Data[i][j] = val
			}

			if valid {
				break
			}
		}
	}
	return matrix
}

// Add складывает две матрицы
func Add(a, b *Matrix) (*Matrix, error) {
	if a.Rows != b.Rows || a.Cols != b.Cols {
		return nil, fmt.Errorf("матрицы должны быть одного размера")
	}

	result := NewMatrix(a.Rows, a.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			result.Data[i][j] = a.Data[i][j] + b.Data[i][j]
		}
	}
	return result, nil
}

// MultiplyByScalar умножает матрицу на число
func MultiplyByScalar(a *Matrix, scalar float64) *Matrix {
	result := NewMatrix(a.Rows, a.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < a.Cols; j++ {
			result.Data[i][j] = a.Data[i][j] * scalar
		}
	}
	return result
}

// Multiply умножает две матрицы
func Multiply(a, b *Matrix) (*Matrix, error) {
	if a.Cols != b.Rows {
		return nil, fmt.Errorf("число столбцов первой матрицы должно равняться числу строк второй")
	}

	result := NewMatrix(a.Rows, b.Cols)
	for i := 0; i < a.Rows; i++ {
		for j := 0; j < b.Cols; j++ {
			for k := 0; k < a.Cols; k++ {
				result.Data[i][j] += a.Data[i][k] * b.Data[k][j]
			}
		}
	}
	return result, nil
}

// DisplayMenu показывает меню
func DisplayMenu() {
	fmt.Println("\n=== КАЛЬКУЛЯТОР МАТРИЦ ===")
	fmt.Println("1. Сложение матриц")
	fmt.Println("2. Умножение матрицы на число")
	fmt.Println("3. Умножение матриц")
	fmt.Println("4. Выход")
	fmt.Print("Выберите операцию: ")
}

// GetMatrixSize получает размер матрицы от пользователя
func GetMatrixSize() (int, int) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Выберите размер матрицы (1 - 2x2, 2 - 3x3): ")
		if !scanner.Scan() {
			return 0, 0
		}

		choice := strings.TrimSpace(scanner.Text())
		switch choice {
		case "1":
			return 2, 2
		case "2":
			return 3, 3
		default:
			fmt.Println("Ошибка: выберите 1 или 2")
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		DisplayMenu()

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			// Сложение матриц
			rows, cols := GetMatrixSize()
			if rows == 0 {
				continue
			}

			fmt.Println("\n--- Сложение матриц ---")
			matrixA := InputMatrix(rows, cols, "A")
			if matrixA == nil {
				continue
			}

			matrixB := InputMatrix(rows, cols, "B")
			if matrixB == nil {
				continue
			}

			result, err := Add(matrixA, matrixB)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}

			fmt.Println("\nМатрица A:")
			matrixA.Print()
			fmt.Println("\nМатрица B:")
			matrixB.Print()
			fmt.Println("\nРезультат сложения (A + B):")
			result.Print()

		case "2":
			// Умножение на число
			rows, cols := GetMatrixSize()
			if rows == 0 {
				continue
			}

			fmt.Println("\n--- Умножение матрицы на число ---")
			matrixA := InputMatrix(rows, cols, "A")
			if matrixA == nil {
				continue
			}

			for {
				fmt.Print("Введите число для умножения: ")
				if !scanner.Scan() {
					continue
				}

				scalar, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
				if err != nil {
					fmt.Println("Ошибка: введите корректное число")
					continue
				}

				result := MultiplyByScalar(matrixA, scalar)

				fmt.Println("\nМатрица A:")
				matrixA.Print()
				fmt.Printf("\nЧисло: %.2f\n", scalar)
				fmt.Printf("\nРезультат умножения (A × %.2f):\n", scalar)
				result.Print()
				break
			}

		case "3":
			// Умножение матриц
			fmt.Println("\n--- Умножение матриц ---")

			// Ввод первой матрицы
			rows1, cols1 := GetMatrixSize()
			if rows1 == 0 {
				continue
			}

			matrixA := InputMatrix(rows1, cols1, "A")
			if matrixA == nil {
				continue
			}

			// Ввод второй матрицы
			fmt.Printf("Матрица B должна иметь размер %dx?\n", cols1)
			rows2, cols2 := GetMatrixSize()
			if rows2 == 0 {
				continue
			}

			if rows2 != cols1 {
				fmt.Printf("Ошибка: матрица B должна иметь %d строк\n", cols1)
				continue
			}

			matrixB := InputMatrix(rows2, cols2, "B")
			if matrixB == nil {
				continue
			}

			result, err := Multiply(matrixA, matrixB)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
				continue
			}

			fmt.Println("\nМатрица A:")
			matrixA.Print()
			fmt.Println("\nМатрица B:")
			matrixB.Print()
			fmt.Println("\nРезультат умножения (A × B):")
			result.Print()

		case "4":
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Ошибка: выберите операцию от 1 до 4")
		}

		fmt.Print("\nНажмите Enter для продолжения...")
		scanner.Scan()
	}
}
