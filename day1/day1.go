package day1

import (
	"fmt"
	"log"
	"strconv"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	sumPart1 := 0
	sumPart2 := 0
	utils.ReadFile(fmt.Sprintf("./day1/input.txt"), func(line string) {
		mass, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		sumPart1 += GetModuleMass(mass)
		sumPart2 += GetModuleMass2(mass)
	})
	fmt.Printf("Part 1: %d\n", sumPart1)
	fmt.Printf("Part 2: %d\n", sumPart2)
}

func GetModuleMass(mass int) int {
	return (mass / 3) - 2
}

func GetModuleMass2(mass int) int {
	sum := 0
	current := mass
	nonZero := true
	for nonZero {
		current = GetModuleMass(current)
		nonZero = current > 0
		if nonZero {
			sum += current
		}
	}
	return sum
}
