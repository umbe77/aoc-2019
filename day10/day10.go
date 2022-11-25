package day10

import (
	"fmt"

	"github.com/umbe77/aoc-2019/utils"
)

var EPSILON float64 = 0.000001

func Execute() {
	matrix := make([]string, 0)
	utils.ReadFile(fmt.Sprintf("./day10/input.txt"), func(line string) {
		matrix = append(matrix, line)
	})

	// matrix1 := []string{
	// 	".#..##.###...#######",
	// 	"##.############..##.",
	// 	".#.######.########.#",
	// 	".###.#######.####.#.",
	// 	"#####.##.#.##.###.##",
	// 	"..#####..#.#########",
	// 	"####################",
	// 	"#.####....###.#.#.##",
	// 	"##.#################",
	// 	"#####.##.###..####..",
	// 	"..######..##.#######",
	// 	"####.##.####...##..#",
	// 	".#####..#.######.###",
	// 	"##...#.##########...",
	// 	"#.##########.#######",
	// 	".####.#.###.###.#.##",
	// 	"....##.##.###..#####",
	// 	".#.#.###########.###",
	// 	"#.#.#.#####.####.###",
	// 	"###.##.####.##.#..##",
	// }
	asteroids := ParseMapLine(matrix)
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

type Asteroid struct {
	X int
	Y int
}

func BestAsteroid(asteroids []Asteroid) int {
	maxCountVisible := 0
	for _, asteroid := range asteroids {
		leftQuadrant, rightQuadrant := make(map[float64]bool), make(map[float64]bool)
		var checkVerticalUp, checkVerticalDown bool
		count := 0
		for _, other := range asteroids {
			if asteroid != other {
				dy := other.Y - asteroid.Y
				dx := other.X - asteroid.X
				if dx == 0 {
					if dy > 0 {
						if !checkVerticalUp {
							checkVerticalUp = true
							count++
						}
					} else {
						if !checkVerticalDown {
							checkVerticalDown = true
							count++
						}
					}
				} else {
					ratio := float64(dy) / float64(dx)
					if dx < 0 {
						if _, inLeft := leftQuadrant[ratio]; !inLeft {
							leftQuadrant[ratio] = true
							count++
						}
					} else {
						if _, inRight := rightQuadrant[ratio]; !inRight {
							rightQuadrant[ratio] = true
							count++
						}
					}
				}
			}
		}
		if maxCountVisible < count {
			maxCountVisible = count
		}
	}

	return maxCountVisible
}
