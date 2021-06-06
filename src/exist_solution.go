package main

import (
	"math"
)

func create_goal(size int) [][]byte {
	ideal_gread := make([][]byte, size)
	for i := range ideal_gread {
		ideal_gread[i] = make([]byte, size)
	}
	var number byte = 1
	left := 0
	right := size - 1
	up := 0
	down := size - 1
	for i, j := 0, 0; int(number) < size * size; {
		for j < right {
			ideal_gread[i][j] = number
			j++
			number++
		}
		up++
		for i < down {
			ideal_gread[i][j] = number
			i++
			number++
		}
		right--
		for j > left {
			ideal_gread[i][j] = number
			j--
			number++
		}
		down--
		for i > up {
			ideal_gread[i][j] = number
			i--
			number++
		}
		left++
	}
	return ideal_gread
}

func exist_solution(grid [][]byte, ideal [][]byte) bool {
	size := len(grid)
	inv := 0
	zero_dist := 0
	for i := 0; i < size * size; i++ {
		if grid[i / size][i % size] == 0 {
			i_tmp, j_tmp := find_index(ideal, 0)
			zero_dist = int(math.Abs(float64(i / size - i_tmp)) + math.Abs(float64(i % size - j_tmp)))
		}
		for j := 0; j < i; j++ {
			i1, j1 := find_index(ideal, grid[j / size][j % size])
			i2, j2 := find_index(ideal, grid[i / size][i % size])
			if i1 * size + j1 > i2 * size + j2 {
				inv++
			}
		}
	}
	if inv % 2 == 0 && zero_dist % 2 == 0 {
		return true
	} else if inv % 2 == 1 && zero_dist % 2 == 1 {
		return true
	}
	return false
}