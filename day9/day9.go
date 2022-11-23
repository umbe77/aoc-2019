package day9

import (
	"fmt"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var data string
	utils.ReadFile(fmt.Sprintf("./day9/input.txt"), func(line string) {
		data = line
	})
	fmt.Printf("Part 1: %d\n", Boost(data, 1))
	fmt.Printf("Part 2: %d\n", Boost(data, 2))
}

func Boost(boost string, input int) int {
	ic := intcode.New(boost)
	return ic.Run([]int{input})
}
