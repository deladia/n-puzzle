package main

import (
	"fmt"
	"os"
	"errors"
	"io"
	"strconv"
	"strings"
)

func create_map(data []string) ([][]byte, error) {
	map_size, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}
	if map_size > 16 || map_size < 3 {
		return nil, errors.New(fmt.Sprintf("Bad map size\n"))
	}
	karta := make([][]byte, 0)
	cnt := 1
	if map_size * map_size != len(data) - 1 {
		return nil, errors.New(fmt.Sprintf("Wrong number of stroke\n"))
	}
	for i := 0; i < map_size; i++ {
		tmp := make([]byte, map_size)
		for j := 0; j < map_size; j++ {
			buf, err := strconv.Atoi(data[cnt])
			if err != nil {
				return nil, err
			}
			if buf < 0 || buf > map_size * map_size - 1 {
				return nil, errors.New(fmt.Sprintf("Bad character in map\n"))
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
		return nil, errors.New(fmt.Sprintf("Empty file\n"))
	}
	data := strings.Fields(string(skip_comment(map_str)))
	if len(data) == 0 {
		return nil, errors.New(fmt.Sprintf("No map\n"))
	}
	karta, err := create_map(data)
	if err != nil {
		return nil, err
	}
	if valid_digits(karta) == false {
		return nil, errors.New(fmt.Sprintf("Repeat digits in map\n"))
	}
	ideal := create_goal(len(karta))
	if exist_solution(karta, ideal) == false {
		return nil, errors.New(fmt.Sprintf("Solution not exist\n"))
	} else {
		fmt.Printf(Green, "Solution exist\n")
	}
	return karta, nil
}