package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverseSlice(slice []int) {
	n := len(slice) - 1
	for i, j := 0, n; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func printMatrix(matrix [][]int) {
	var str string
	for _, list := range matrix {
		for _, v := range list {
			str += strconv.Itoa(v) + ", "
		}
		str = str[:len(str)-2] + "\n"
	}
	fmt.Println(str)
}

func generateMagicSquare() [][]int {
	x := 3
	m := make([][]int, x)
	for i := range m {
		m[i] = make([]int, x)
	}
	a, b := 0, x/2
	m[a][b] = 1
	for i := 2; i <= x*x; i++ {
		var c, d int
		if a > 0 {
			c = a - 1
		} else {
			c = x - 1
		}
		if b < x-1 {
			d = b + 1
		} else {
			d = 0
		}

		if m[c][d] != 0 {
			c, d = a+1, b
		}
		m[c][d] = i
		a, b = c, d
	}
	return m
}

func rotateMatrix(matrix [][]int) {
	n := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < n-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = matrix[j][n-i]
			matrix[j][n-i] = temp
		}
	}
}

func flipMatrix(matrix [][]int) [][]int {
	x := [][]int{}
	for _, list := range matrix {
		listCopy := make([]int, len(list))
		copy(listCopy, list)
		reverseSlice(listCopy)
		x = append(x, listCopy)
	}
	return x
}

func convertMatrixToArray(matrix [][]int) []int {
	x := []int{}
	for _, list := range matrix {
		for _, v := range list {
			x = append(x, v)
		}
	}
	return x
}

func magicSquareInit() [][]int {
	x := make([][]int, 0)

	msRegular := generateMagicSquare()
	msFlipped := flipMatrix(msRegular)
	x = append(x, convertMatrixToArray(msRegular), convertMatrixToArray(msFlipped))

	for i := 0; i < 3; i++ {
		rotateMatrix(msRegular)
		msFlipped = flipMatrix(msRegular)
		x = append(x, convertMatrixToArray(msRegular), convertMatrixToArray(msFlipped))
	}
	return x
}

func magicSquareForming(matrix [][]int) int {
	g := magicSquareInit()
	// printMatrix(g)

	f := convertMatrixToArray(matrix)
	r := 99

	for _, list := range g {
		s := 0
		for i, v := range f {
			a := math.Abs(float64(v) - float64(list[i]))
			s += int(a)
		}
		if s < r {
			r = s
		}
	}
	return r
}

func main() {
	// matrix := generateMagicSquare()
	// printMatrix(matrix)

	// rotateMatrix(matrix)
	// printMatrix(matrix)

	// flipped := flipMatrix(matrix)
	// printMatrix(flipped)

	// fmt.Println(convertMatrixToArray(matrix))
	// fmt.Println(convertMatrixToArray(flipped))
	x := [][]int{
		{5, 3, 4},
		{1, 5, 8},
		{6, 4, 2},
	}
	fmt.Println(magicSquareForming(x))
}
