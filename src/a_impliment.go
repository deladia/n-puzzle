package main

import "fmt"

func create_4_child(grid_src [][]byte, index int) [][]byte {
	size := len(grid_src)
	i, j := find_index(grid_src, 0)

	grid := make([][]byte, size)
	for k := range grid_src {
		grid[k] = make([]byte, size)
		copy(grid[k], grid_src[k])
	}
	switch index {
	case 0:
		if j + 1 < size {
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
		if i + 1 < size {
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
	max := 1
	for open_list != nil {
		process := open_list
		if process.h == 0 {
			fmt.Printf(Purple + " %v " + Purple, "compexity in time =", len(closed_map), "evaluated node\n")
			fmt.Printf(Purple + " %v " + Purple, "compexity in size =", max, "max node in memory\n")
			fmt.Printf(Purple + " %v " + Purple, "total steps", process.len_answer(), "\n\n")
			return process
		}
		open_list = open_list.next
		closed_map[to_string(process.grid)] = true
		for i := 0; i < 4; i++ {
			if new_grid := create_4_child(process.grid, i); new_grid != nil {
				if _, ok := closed_map[to_string(new_grid)]; ok == true {
					continue
				}
				max++
				node := initial(new_grid, ideal_grid, hueristic, process.g + 1)
				node.parent = process
				actual_node := open_list.find_repeat(node)
				if actual_node == nil {
					open_list = open_list.insert(node)
				} else {
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

func greedy_search(karta [][]byte, hueristic *string) *Node {
	ideal_grid := create_goal(len(karta))
	open_list := initial(karta, ideal_grid, hueristic, 0)
	closed_map := make(map[string]bool)
	max := 1
	for open_list != nil {
		process := open_list
		if process.h == 0 {
			fmt.Printf(Purple + " %v " + Purple, "compexity in time =", len(closed_map), "evaluated node\n")
			fmt.Printf(Purple + " %v " + Purple, "compexity in size =", max, "max node in memory\n")
			fmt.Printf(Purple + " %v " + Purple, "total steps", process.len_answer(), "\n\n")
			return process
		}
		open_list = open_list.next
		closed_map[to_string(process.grid)] = true
		for i := 0; i < 4; i++ {
			if new_grid := create_4_child(process.grid, i); new_grid != nil {
				if _, ok := closed_map[to_string(new_grid)]; ok == false {
					node := initial(new_grid, ideal_grid, hueristic, 0)
					node.parent = process
					open_list = open_list.insert(node)
					max++
				}
			}
		}
	}
	return nil
}

func uniformcost_search(karta [][]byte) *Node {
	ideal_grid := create_goal(len(karta))
	var open_list *Node = &Node{
		grid:	karta,
		g:		0,
		h:		0,
		f:		0,
		parent: nil,
		next:	nil}
	closed_map := make(map[string]bool)
	max := 1
	for open_list != nil {
		process := open_list
		if equal_grid(process.grid, ideal_grid) {
			fmt.Printf(Purple + " %v " + Purple, "compexity in time =", len(closed_map), "evaluated node\n")
			fmt.Printf(Purple + " %v " + Purple, "compexity in size =", max, "max node in memory\n")
			fmt.Printf(Purple + " %v " + Purple, "total steps", process.len_answer(), "\n\n")
			return process
		}
		open_list = open_list.next
		closed_map[to_string(process.grid)] = true
		for i := 0; i < 4; i++ {
			if new_grid := create_4_child(process.grid, i); new_grid != nil {
				if _, ok := closed_map[to_string(new_grid)]; ok == true {
					continue
				}
				max++
				var node *Node = &Node{
					grid:	new_grid,
					g:		process.g + 1,
					h:		0,
					f:		0,
					parent: process,
					next:	nil}
				node.parent = process
				actual_node := open_list.find_repeat(node)
				if actual_node == nil {
					open_list = open_list.insert(node)
				} else {
					if node.g < actual_node.g {
						actual_node.g = node.g
						actual_node.parent = node.parent
					}
				}
			}
		}
	}
	return nil
}