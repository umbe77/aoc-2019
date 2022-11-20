package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2019/day1"
	"github.com/umbe77/aoc-2019/day2"
)

func main() {
	var day string
	flag.StringVar(&day, "day", "a", "day in format 01")
	flag.Parse()

	fmt.Printf("Day %s\n", day)
	switch day {
	case "01":
		day1.Execute()
		break
	case "02":
		day2.Execute()
		break
	}
}
