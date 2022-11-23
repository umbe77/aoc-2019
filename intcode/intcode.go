package intcode

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type IntCodeStatus int

const (
	Running IntCodeStatus = iota
	Done
)

type IntCode struct {
	intCode            []int
	instructionPointer *int
	compiledInputs     map[int]int
	Status             *IntCodeStatus
	InitInputs         []int
	initInputsRead     *bool
	relativeBase       *int
}

func New(input string) IntCode {
	ip := 0
	s := Running
	iir := false
	rb := 0
	ic := IntCode{
		instructionPointer: &ip,
		compiledInputs:     make(map[int]int),
		Status:             &s,
		initInputsRead:     &iir,
		relativeBase:       &rb,
	}
	ic.intCode = compile(input)
	return ic
}

func (ic IntCode) Copy() IntCode {
	newCode := make([]int, len(ic.intCode))
	copy(newCode, ic.intCode)
	ip := 0
	s := Running
	iir := false
	rb := 0
	return IntCode{
		intCode:            newCode,
		instructionPointer: &ip,
		compiledInputs:     make(map[int]int),
		Status:             &s,
		InitInputs:         ic.InitInputs,
		initInputsRead:     &iir,
		relativeBase:       &rb,
	}
}

func (ic IntCode) Set(pointer, value int) {
	ic.intCode[pointer] = value
}
func (ic IntCode) Get(poiner int) int {
	return ic.intCode[poiner]
}
func (ic IntCode) Size() int {
	return len(ic.intCode)
}

func compile(input string) []int {

	strIntCode := strings.Split(input, ",")
	intCode := make([]int, len(strIntCode)+(1024))

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

func printIntCode(code []int) {
	fmt.Print("[ ")
	sc := make([]string, len(code))
	for i, b := range code {
		sc[i] = strconv.Itoa(b)
	}
	fmt.Print(strings.Join(sc, ","))
	fmt.Print(" ]")
	fmt.Println()
}

func (ic IntCode) Run(inputs []int) int {
	opCode := OpCode{}
	currentInputPointer := 0
	var getInputValue = func() int {
		if len(ic.InitInputs) > 0 && !*ic.initInputsRead {
			*ic.initInputsRead = true
			v := ic.InitInputs[0]
			return v
		}
		v := inputs[0]
		currentInputPointer++
		return v
	}
	var getParam = func(p, mode int) int {
		if mode == 0 {
			return ic.intCode[p]
		} else if mode == 2 {
			return ic.intCode[*ic.relativeBase+p]
		}
		return p
	}
	var getDestinationParam = func(p, mode int) int {
		destination := ic.intCode[p]
		if mode == 2 {
			destination = *ic.relativeBase + destination
		}
		return destination
	}
	retrunValue := 0
	for *ic.Status == Running {
		opCode = getOpCode(ic.intCode[*ic.instructionPointer])
		switch opCode.Code {
		case 1: //Add
			lo := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			ro := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			destination := getDestinationParam(*ic.instructionPointer+3, opCode.Param3Mode)
			ic.intCode[destination] = lo + ro
			*ic.instructionPointer += 4
			break
		case 2: //Multiply
			lo := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			ro := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			destination := getDestinationParam(*ic.instructionPointer+3, opCode.Param3Mode)
			ic.intCode[destination] = lo * ro
			*ic.instructionPointer += 4
			break
		case 3: //Inputs
			destination := getDestinationParam(*ic.instructionPointer+1, opCode.Param1Mode)
			ic.intCode[destination] = getInputValue()
			*ic.instructionPointer += 2
		case 4: //Outputs
			source := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			*ic.instructionPointer += 2
			// printIntCode(ic.intCode)
			return source
		case 5: //jump-if-true
			p1 := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			if p1 != 0 {
				*ic.instructionPointer = p2
			} else {
				*ic.instructionPointer += 3
			}
		case 6: //jump-if-false
			p1 := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			if p1 == 0 {
				*ic.instructionPointer = p2
			} else {
				*ic.instructionPointer += 3
			}
		case 7: //less then
			p1 := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			destination := getDestinationParam(*ic.instructionPointer+3, opCode.Param3Mode)
			if p1 < p2 {
				ic.intCode[destination] = 1
			} else {
				ic.intCode[destination] = 0
			}
			*ic.instructionPointer += 4
		case 8: //equals
			p1 := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			p2 := getParam(ic.intCode[*ic.instructionPointer+2], opCode.Param2Mode)
			destination := getDestinationParam(*ic.instructionPointer+3, opCode.Param3Mode)
			if p1 == p2 {
				ic.intCode[destination] = 1
			} else {
				ic.intCode[destination] = 0
			}
			*ic.instructionPointer += 4
		case 9: // modify relative base
			p1 := getParam(ic.intCode[*ic.instructionPointer+1], opCode.Param1Mode)
			*ic.relativeBase += p1
			*ic.instructionPointer += 2
		case 99:
			*ic.Status = Done
			break
		default:
			log.Println(ic.instructionPointer)
			log.Fatalf("Unkown OpCode: %d", opCode)
		}
	}
	return retrunValue
}
