package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
	"math"
)

type Node struct {
	grid	[][]byte
	g		int
	h		int
	f		int
	parent	*Node
	next	*Node
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
	new_node.g = g;
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

func exist_solution(grid [][]byte) bool {
	size := len(grid)
	inv := 0
	buf := 0
	for i := 1; i < size * size; i++ {
		
		buf = int(grid[(i - 1) / size][(i - 1) % size])
		if buf == 0 {
			inv += 1 + i / size
		}
		j := i + 1
		for ; j < size * size; j++ {
			if buf > int(grid[(j - 1) / size][(j - 1) % size]) && grid[(j - 1) / size][(j - 1) % size] != 0 {
				inv++
			}
		}
	}
	if inv % 2 == 0 {
		return true
	}
	return false
}

func a_implement(karta [][]byte) *Node {
	// if valid_start(karta) == false {
	// 	return nil
	// }

	// var open_list *Node
	var closed_list *Node

	// open_list = open_list.insert(&start)
	open_list := initial(karta, 0)
	// closed_list = nil

	// проверка на начальное состояние A + B = четное (есть решение)
	//								   A + B = нечетное (решения нет)
	// Число A — это число пар плиток, в которых плитка с большим 
	// номером идёт перед плиткой с меньшим номером (количество 
	// инверсий перестановок). Число B — это номер строки, в которой
	// находится пустое поле.
	for open_list != nil {
		//find_min_f
		process := open_list
		if process.h == 0 {
			return process
		}
		// open_list = open_list.remove(process)
		open_list = open_list.next
		closed_list = closed_list.insert(process)

		for i := 0; i < 4; i++ {
			if node := create_4_child(process.grid, process.g + 1, i); node != nil {
				node.parent = process
				if closed_list.find_repeat(node) != nil {
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
func create_map(data []string) ([][]byte, error) {
	map_size, err := strconv.Atoi(data[0])
	if err != nil || map_size > 16 {
		return nil, err
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
			//проверка buf на отрицательное и чтобы числа были нужные
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
			for ; for_valid_map[i] != '\n'; i++{
			}
		}
		answer = append(answer, byte(for_valid_map[i]))
	}
	return answer
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

	karta, err := create_map(data)
	if err != nil {
		return nil, err
	}
	if exist_solution(karta) == false {
		return nil, errors.New(fmt.Sprintf("Solution not exist"))
	}
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
				fmt.Println("Solution not exist")
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error number of args")
	}
}