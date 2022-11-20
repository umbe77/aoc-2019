package day2

import (
	"fmt"

	ic "github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	part1Resul := 0
	var intCode []int
	utils.ReadFile(fmt.Sprintf("./day2/input.txt"), func(line string) {
		intCode = ic.CompileIntCode(line)
	})
	intCodePart1 := make([]int, len(intCode))
	copy(intCodePart1, intCode)
	intCodePart1[1] = 12
	intCodePart1[2] = 2
	inputs := make([]int, 0)
	ic.RunIntCode(intCodePart1, inputs)
	part1Resul = intCodePart1[0]

	fmt.Printf("Part 1: %d\n", part1Resul)

	found := false
	var part2Result int
	for !found {
		for noun := 0; noun < len(intCode); noun++ {
			for verb := 0; verb < len(intCode); verb++ {
				icCopy := make([]int, len(intCode))
				copy(icCopy, intCode)
				icCopy[1] = noun
				icCopy[2] = verb
				ic.RunIntCode(icCopy, inputs)
				if icCopy[0] == 19690720 {
					found = true
					part2Result = 100*noun + verb
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2Result)
}
