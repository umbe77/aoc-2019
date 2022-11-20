package day5

import (
	"fmt"

	ic "github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var intCode []int
	utils.ReadFile(fmt.Sprintf("./day5/input.txt"), func(line string) {
		intCode = ic.CompileIntCode(line)
	})
	var intCode1 = make([]int, len(intCode))
	copy(intCode1, intCode)
	var intCode2 = make([]int, len(intCode))
	copy(intCode2, intCode)

	inputs := make([]int, 1)
	inputs[0] = 1
	outputs := ic.RunIntCode(intCode1, inputs)
	fmt.Println("Part 1: ")
	for i, v := range outputs {
		fmt.Printf("%d : %d\n", i, v)
	}
	fmt.Println("---------")
	o2 := ic.RunIntCode(intCode2, []int{5})
	fmt.Printf("Part 2: %d\n", o2[0])

}
