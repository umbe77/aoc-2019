package day1_test

import (
	"testing"

	"github.com/umbe77/aoc-2019/day1"
)

func TestGetModuleMass(t *testing.T) {

	if fuel := day1.GetModuleMass(12); fuel != 2 {
		t.Errorf("Expected 2: got %d", fuel)
	}

	if fuel := day1.GetModuleMass(14); fuel != 2 {
		t.Errorf("Expected 2: got %d", fuel)
	}
	if fuel := day1.GetModuleMass(1969); fuel != 654 {
		t.Errorf("Expected 654: got %d", fuel)
	}
	if fuel := day1.GetModuleMass(100756); fuel != 33583 {
		t.Errorf("Expected 33583: got %d", fuel)
	}
}

func TestGetModuleMass2(t *testing.T) {

	if fuel := day1.GetModuleMass2(14); fuel != 2 {
		t.Errorf("Expected 2: got %d", fuel)
	}
	if fuel := day1.GetModuleMass2(1969); fuel != 966 {
		t.Errorf("Expected 966: got %d", fuel)
	}
	if fuel := day1.GetModuleMass2(100756); fuel != 50346 {
		t.Errorf("Expected 50346: got %d", fuel)
	}
}
