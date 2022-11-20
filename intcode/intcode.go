package intcode

import (
	"log"
	"strconv"
	"strings"
)

func CompileIntCode(input string) []int {

	strIntCode := strings.Split(input, ",")
	intCode := make([]int, len(strIntCode))

	for i, v := range strIntCode {
		iv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		intCode[i] = iv
	}
	return intCode
}

type OpCode struct {
	Code       int
	Param1Mode int
	Param2Mode int
	Param3Mode int
}

func getOpCode(code int) OpCode {
	opCode := OpCode{
		Code:       99,
		Param1Mode: 0,
		Param2Mode: 0,
		Param3Mode: 0,
	}
	s := strconv.Itoa(code)
	switch len(s) {
	case 1:
		opCode.Code = code
	case 2:
		opCode.Code = code
	case 3:
		opCode.Param1Mode, _ = strconv.Atoi(string(s[0]))
		opCode.Code, _ = strconv.Atoi(string(s[1:]))
	case 4:
		opCode.Param1Mode, _ = strconv.Atoi(string(s[1]))
		opCode.Param2Mode, _ = strconv.Atoi(string(s[0]))
		opCode.Code, _ = strconv.Atoi(string(s[2:]))
	case 5:
		opCode.Param1Mode, _ = strconv.Atoi(string(s[2]))
		opCode.Param2Mode, _ = strconv.Atoi(string(s[1]))
		opCode.Param3Mode, _ = strconv.Atoi(string(s[0]))
		opCode.Code, _ = strconv.Atoi(string(s[3:]))
	}
	return opCode
}

func RunIntCode(intCode []int, inputs []int) []int {
	outputs := make([]int, 0)
	instructionPointer := 0
	opCode := OpCode{}
	currentInputPointer := 0
	var getInputValue = func() int {
		v := inputs[currentInputPointer]
		currentInputPointer++
		return v
	}
	var getParam = func(p, mode int) int {
		if mode == 0 {
			return intCode[p]
		}
		return p
	}
	running := true
	for running {
		opCode = getOpCode(intCode[instructionPointer])
		if opCode.Code == 99 {
			break
		}

		switch opCode.Code {
		case 1: //Add
			lo := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			ro := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			destination := intCode[instructionPointer+3]
			intCode[destination] = lo + ro
			instructionPointer += 4
			break
		case 2: //Multiply
			lo := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			ro := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			destination := intCode[instructionPointer+3]
			intCode[destination] = lo * ro
			instructionPointer += 4
			break
		case 3: //Inputs
			destination := intCode[instructionPointer+1]
			intCode[destination] = getInputValue()
			instructionPointer += 2
		case 4: //Outputs
			source := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			outputs = append(outputs, source)
			instructionPointer += 2
		case 5: //jump-if-true
			p1 := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			if p1 != 0 {
				instructionPointer = p2
			} else {
				instructionPointer += 3
			}
		case 6: //jump-if-false
			p1 := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			if p1 == 0 {
				instructionPointer = p2
			} else {
				instructionPointer += 3
			}
		case 7: //less then
			p1 := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			destination := intCode[instructionPointer+3]
			if p1 < p2 {
				intCode[destination] = 1
			} else {
				intCode[destination] = 0
			}
			instructionPointer += 4
		case 8: //equals
			p1 := getParam(intCode[instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(intCode[instructionPointer+2], opCode.Param2Mode)
			destination := intCode[instructionPointer+3]
			if p1 == p2 {
				intCode[destination] = 1
			} else {
				intCode[destination] = 0
			}
			instructionPointer += 4
		case 99:
			running = false
			break
		default:
			log.Fatalf("Unkown OpCode: %d", opCode)
		}
	}
	return outputs
}
