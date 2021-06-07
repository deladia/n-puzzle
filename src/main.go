package main

import (
	"flag"
	"fmt"
	"os"
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

func duration(start time.Time) {
	end := time.Now()
	fmt.Printf(Yellow + Yellow, end.Sub(start), " for solution\n\n")
}

func main() {
	start := time.Now()
	
	hueristic_flag := flag.String("h", "conflict", "hueristic: [hamming, manhattan, conflict]")
	search_algo := flag.String("s", "a_star", "search algorithm: [a_star, greedy, uniform]")
	flag.Parse()
	args := os.Args
	
	if len(args) == 2 || len(args) == 4 || len(args) == 6 {
		if parse_flag_h(*hueristic_flag) == false || parse_flag_s(*search_algo) == false {
			fmt.Printf(Yellow, "Usage of ./n-puzzle:\n  -h string\n        hueristic: [hamming, manhattan, conflict] (default \"conflict\") file\n")
		}
		grid, err := parse(args[len(args) - 1])
		if err == nil {
			var answer *Node
			switch *search_algo {
			case "a_star":
				answer = a_implement(grid, hueristic_flag)
			case "greedy":
				answer = greedy_search(grid, hueristic_flag)
			case "uniform":
				answer = uniformcost_search(grid)
			}
			duration(start)
			if answer != nil {
				answer.print_list()
			} else {
				fmt.Printf(Red, "Solution not found\n")
			}
		} else {
			fmt.Printf(Red, err)
		}
	} else {
		fmt.Printf(Yellow, "Usage of ./n-puzzle:\n  -h string\n        hueristic: [hamming, manhattan, conflict] (default \"conflict\") file\n")
	}
}