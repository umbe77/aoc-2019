package day13

import (
	"fmt"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var data string
	utils.ReadFile(fmt.Sprintf("./day13/input.txt"), func(line string) {
		data = line
	})

	fmt.Printf("Part 1: %d\n", Part1(data))
}

type TileKind int

const (
	Empty TileKind = iota
	Wall
	Block
	Paddle
	Ball
)

func Part1(code string) int {
	ic := intcode.New(code)
	count := 0
	for *ic.Status == intcode.Running {
		ic.Run([]int{})
		ic.Run([]int{})
		id := ic.Run([]int{})

		if id == int(Block) {
			count++
		}

	}
	return count
}
