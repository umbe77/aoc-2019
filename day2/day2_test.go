package day2_test

import (
	"fmt"
	"testing"

	"github.com/umbe77/aoc-2019/day2"
)

func TestRunIntCode(t *testing.T) {

	intCode := day2.ConvertIntCode("1,1,1,4,99,5,6,0,99")
	fmt.Printf("%v\n", intCode)
	if result := day2.RunIntCode(intCode); result != 30 {
		t.Errorf("Expect 30 got %d", result)
	}
}
