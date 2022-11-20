package day3

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/umbe77/aoc-2019/utils"
)

func Execute() {
	var directions [][]Direction
	utils.ReadFile(fmt.Sprintf("./day3/input.txt"), func(line string) {
		directions = append(directions, ParseDirs(line))
	})

	wire1 := GetPoints(directions[0])
	wire2 := GetPoints(directions[1])

	points := FindCrossingPoints(wire1, wire2)
	part1Result := FindNearCrossingPoint(points)

	fmt.Printf("Part 1: %d\n", part1Result)

	part2Result := FindBestStep(points)
	fmt.Printf("Part 2: %d\n", part2Result)
}

type Point struct {
	X      int
	Y      int
	Weight int
}
type Direction struct {
	Dir    string
	Length int
}

func FindNearCrossingPoint(points []Point) int {
	distances := make([]int, len(points))
	for i, v := range points {
		distances[i] = utils.Abs(v.X) + utils.Abs(v.Y)
	}
	sort.Ints(distances)
	return distances[0]
}

func FindBestStep(points []Point) int {
	steps := make([]int, len(points))

	for i, v := range points {
		steps[i] = v.Weight
	}

	sort.Ints(steps)
	return steps[0]
}

func FindCrossingPoints(wire1, wire2 []Point) []Point {
	crossingPoints := make([]Point, 0)
	for i := 1; i < len(wire1); i++ {
		for j := 1; j < len(wire2); j++ {
			w1_p0 := wire1[i-1]
			w1_p1 := wire1[i]
			w2_p0 := wire2[j-1]
			w2_p1 := wire2[j]

			if isVertical(w1_p0, w1_p1) && isVertical(w2_p0, w2_p1) ||
				isHorizontal(w1_p0, w1_p1) && isHorizontal(w2_p0, w2_p1) {
				continue
			}
			if isVertical(w1_p0, w1_p1) { //w2 segment is actually horizontal
				if w1_p1.X > utils.Min(w2_p0.X, w2_p1.X) && w1_p1.X < utils.Max(w2_p0.X, w2_p1.X) &&
					w2_p0.Y > utils.Min(w1_p0.Y, w1_p1.Y) && w2_p0.Y < utils.Max(w1_p0.Y, w1_p1.Y) {

					wire1Weight := w1_p1.Weight - utils.Abs(w1_p1.Y-w2_p0.Y)
					wire2Weight := w2_p1.Weight - utils.Abs(w2_p1.X-w1_p0.X)
					weight := wire1Weight + wire2Weight
					crossingPoints = append(crossingPoints, Point{
						X:      w1_p0.X,
						Y:      w2_p0.Y,
						Weight: weight,
					})
				}
			}
			if isHorizontal(w1_p0, w1_p1) { //w2 segment is actually vertical
				if w1_p1.Y > utils.Min(w2_p0.Y, w2_p1.Y) && w1_p1.Y < utils.Max(w2_p0.Y, w2_p1.Y) &&
					w2_p0.X > utils.Min(w1_p0.X, w1_p1.X) && w2_p0.X < utils.Max(w1_p0.X, w1_p1.X) {

					wire1Weight := w1_p1.Weight - utils.Abs(w1_p1.X-w2_p0.X)
					wire2Weight := w2_p1.Weight - utils.Abs(w2_p1.Y-w1_p0.Y)
					weight := wire1Weight + wire2Weight
					crossingPoints = append(crossingPoints, Point{
						X:      w2_p0.X,
						Y:      w1_p0.Y,
						Weight: weight,
					})
				}
			}
		}
	}
	return crossingPoints
}

func isVertical(p0, p1 Point) bool {
	return p0.X == p1.X
}
func isHorizontal(p0, p1 Point) bool {
	return p0.Y == p1.Y
}

func GetPoints(dirs []Direction) []Point {
	points := make([]Point, 0)
	currentPoint := Point{X: 0, Y: 0, Weight: 0}
	points = append(points, currentPoint)

	for _, d := range dirs {
		dest := Point{
			X:      currentPoint.X,
			Y:      currentPoint.Y,
			Weight: currentPoint.Weight,
		}
		switch d.Dir {
		case "UP":
			dest.Y = dest.Y + d.Length
		case "DOWN":
			dest.Y = dest.Y - d.Length
		case "LEFT":
			dest.X = dest.X - d.Length
		case "RIGHT":
			dest.X = dest.X + d.Length
		}
		dest.Weight = dest.Weight + d.Length
		points = append(points, dest)
		currentPoint = dest
	}
	return points
}

func ParseDirs(line string) []Direction {
	directions := make([]Direction, 0)
	for _, v := range strings.Split(line, ",") {
		dir := Direction{}
		switch string(v[0]) {
		case "U":
			dir.Dir = "UP"
		case "D":
			dir.Dir = "DOWN"
		case "L":
			dir.Dir = "LEFT"
		case "R":
			dir.Dir = "RIGHT"
		}
		l, err := strconv.Atoi(v[1:])
		if err != nil {
			log.Fatal(err)
		}
		dir.Length = l
		directions = append(directions, dir)
	}
	return directions
}
