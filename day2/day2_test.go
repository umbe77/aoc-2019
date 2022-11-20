package day2_test

import (
	"fmt"
	"testing"

	ic "github.com/umbe77/aoc-2019/intcode"
)

func TestRunIntCode(t *testing.T) {

	intCode := ic.CompileIntCode("1,1,1,4,99,5,6,0,99")
	fmt.Printf("%v\n", intCode)
	inputs := make([]int, 0)
	ic.RunIntCode(intCode, inputs)
	if intCode[0] != 30 {
		t.Errorf("Expect 30 got %d", intCode[0])
	}
}
