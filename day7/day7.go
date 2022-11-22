package day7

import (
	"fmt"
	"sort"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {

	var ic intcode.IntCode
	utils.ReadFile(fmt.Sprintf("./day7/input.txt"), func(line string) {
		ic = intcode.New(line)
	})
	part1 := GetThruster1(ic)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := GetThruster2(ic.Copy())
	fmt.Printf("Part 2: %d\n", part2)
}

func RunAmplifier(ic intcode.IntCode, p1, p2 int) int {
	p := ic.Copy()
	p.InitInputs = []int{p1}
	inputs := []int{p2}
	return p.Run(inputs)
}
func RunAmplifier2(ic intcode.IntCode, p1, p2 int) int {
	inputs := []int{p1, p2}
	return ic.Run(inputs)
}

func GetThruster1(master intcode.IntCode) int {
	thrusters := make([]int, 0)
	for a := 0; a <= 4; a++ {
		outA := RunAmplifier(master, a, 0)
		for b := 0; b <= 4; b++ {
			if b == a {
				continue
			}
			outB := RunAmplifier(master, b, outA)
			for c := 0; c <= 4; c++ {
				if c == a || c == b {
					continue
				}
				outC := RunAmplifier(master, c, outB)
				for d := 0; d <= 4; d++ {
					if d == a || d == b || d == c {
						continue
					}
					outD := RunAmplifier(master, d, outC)
					for e := 0; e <= 4; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}
						outE := RunAmplifier(master, e, outD)
						thrusters = append(thrusters, outE)
					}
				}
			}
		}
	}
	sort.Ints(thrusters)
	return thrusters[len(thrusters)-1]
}
func GetThruster2(ic intcode.IntCode) int {
	thrusters := make([]int, 0)
	for a := 5; a <= 9; a++ {
		for b := 5; b <= 9; b++ {
			if b == a {
				continue
			}
			for c := 5; c <= 9; c++ {
				if c == a || c == b {
					continue
				}
				for d := 5; d <= 9; d++ {
					if d == a || d == b || d == c {
						continue
					}
					for e := 5; e <= 9; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}

						icA := ic.Copy()
						icA.InitInputs = []int{a}
						icB := ic.Copy()
						icB.InitInputs = []int{b}
						icC := ic.Copy()
						icC.InitInputs = []int{c}
						icD := ic.Copy()
						icD.InitInputs = []int{d}
						icE := ic.Copy()
						icE.InitInputs = []int{e}

						outA := 0
						outB := 0
						outC := 0
						outD := 0
						outE := 0
						lastOut := outE

						for *icA.Status == intcode.Running &&
							*icB.Status == intcode.Running &&
							*icC.Status == intcode.Running &&
							*icD.Status == intcode.Running &&
							*icE.Status == intcode.Running {

							outA = icA.Run([]int{outE})
							outB = icB.Run([]int{outA})
							outC = icC.Run([]int{outB})
							outD = icD.Run([]int{outC})
							outE = icE.Run([]int{outD})
							if outE != 0 {
								lastOut = outE
							}
						}

						thrusters = append(thrusters, lastOut)
					}
				}
			}
		}
	}
	sort.Ints(thrusters)
	return thrusters[len(thrusters)-1]
}
