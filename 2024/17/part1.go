package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	adv = iota
	bxl
	bst
	jnz
	bxc
	out
	bdv
	cdv
)

var regA int
var regB int
var regC int
var program []int
var output []int

func main() {
	// (assume input is valid)
	scanner := bufio.NewScanner(os.Stdin)
	// parse initial register values
	scanner.Scan()
	regString := scanner.Text()[len("Register A: "):]
	regA, _ = strconv.Atoi(regString)
	scanner.Scan()
	regString = scanner.Text()[len("Register B: "):]
	regB, _ = strconv.Atoi(regString)
	scanner.Scan()
	regString = scanner.Text()[len("Register C: "):]
	regC, _ = strconv.Atoi(regString)
	// parse the empty line between registers and instructions
	scanner.Scan()
	// parse instructions
	scanner.Scan()
	instructionsString := scanner.Text()[len("Program: "):]
	instructions := strings.Split(instructionsString, ",")
	for _, instruction := range instructions {
		instructionInt, _ := strconv.Atoi(instruction)
		program = append(program, instructionInt)
	}

	// execute program
	instructionPointer := 0
	for instructionPointer < len(program)-1 {
		switch program[instructionPointer] {
		case adv:
			comboOperand := getComboOperand(program[instructionPointer+1])
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			regA = regA / denominator
		case bxl:
			literalOperand := program[instructionPointer+1]
			regB = regB ^ literalOperand
		case bst:
			comboOperand := getComboOperand(program[instructionPointer+1])
			regB = comboOperand % 8
		case jnz:
			if regA != 0 {
				literalOperand := program[instructionPointer+1]
				instructionPointer = literalOperand
				continue
			}
		case bxc:
			regB = regB ^ regC
		case out:
			comboOperand := getComboOperand(program[instructionPointer+1])
			output = append(output, comboOperand%8)
		case bdv:
			comboOperand := getComboOperand(program[instructionPointer+1])
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			regB = regA / denominator
		case cdv:
			comboOperand := getComboOperand(program[instructionPointer+1])
			denominator := int(math.Pow(float64(2), float64(comboOperand)))
			regC = regA / denominator
		default:
			panic("Unexpected opcode")
		}
		instructionPointer += 2
	}

	for i := 0; i < len(output)-1; i++ {
		fmt.Print(output[i], ",")
	}
	fmt.Println(output[len(output)-1])

}

func printState() {
	fmt.Println(regA, regB, regC, program)
}

func getComboOperand(instruction int) int {
	if instruction >= 0 && instruction <= 3 {
		return instruction
	}
	switch instruction {
	case 4:
		return regA
	case 5:
		return regB
	case 6:
		return regC
	default:
		panic("Unable to determine combo operand")
	}
}
