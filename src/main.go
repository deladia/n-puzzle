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
	fmt.Println(end.Sub(start))
}

func main() {
	start := time.Now()
	defer duration(start)
	
	hueristicFlag := flag.String("h", "conflict", "hueristic: [hamming, manhattan, conflict]")
	flag.Parse()
	args := os.Args
	
	if len(args) == 2 || (len(args) == 4 && parse_flag(*hueristicFlag)){
		
		grid, err := parse(args[len(args) - 1])
		if err == nil {
			answer := a_implement(grid, hueristicFlag)
			if answer != nil {
				for answer != nil {
					fmt.Println("g =", answer.g, "h =", answer.h, "f =", answer.f)
					for _, v := range answer.grid {
						fmt.Println(v)
					}
					fmt.Println()
					answer = answer.parent
				}
			} else {
				fmt.Println("Solution not found")
			}
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Usage of ./n-puzzle:\n  -h string\n        hueristic: [hamming, manhattan, conflict] (default \"conflict\") file")
	}
}