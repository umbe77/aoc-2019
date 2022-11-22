package day5_test

import (
	"testing"

	ic "github.com/umbe77/aoc-2019/intcode"
)

func TestEngine(t *testing.T) {
	code := "3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99"
	ic1 := ic.New(code)
	i1 := []int{3}
	o1 := ic1.Run(i1)
	if o1 != 999 {
		t.Errorf("Expected 999 got %v", o1)
	}

	ic1 = ic.New(code)
	i1_1 := []int{8}
	o1_1 := ic1.Run(i1_1)
	if o1_1 != 1000 {
		t.Errorf("Expected 1000 got %v", o1_1)
	}

	ic1 = ic.New(code)
	i1_2 := []int{23}
	o1_2 := ic1.Run(i1_2)
	if o1_2 != 1001 {
		t.Errorf("Expected 1001 got %v", o1_2)
	}
}

func TestPart2Jump(t *testing.T) {
	code1 := "3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9"
	ic1 := ic.New(code1)
	i1 := []int{3}
	o1 := ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code1)
	i1_1 := []int{0}
	o1_1 := ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected 0 got %v", o1_1)
	}
	code2 := "3,3,1105,-1,9,1101,0,0,12,4,12,99,1"
	ic1 = ic.New(code2)
	i1 = []int{3}
	o1 = ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code2)
	i1_1 = []int{0}
	o1_1 = ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected 0 got %v", o1_1)
	}
}

func TestPart2Compare(t *testing.T) {
	code1 := "3,9,8,9,10,9,4,9,99,-1,8"
	ic1 := ic.New(code1)
	i1 := []int{8}
	o1 := ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code1)
	i1_1 := []int{3}
	o1_1 := ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected 0 got %v", o1_1)
	}

	code2 := "3,9,7,9,10,9,4,9,99,-1,8"
	ic1 = ic.New(code2)
	i1 = []int{4}
	o1 = ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code2)
	i1_1 = []int{8}
	o1_1 = ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected lessthen 0 got %v", o1_1)
	}

	code3 := "3,3,1108,-1,8,3,4,3,99"
	ic1 = ic.New(code3)
	i1 = []int{8}
	o1 = ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code3)
	i1_1 = []int{3}
	o1_1 = ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected 0 got %v", o1_1)
	}

	code4 := "3,3,1107,-1,8,3,4,3,99"
	ic1 = ic.New(code4)
	i1 = []int{4}
	o1 = ic1.Run(i1)
	if o1 != 1 {
		t.Errorf("Expected 1 got %v", o1)
	}
	ic1 = ic.New(code4)
	i1_1 = []int{8}
	o1_1 = ic1.Run(i1_1)
	if o1_1 != 0 {
		t.Errorf("Expected lessthen 0 got %v", o1_1)
	}
}
