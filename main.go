package main

import "fmt"

func main() {
	getDiogonalMatix(10, 10)
	getReverseDiogonalMatix(10, 10)
}

func getDiogonalMatix(axisX int, axisY int) {
	matrix := make([][]int, axisX)

	for x := 0; x < axisX; x++ {
		matrix[x] = make([]int, axisY)

		for y := 0; y < axisY; y++ {
			if x == y {
				matrix[x][y] = y
			}
		}

		fmt.Println(matrix[x])
	}
}

func getReverseDiogonalMatix(axisX int, axisY int) {
	matrix := make([][]int, axisX)

	for x := 0; x < axisX; x++ {
		matrix[x] = make([]int, axisY)

		for y := axisY - 1; y > 0; y-- {
			if x + y == axisY - 1{
				matrix[x][y] = y
			}
		}

		fmt.Println(matrix[x])
	}

}