package day2_test

import (
	"testing"

	"github.com/umbe77/aoc-2019/intcode"
)

func TestRunIntCode(t *testing.T) {

	ic := intcode.New("1,1,1,4,99,5,6,0,99")
	ic.Run([]int{})
	if ic.Get(0) != 30 {
		t.Errorf("Expect 30 got %d", ic.Get(0))
	}
}
