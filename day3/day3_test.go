package day3_test

import (
	"testing"

	"github.com/umbe77/aoc-2019/day3"
)

func getNearest(dir1, dir2 []day3.Direction) (int, int) {

	wire1 := day3.GetPoints(dir1)
	wire2 := day3.GetPoints(dir2)

	crossing := day3.FindCrossingPoints(wire1, wire2)
	distance := day3.FindNearCrossingPoint(crossing)
	step := day3.FindBestStep(crossing)
	return distance, step
}

func TestFindNearest1(t *testing.T) {
	wire1Directions := day3.ParseDirs("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	wire2Directions := day3.ParseDirs("U62,R66,U55,R34,D71,R55,D58,R83")

	distance, _ := getNearest(wire1Directions, wire2Directions)

	if distance != 159 {
		t.Errorf("Expected 159 got %d", distance)
	}
}
func TestFindNearest2(t *testing.T) {
	wire1Directions := day3.ParseDirs("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	wire2Directions := day3.ParseDirs("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")

	distance, _ := getNearest(wire1Directions, wire2Directions)

	if distance != 135 {
		t.Errorf("Expected 135 got %d", distance)
	}
}

func TestFindSteps1(t *testing.T) {
	wire1Directions := day3.ParseDirs("R75,D30,R83,U83,L12,D49,R71,U7,L72")
	wire2Directions := day3.ParseDirs("U62,R66,U55,R34,D71,R55,D58,R83")

	_, steps := getNearest(wire1Directions, wire2Directions)

	if steps != 610 {
		t.Errorf("Expected 610 got %d", steps)
	}
}
func TestFindSteps2(t *testing.T) {
	wire1Directions := day3.ParseDirs("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	wire2Directions := day3.ParseDirs("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")

	_, steps := getNearest(wire1Directions, wire2Directions)

	if steps != 410 {
		t.Errorf("Expected 410 got %d", steps)
	}
}
