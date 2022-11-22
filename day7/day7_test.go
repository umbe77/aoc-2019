package day7_test

import (
	"github.com/umbe77/aoc-2019/day7"
	ic "github.com/umbe77/aoc-2019/intcode"
	"testing"
)

func TestThrusterMax(t *testing.T) {
	input1 := "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0"
	code1 := ic.New(input1)
	t1 := day7.GetThruster1(code1)
	if t1 != 43210 {
		t.Errorf("Expected 43210, Got %d", t1)
	}

	input2 := "3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0"
	code2 := ic.New(input2)
	t2 := day7.GetThruster1(code2)
	if t2 != 54321 {
		t.Errorf("Expected 54321, Got %d", t2)
	}

	input3 := "3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0"
	code3 := ic.New(input3)
	t3 := day7.GetThruster1(code3)
	if t3 != 65210 {
		t.Errorf("Expected 65210, Got %d", t3)
	}
}
func TestThrusterMax2(t *testing.T) {
	input1 := "3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5"
	code1 := ic.New(input1)
	t1 := day7.GetThruster2(code1)
	if t1 != 139629729 {
		t.Errorf("Expected 139629729, Got %d", t1)
	}

	input2 := "3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10"
	code2 := ic.New(input2)
	t2 := day7.GetThruster2(code2)
	if t2 != 18216 {
		t.Errorf("Expected 18216, Got %d", t2)
	}
}
