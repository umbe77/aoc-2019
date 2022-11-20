package day4

import (
	"fmt"
	"strconv"
)

func Execute() {
	resultPart1 := run(367479, 893698, CheckPassword)
	fmt.Printf("Part 1: %d\n", resultPart1)
	resultPart2 := run(367479, 893698, CheckPassword2)
	fmt.Printf("Part 2: %d\n", resultPart2)
}

func run(a, b int, check func(pwd string) bool) int {
	count := 0
	for i := a; i <= b; i++ {
		if check(strconv.Itoa(i)) {
			count++
		}
	}
	return count
}

func CheckPassword(pwd string) bool {
	hasDouble := false
	for i := 1; i < len(pwd); i++ {
		digit, _ := strconv.Atoi(string(pwd[i]))
		prev, _ := strconv.Atoi(string(pwd[i-1]))
		if prev > digit {
			return false
		}
		if !hasDouble && prev == digit {
			hasDouble = true
		}
	}
	return hasDouble
}

func CheckPassword2(pwd string) bool {
	for i := 1; i < len(pwd); i++ {
		digit, _ := strconv.Atoi(string(pwd[i]))
		prev, _ := strconv.Atoi(string(pwd[i-1]))
		if prev > digit {
			return false
		}
	}
	i := 0
	for i < len(pwd)-1 {
		digit, _ := strconv.Atoi(string(pwd[i]))
		countEquals := 1
		for j := i + 1; j < len(pwd); j++ {
			next, _ := strconv.Atoi(string(pwd[j]))
			if next > digit {
				i = j
				break
			}
			countEquals++
			if j == len(pwd)-1 {
				i = len(pwd)
			}
		}
		if countEquals == 2 {
			return true
		}
	}
	return false
}
