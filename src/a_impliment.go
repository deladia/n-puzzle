package main

func create_4_child(grid_src [][]byte, index int) [][]byte {
	
	i, j := find_index(grid_src, 0)

	grid := make([][]byte, len(grid_src))
	for k := range grid_src {
		grid[k] = make([]byte, len(grid_src[i]))
		copy(grid[k], grid_src[k])
	}
	switch index {
	case 0:
		if j + 1 < len(grid) {
			grid[i][j], grid[i][j + 1] = grid[i][j + 1], grid[i][j]
		} else {
			return nil
		}
	case 1:
		if j - 1 >= 0 {
			grid[i][j], grid[i][j - 1] = grid[i][j - 1], grid[i][j]
		} else {
			return nil
		}
	case 2:
		if i + 1 < len(grid) {
			grid[i][j], grid[i + 1][j] = grid[i + 1][j], grid[i][j]
		} else {
			return nil
		}
	case 3:
		if i - 1 >= 0 {
			grid[i][j], grid[i - 1][j] = grid[i - 1][j], grid[i][j]
		} else {
			return nil
		}
	}
	return grid
}

func a_implement(karta [][]byte, hueristic *string) *Node {
	ideal_grid := create_goal(len(karta))
	open_list := initial(karta, ideal_grid, hueristic, 0)
	closed_map := make(map[string]bool)
	for open_list != nil {
		process := open_list
		if process.h == 0 {
			return process
		}
		open_list = open_list.next
		closed_map[to_string(process.grid)] = true
		for i := 0; i < 4; i++ {
			if new_grid := create_4_child(process.grid, i); new_grid != nil {
				node := initial(new_grid, ideal_grid, hueristic, process.g + 1)
				node.parent = process
				if _, ok := closed_map[to_string(node.grid)]; ok == true {
					continue
				}
				if open_list.find_repeat(node) == nil {
					open_list = open_list.insert(node)
				} else {
					actual_node := open_list.find_repeat(node)
					if node.g < actual_node.g {
						actual_node.g = node.g
						actual_node.f = node.f
						actual_node.parent = node.parent
					}
				}
			}
		}
	}
	return nil
}