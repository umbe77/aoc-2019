package day2

import (
	"fmt"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var ic intcode.IntCode
	utils.ReadFile(fmt.Sprintf("./day2/input.txt"), func(line string) {
		ic = intcode.New(line)
	})

	ic1 := ic.Copy()
	ic1.Set(1, 12)
	ic1.Set(2, 2)
	ic1.Run([]int{})
	part1Result := ic1.Get(0)

	fmt.Printf("Part 1: %d\n", part1Result)

	found := false
	var part2Result int
	for !found {
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				icCopy := ic.Copy()
				icCopy.Set(1, noun)
				icCopy.Set(2, verb)
				icCopy.Run([]int{})
				if icCopy.Get(0) == 19690720 {
					found = true
					part2Result = 100*noun + verb
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2Result)
}
