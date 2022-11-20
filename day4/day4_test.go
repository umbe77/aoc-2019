package day4_test

import (
	"testing"

	"github.com/umbe77/aoc-2019/day4"
)

func TestCheckPwd(t *testing.T) {
	if !day4.CheckPassword("111111") {
		t.Error("111111 shuld be good")
	}
	if !day4.CheckPassword("122345") {
		t.Error("122345 shuld be good")
	}
	if !day4.CheckPassword("111123") {
		t.Error("111123 shuld be good")
	}
	if day4.CheckPassword("223450") {
		t.Error("223450 shuld be bad")
	}
	if day4.CheckPassword("123789") {
		t.Error("123789 shuld be bad")
	}
}
func TestCheckPwd2(t *testing.T) {
	if !day4.CheckPassword2("112233") {
		t.Error("112233 shuld be good")
	}
	if !day4.CheckPassword2("111122") {
		t.Error("111122 shuld be good")
	}
	if day4.CheckPassword2("123444") {
		t.Error("123444 shuld be bad")
	}
}
