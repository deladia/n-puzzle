package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	grid	[][]byte
	g		int
	h		int
	f		int
	parent	*Node
	next	*Node
}

func create_goal(size int) [][]byte {
	ideal_gread := make([][]byte, size)
	for i := range ideal_gread {
		ideal_gread[i] = make([]byte, size)
	}
	var number byte = 0
	left := 0
	right := size - 1
	up := 0
	down := size * size - 1
	for i := 0; i < size * size - 1; i++ {
		for i <= right {
			ideal_gread[i / size][i % size] = number
			i++
			number++
		}
		up++
		for i <= down {
			ideal_gread[i / size][i % size] = number
			i += size
			number++
		}
		right--
		for i >= left {
			ideal_gread[i / size][i % size] = number
			i--
			number++
		}
		down -= size
		for i <= up {
			ideal_gread[i / size][i % size] = number
			i++
			number++
		}
		left++
	}
	return ideal_gread
}

func hamming_distance(grid [][]byte) int {
	size := len(grid)
	h := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if size * i + j + 1 != int(grid[i][j]) && size * size > size * i + j + 1 {
				h++
			} 
		}
	}
	return h
}

func manhattan_distance(grid [][]byte) int {
	size := len(grid)
	h := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if size * i + j + 1 != int(grid[i][j]) {
				i_tmp := int((grid[i][j] - 1)) / size
				j_tmp := int((grid[i][j] - 1)) % size
				h += int(math.Abs(float64(i - i_tmp)) + math.Abs(float64(j - j_tmp)))
			}
		}
	}
	return h
}

//Херня
func linear_conflict_manhattan_distance(grid [][]byte) int {
	size := len(grid)
	h := 0
	line_conf := 0
	col_coef := make([]int, size) 
	for i := 0; i < size; i++ {
		row_coef := 0
		for j := 0; j < size; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if size * i + j + 1 != int(grid[i][j]) {
				i_tmp := int((grid[i][j] - 1)) / size
				j_tmp := int((grid[i][j] - 1)) % size
				h += int(math.Abs(float64(i - i_tmp)) + math.Abs(float64(j - j_tmp)))
				//По столбцам
				if j_tmp == 0 {
					col_coef[j] += 1
				}
				//По строкам
				if i_tmp == 0 {
					row_coef++
				}
			} 
		}
		if row_coef != 0 {
			line_conf += (row_coef - 1)
		}
	}
	for _, v := range col_coef {
		if v != 0 {
			line_conf += (v - 1)
		}
	}
	h += line_conf * 2
	return h
}

