package day11

import (
	"fmt"
	"sort"

	"github.com/umbe77/aoc-2019/intcode"
	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var data string
	utils.ReadFile(fmt.Sprintf("./day11/input.txt"), func(line string) {
		data = line
	})
	count, _ := Paint(data, 0)
	fmt.Printf("Part 1: %d\n", count)
	_, tiles := Paint(data, 1)
	fmt.Println("Part 2:")
	Draw(tiles)
}

type Tile struct {
	X int
	Y int
}

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Ehpr struct {
	Dir         Direction
	CurrentTile Tile
}

var emptyTile Tile

func MoveEhpr(direction int, robot *Ehpr) {
	switch direction {
	case 0: //rotate left
		d := (robot.Dir - 1) % 4
		if d < 0 {
			d = d + 4
		}
		robot.Dir = d
	case 1: //rotate right
		robot.Dir = (robot.Dir + 1) % 4
	}
	switch robot.Dir {
	case Up:
		robot.CurrentTile.Y -= 1
	case Left:
		robot.CurrentTile.X -= 1
	case Down:
		robot.CurrentTile.Y += 1
	case Right:
		robot.CurrentTile.X += 1
	}
}

func FindTile(tiles []Tile, current Tile) int {
	for i, t := range tiles {
		if t.X == current.X && t.Y == current.Y {
			return i
		}
	}
	return -1
}

func Paint(code string, firstColor int) (int, map[Tile]int) {
	ic := intcode.New(code)
	tiles := make(map[Tile]int)
	tiles[Tile{X: 0, Y: 0}] = firstColor
	robot := Ehpr{
		Dir:         Up,
		CurrentTile: Tile{X: 0, Y: 0},
	}
	for *ic.Status == intcode.Running {
		tileColor, ok := tiles[robot.CurrentTile]
		if !ok {
			tileColor = 0
		}

		paintColor := ic.Run([]int{tileColor})
		tiles[robot.CurrentTile] = paintColor
		dir := ic.Run([]int{})
		MoveEhpr(dir, &robot)
	}
	return len(tiles), tiles
}

func Draw(tiles map[Tile]int) {
	xs := make([]int, 0)
	ys := make([]int, 0)
	for k := range tiles {
		xs = append(xs, k.X)
		ys = append(ys, k.Y)
	}
	sort.Ints(xs)
	sort.Ints(ys)

	for y := ys[0] - 2; y <= ys[len(ys)-1]+2; y++ {
		for x := xs[0] - 2; x <= xs[len(xs)-1]+2; x++ {
			colorToPaint := tiles[Tile{X: x, Y: y}]
			color := " "
			if colorToPaint == 1 {
				color = "█"
			}
			fmt.Print(color)
		}
		fmt.Println()
	}

}
