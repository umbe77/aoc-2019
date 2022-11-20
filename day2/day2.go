package day2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	part1Resul := 0
	var intCode []int
	utils.ReadFile(fmt.Sprintf("./day2/input.txt"), func(line string) {
		intCode = ConvertIntCode(line)
	})
	intCodePart1 := make([]int, len(intCode))
	copy(intCodePart1, intCode)
	intCodePart1[1] = 12
	intCodePart1[2] = 2
	part1Resul = RunIntCode(intCodePart1)

	fmt.Printf("Part 1: %d\n", part1Resul)

	found := false
	var part2Result int
	for !found {
		for noun := 0; noun < len(intCode); noun++ {
			for verb := 0; verb < len(intCode); verb++ {
				ic := make([]int, len(intCode))
				copy(ic, intCode)
				ic[1] = noun
				ic[2] = verb
				if RunIntCode(ic) == 19690720 {
					found = true
					part2Result = 100*noun + verb
				}
			}
		}
	}
	fmt.Printf("Part 2: %d\n", part2Result)
}

func ConvertIntCode(input string) []int {

	strIntCode := strings.Split(input, ",")
	intCode := make([]int, len(strIntCode))

	for i, v := range strIntCode {
		iv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intCode[i] = iv
	}
	return intCode
}

func RunIntCode(intCode []int) int {
	instructinPointer := 0
	opCode := 0
	running := true
	for running {
		opCode = intCode[instructinPointer]
		if opCode == 99 {
			break
		}

		switch opCode {
		case 1: //Add
			lo := intCode[instructinPointer+1]
			ro := intCode[instructinPointer+2]
			destination := intCode[instructinPointer+3]
			intCode[destination] = intCode[lo] + intCode[ro]
			instructinPointer += 4
			break
		case 2: //Multiply
			lo := intCode[instructinPointer+1]
			ro := intCode[instructinPointer+2]
			destination := intCode[instructinPointer+3]
			intCode[destination] = intCode[lo] * intCode[ro]
			instructinPointer += 4
			break
		case 99:
			running = false
			break
		default:
			log.Fatalf("Unkown OpCode: %d", opCode)
		}
	}
	return intCode[0]
}
