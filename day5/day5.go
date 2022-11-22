package day5

import (
	"fmt"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var ic intcode.IntCode
	utils.ReadFile(fmt.Sprintf("./day5/input.txt"), func(line string) {
		ic = intcode.New(line)
	})

	var intCode1 = ic.Copy()
	var intCode2 = ic.Copy()

	inputs := make([]int, 1)
	inputs[0] = 1
	fmt.Println("Part 1: ")
	for *intCode1.Status == intcode.Running {
		fmt.Printf("%d\n", intCode1.Run(inputs))
		// intCode1.Run(inputs)
	}
	fmt.Println("---------")
	o2 := intCode2.Run([]int{5})
	fmt.Printf("Part 2: %d\n", o2)

}
