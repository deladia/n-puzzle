package main

import "fmt"

var (
	Red     = "\033[1;31m%s\033[0m"
	Green   = "\033[1;32m%s\033[0m"
	Yellow  = "\033[1;33m%s\033[0m"
	Purple  = "\033[1;34m%s\033[0m"
	White   = "\033[1;37m%s\033[0m"
  )

func (answer *Node) print_list() {
	for answer != nil {
		fmt.Println("g =", answer.g, "h =", answer.h, "f =", answer.f)
		for _, v := range answer.grid {
			fmt.Println(v)
		}
		fmt.Println()
		answer = answer.parent
	}
}

func find_index(grid [][]byte, sym byte) (int, int) {
	for i, v := range grid {
		for j, v2 := range v {
			if v2 == sym {
				return i, j
			}
		}
	}
	return -1, -1
}

func initial(grid [][]byte, ideal [][]byte, hueristic *string, g int) *Node {
	var new_node Node

	new_node.grid = grid
	new_node.g = g
	switch *hueristic {
	case "hamming":
		new_node.h = hamming_distance(grid, ideal)
	case "manhattan":
		new_node.h = manhattan_distance(grid, ideal)
	case "conflict":
		new_node.h = linear_conflict_manhattan_distance(grid, ideal)
	default:
		new_node.h = 0
	}
	new_node.f = new_node.h + g
	new_node.parent = nil
	new_node.next = nil
	return &new_node
}

func to_string(grid [][]byte) string {
	answer := ""
	for _, v := range grid {
		answer += string(v)
	}
	return answer
}

func (n *Node) len_list() int {
	tmp := n
	cnt := 0
	for tmp != nil {
		cnt++
		tmp = tmp.next
	}
	return cnt
}

func (n *Node) len_answer() int {
	tmp := n
	cnt := 0
	for tmp != nil {
		cnt++
		tmp = tmp.parent
	}
	return cnt - 1
}

func (n *Node) insert(new *Node) *Node {
	if n != nil {
		var tmp_list *Node
		var prev *Node

		tmp_list = n
		prev = nil
		for tmp_list != nil {
			if tmp_list.next != nil && tmp_list.next.f >= new.f {
				new.next = tmp_list.next
				tmp_list.next = new
				return n
			} else if tmp_list.next == nil && tmp_list.f >= new.f {
				new.next = tmp_list
				if prev != nil {
					prev.next = new
					return n
				} else {
					return new
				}
			} else if tmp_list.next == nil && tmp_list.f < new.f {
				new.next = tmp_list.next
				tmp_list.next = new
				return n
			}
			prev = tmp_list
			tmp_list = tmp_list.next
		}
	}
	return new
}

func (list *Node) find_repeat(to_find *Node) *Node {
	var tmp_list *Node

	tmp_list = list
	for tmp_list != nil {
		if equal_grid(to_find.grid, tmp_list.grid) {
			return tmp_list
		}
		tmp_list = tmp_list.next
	}
	return nil
}

func equal_grid(node1 [][]byte, node2 [][]byte) bool {
	for i, v := range node1 {
		for j, vv := range v {
			if vv != node2[i][j] {
				return false
			}
		}
	}
	return true
}