func initial(grid [][]byte, g int) *Node {
	var new_node Node

	new_node.grid = grid
	new_node.g = g
	new_node.h = manhattan_distance(grid)
	new_node.f = new_node.h + g
	new_node.parent = nil
	new_node.next = nil
	return &new_node
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

func equal_grid(node1 *Node, node2 *Node) bool {
	for i, v := range node1.grid {
		for j, vv := range v {
			if vv != node2.grid[i][j] {
				return false
			}
		}
	}
	return true
}

func (list *Node) find_repeat(to_find *Node) *Node {
	var tmp_list *Node

	tmp_list = list
	for tmp_list != nil {
		if equal_grid(to_find, tmp_list) {
			return tmp_list
		}
		tmp_list = tmp_list.next
	}
	return nil
}

// func (list *Node) remove(to_remove *Node) *Node {
// 	var tmp_list *Node

// 	tmp_list = list
// 	if equal_grid(list, to_remove) {
// 		list = list.next
// 		tmp_list.next = nil
// 		return list
// 	}
// 	for tmp_list != nil {
// 		if equal_grid(tmp_list.next, to_remove) {
// 			buf := tmp_list.next
// 			tmp_list.next = tmp_list.next.next
// 			buf.next = nil
// 			return list
// 		}
// 		tmp_list = tmp_list.next
// 	}
// 	return list
// }

// func find_min_f(list *Node) *Node {
// 	var search *Node
// 	var tmp_list *Node

// 	tmp_list = list
// 	for tmp_f := 9999; tmp_list != nil; {
// 		if tmp_list.f < tmp_f {
// 			tmp_f = tmp_list.f
// 			search = tmp_list
// 		}
// 		tmp_list = tmp_list.next
// 	}
// 	return search
// }

func find_zero(grid [][]byte) (int, int) {
	for i, v := range grid {
		for j, vv := range v {
			if vv == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func create_4_child(grid_src [][]byte, g int, index int) *Node {
	
	i, j := find_zero(grid_src)

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
	return initial(grid, g)
}

func to_string(grid [][]byte) string {
	answer := ""
	for _, v := range grid {
		answer += string(v)
	}
	return answer
}

func a_implement(karta [][]byte) *Node {
	open_list := initial(karta, 0)

	// проверка на начальное состояние A + B = четное (есть решение)
	//								   A + B = нечетное (решения нет)
	// Число A — это число пар плиток, в которых плитка с большим 
	// номером идёт перед плиткой с меньшим номером (количество 
	// инверсий перестановок). Число B — это номер строки, в которой
	// находится пустое поле.
	closed_map := make(map[string]bool)
	for open_list != nil {
		//find_min_f
		// fmt.Println(open_list)
		process := open_list
		if process.h == 0 {
			return process
		}
		// open_list = open_list.remove(process)
		open_list = open_list.next
		// closed_list = closed_list.insert(process)
		closed_map[to_string(process.grid)] = true
		
		for i := 0; i < 4; i++ {
			if node := create_4_child(process.grid, process.g + 1, i); node != nil {
				node.parent = process
				// if closed_list.find_repeat(node) != nil {
				// 	continue
				// }
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


////////////////////////////////////////////////

func exist_solution(grid [][]byte) bool {
	size := len(grid)
	inv := 0
	// for i := 1; i < size * size; i++ {
		
	// 	buf = int(grid[(i - 1) / size][(i - 1) % size])
	// 	if buf == 0 {
	// 		inv += 1 + i / size
	// 		continue
	// 	}
	// 	j := i + 1
	// 	for ; j < size * size; j++ {
	// 		if buf > int(grid[(j - 1) / size][(j - 1) % size]) {
	// 			inv++
	// 		}
	// 	}
	// }
	zero := 0
	for i := 0; i < size * size; i++ {
		if grid[i / size][i % size] == 0 {
			// inv += 1 + i / size
			zero = ((size - 1) - i / size) + ((size - 1) - i % size)
			//zero = i + 1
		} else {
			for j := 0; j < i; j++ {
				if grid[j / size][j % size] > grid[i / size][i % size] {
					// fmt.Println(grid[j / size][j % size], grid[i / size][i % size])
					inv++
				}
			}
		}
	}
	//if size % 2 == 1 {
	//	// fmt.Println("aaa", zero, inv)
	//	return inv % 2 == 1
	//} else if (zero / size) % 2 == 0 {
	//	return inv % 2 == 1
	//}
	if inv % 2 == 0 && zero % 2 == 0 {
		return true
	} else if inv % 2 == 1 && zero % 2 == 1 {
		return true
	}
	
	return false
	//return inv % 2 == 0
}

func create_map(data []string) ([][]byte, error) {
	map_size, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}
	if map_size > 16 || map_size < 3 {
		return nil, errors.New(fmt.Sprintf("Bad map size"))
	}
	karta := make([][]byte, 0)
	cnt := 1
	if map_size * map_size != len(data) - 1 {
		return nil, errors.New(fmt.Sprintf("Wrong number of stroke"))
	}
	for i := 0; i < map_size; i++ {
		tmp := make([]byte, map_size)
		for j := 0; j < map_size; j++ {
			buf, err := strconv.Atoi(data[cnt])
			if err != nil {
				return nil, err
			}
			if buf < 0 || buf > map_size * map_size - 1 {
				return nil, errors.New(fmt.Sprintf("Bad character in map"))
			}
			tmp[j] = byte(buf)
			cnt++
		}
		karta = append(karta, tmp)
	}
	return karta, err
}

func skip_comment(for_valid_map []byte) []byte {
	answer := make([]byte, 0)

	for i := 0; i < len(for_valid_map); i++ {
		if for_valid_map[i] == '#' {
			for ; i < len(for_valid_map) && for_valid_map[i] != '\n'; i++{
			}
		} else {
			answer = append(answer, byte(for_valid_map[i]))
		}
	}
	return answer
}

func valid_digits(grid [][]byte) bool {
	size := len(grid)
	for i := 0; i < size * size - 1; i++ {
		tmp := grid[i / size][i % size]
		for j := i + 1; j < size * size; j++ {
			if tmp == grid[j / size][j % size] {
				return false
			}
		}
	}
	return true
}

func parse(fname string) ([][]byte, error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	buf := make([]byte, 64)
	map_str := make([]byte, 0)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		for i := 0; i < n; i++ {
			map_str = append(map_str, buf[i])
		}
	}
	if (len(map_str) == 0) {
		return nil, errors.New(fmt.Sprintf("Empty file"))
	}
	data := strings.Fields(string(skip_comment(map_str)))
	if len(data) == 0 {
		return nil, errors.New(fmt.Sprintf("No map"))
	}
	karta, err := create_map(data)
	if err != nil {
		return nil, err
	}
	if valid_digits(karta) == false {
		return nil, errors.New(fmt.Sprintf("Repeat digits in map"))
	}
	//if !exist_solution(karta){
	//	return nil, errors.New(fmt.Sprintf("Solution not exist"))
	//}
	return karta, nil
}

func duration(start time.Time) {
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	start := time.Now()
	defer duration(start)
	
	args := os.Args
	if len(args) == 2 {
		grid, err := parse(args[1])
		if err == nil {
			answer := a_implement(grid)
			//Печать ответа
			if answer != nil {
				for answer != nil {
					fmt.Println("g =", answer.g, "h =", answer.h, "f =", answer.f)
					for _, v := range answer.grid {
						fmt.Println(v)
					}
					fmt.Println("")
					answer = answer.parent
				}
			} else {
				fmt.Println("Solution not found")
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error number of args")
	}
	fmt.Println(create_goal(3))
}