package day6_test

import (
	"fmt"
	"testing"

	"github.com/umbe77/aoc-2019/day6"
)

func TestCount(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}

	com := day6.InitializeOrbits(input)
	count := day6.CountOrbits(com)
	if count != 42 {
		t.Errorf("Expected 42 got %d", count)
	}
}

func TestShortest(t *testing.T) {
	input := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}

	com := day6.InitializeOrbitsGraph(input)
	shortest := day6.FindShortest(com, "YOU", "SAN", []string{})
	fmt.Printf("%q\n", shortest)
	orbitalTranfers := len(shortest) - 3
	if orbitalTranfers != 4 {
		t.Errorf("Expected 4, Got %d", orbitalTranfers)
	}
}
