package main

import (
	"flag"
	"fmt"

	"github.com/umbe77/aoc-2019/day1"
	"github.com/umbe77/aoc-2019/day2"
	"github.com/umbe77/aoc-2019/day3"
	"github.com/umbe77/aoc-2019/day4"
	"github.com/umbe77/aoc-2019/day5"
	"github.com/umbe77/aoc-2019/day6"
	"github.com/umbe77/aoc-2019/day7"
	"github.com/umbe77/aoc-2019/day8"
	"github.com/umbe77/aoc-2019/day9"
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
	case "03":
		day3.Execute()
		break
	case "04":
		day4.Execute()
		break
	case "05":
		day5.Execute()
		break
	case "06":
		day6.Execute()
		break
	case "07":
		day7.Execute()
		break
	case "08":
		day8.Execute()
		break
	case "09":
		day9.Execute()
		break
	}
}
