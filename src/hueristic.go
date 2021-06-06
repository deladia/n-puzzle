package main

import "math"

func hamming_distance(grid [][]byte, ideal [][]byte) int {
	size := len(grid)
	h := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			i_tmp, j_tmp := find_index(ideal, grid[i][j])
			if i_tmp != i || j_tmp != j {
				h++
			}
		}
	}
	return h
}

func manhattan_distance(grid [][]byte, ideal [][]byte) int {
	size := len(grid)
	h := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				continue
			}
			i_tmp, j_tmp := find_index(ideal, grid[i][j])
			h += int(math.Abs(float64(i-i_tmp)) + math.Abs(float64(j-j_tmp)))
		}
	}
	return h
}

func linear_conflict_manhattan_distance(grid [][]byte, ideal [][]byte) int {
	size := len(grid)
	h := 0
	line_conf := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size - 1; j++ {
			if grid[i][j] == 0 {
				continue 
			}
			for k := j + 1; k < size; k++ {
				if grid[i][k] == 0 {
					continue 
				}
				i_tmp_1, j_tmp_1 := find_index(ideal, grid[i][j])
				i_tmp_2, j_tmp_2 := find_index(ideal, grid[i][k])
				if i_tmp_1 != i || i_tmp_2 != i {
					continue
				}
				if j_tmp_1 > j_tmp_2 && j < k {
					line_conf++
				} else if j_tmp_1 < j_tmp_2 && j > k {
					line_conf++
				}
			}
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size - 1; j++ {
			if grid[j][i] == 0 {
				continue 
			}
			for k := j + 1; k < size; k++ {
				if grid[k][i] == 0 {
					continue 
				}
				i_tmp_1, j_tmp_1 := find_index(ideal, grid[j][i])
				i_tmp_2, j_tmp_2 := find_index(ideal, grid[k][i])
				if j_tmp_1 != i || j_tmp_2 != i {
					continue
				}
				if i_tmp_1 > i_tmp_2 && j < k {
					line_conf++
				} else if i_tmp_1 < i_tmp_2 && j > k {
					line_conf++
				}
			}
		}
	}
	h += line_conf * 2 + manhattan_distance(grid, ideal)
	return h
}
