package day13

import (
	"fmt"
	"sort"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var data string
	utils.ReadFile(fmt.Sprintf("./day13/input.txt"), func(line string) {
		data = line
	})

	fmt.Printf("Part 1: %d\n", Part1(data))
	fmt.Println("Part 2:")
	game := RunGame(data, true)
	DrawGame(game)
}

func Part1(code string) int {
	count := 0
	for _, v := range RunGame(code, false) {
		if v == Block {
			count++
		}
	}
	return count
}

func DrawGame(game map[Cell]TileKind) {
	xs := make([]int, 0)
	ys := make([]int, 0)
	for k := range game {
		xs = append(xs, k.X)
		ys = append(xs, k.Y)
	}
	sort.Ints(xs)
	sort.Ints(ys)
	maxY := ys[len(ys)-1]
	maxX := xs[len(xs)-1]

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			v := game[Cell{X: x, Y: y}]
			switch v {
			case Empty:
				fmt.Print(" ")
			case Wall:
				fmt.Print("█")
			case Block:
				fmt.Print(".")
			case Paddle:
				fmt.Print("-")
			case Ball:
				fmt.Print("o")
			}
		}
		fmt.Println()
	}

}

type TileKind int

const (
	Empty TileKind = iota
	Wall
	Block
	Paddle
	Ball
)

type Cell struct {
	X int
	Y int
}

func RunGame(code string, quarter bool) map[Cell]TileKind {
	game := make(map[Cell]TileKind)
	ic := intcode.New(code)
	if quarter {
		ic.Set(0, 2)
	}
	for *ic.Status == intcode.Running {
		x := ic.Run([]int{0})
		y := ic.Run([]int{1})
		id := ic.Run([]int{})

		if x == -1 && y == 0 {
			fmt.Println(id)
		}

		game[Cell{X: x, Y: y}] = TileKind(id)

	}
	return game
}
