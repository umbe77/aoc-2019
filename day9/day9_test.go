package day9_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/umbe77/aoc-2019/intcode"
)

func TestBOOST(t *testing.T) {
	input1 := "109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99"
	ic := intcode.New(input1)
	result := make([]string, 0)
	for *ic.Status == intcode.Running {
		result = append(result, strconv.Itoa(ic.Run([]int{})))
	}
	if strings.Join(result, ",") != input1+",0" {
		t.Errorf("Expected: '%s'\nGot:      '%s'\n", input1, strings.Join(result, ","))
	}

	input2 := "1102,34915192,34915192,7,4,7,99,0"
	ic = intcode.New(input2)
	result1 := ic.Run([]int{})
	if result1 != 1219070632396864 {
		t.Errorf("Expected: 1219070632396864, Got: '%d'\n", result1)
	}

	input3 := "104,1125899906842624,99"
	ic = intcode.New(input3)
	result3 := ic.Run([]int{})
	if result3 != 1125899906842624 {
		t.Errorf("Expected: 1125899906842624, Got: '%d'\n", result1)
	}
}
