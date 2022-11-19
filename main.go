package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2019/day1"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "a", "day in format 01")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	if day == "01" {
		day1.Execute()
	}

}
