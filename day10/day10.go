package day10

import (
	"fmt"
	"sort"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	matrix := make([]string, 0)
	utils.ReadFile(fmt.Sprintf("./day10/input.txt"), func(line string) {
		matrix = append(matrix, line)
	})

	matrix1 := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	asteroids := ParseMapLine(matrix1)
	fmt.Printf("Part 1: %d\n", BestAsteroid(asteroids))
}

func ParseMapLine(matrix []string) []Asteroid {
	asteroids := make([]Asteroid, 0)
	for y, l := range matrix {
		for x, p := range l {
			a := string(p)
			if a == "#" {
				asteroids = append(asteroids, Asteroid{
					X: x,
					Y: y,
				})
			}
		}
	}
	return asteroids
}

func checkOnSameLine(a, b, c Asteroid) bool {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dx1 := a.X - c.X
	dy1 := a.Y - c.Y
	return dx1*dy == dy1*dx
}

type Asteroid struct {
	X int
	Y int
}

func isBetween(a, p1, p2 Asteroid) bool {
	dxl := p1.X - p2.X
	dyl := p1.Y - p2.Y
	if utils.Abs(dxl) >= utils.Abs(dyl) {
		if dxl > 0 {
			return p1.X <= a.X && a.X <= p2.X
		} else {
			return p2.X <= a.X && a.X <= p1.X
		}
	} else {
		if dyl > 0 {
			return p1.Y <= a.Y && a.Y <= p2.Y
		} else {
			return p2.Y <= a.Y && a.Y <= p1.Y
		}
	}
}

func BestAsteroid(asteroids []Asteroid) int {
	locations := make(map[string][]Asteroid)
	for _, a := range asteroids {
		name := fmt.Sprintf("%d%d", a.X, a.Y)
		locations[name] = make([]Asteroid, 0)
		inLine := make([]Asteroid, 0)
		for _, b := range asteroids {
			if a.X == b.X && a.Y == a.Y {
				continue
			}
			inLine = append(inLine, b)
			for _, c := range asteroids {
				if (a.X == c.X && a.Y == c.Y) ||
					b.X == c.X && b.Y == c.Y {
					continue
				}
				if checkOnSameLine(a, b, c) {
					inLine = append(inLine, c)
				}
			}
		}
		sort.Slice(inLine, func(i, j int) bool {
			return inLine[i].X <= inLine[j].X && inLine[i].Y <= inLine[j].Y
		})
		if a.X <= inLine[0].X && a.Y <= inLine[0].Y {
			locations[name] = append(locations[name], inLine[0])
		} else if a.X >= inLine[len(inLine)-1].X && a.Y >= inLine[len(inLine)-1].Y {
			locations[name] = append(locations[name], inLine[len(inLine)-1])
		} else {
			for i := 1; i < len(inLine); i++ {
				p1 := inLine[i-1]
				p2 := inLine[i]
				if isBetween(a, p1, p2) {
					locations[name] = append(locations[name], p1)
					locations[name] = append(locations[name], p2)
					break
				}
			}
		}
	}
	fmt.Printf("%+v\n", locations)
	astList := make([]int, 0)
	for _, l := range locations {
		astList = append(astList, len(l))
	}
	sort.Ints(astList)
	return astList[len(astList)-1]
}
