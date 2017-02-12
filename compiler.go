package main

import (
	"fmt"
)

func Compile(code string) ([]int, error) {
	var program []int
	var jmpStack []int
	for i := 0; i < len(code); i++ {
		switch {
		// IncP
		case string(code[i]) == ">":
			temp := 1
			for {
				if len(code) <= i+1 {
					break
				}
				if string(code[i+1]) != ">" {
					break
				}
				temp++
				i++
			}
			program = append(program, []int{IncP, temp}...)
		// DecP
		case string(code[i]) == "<":
			temp := 1
			for {
				if len(code) <= i+1 {
					break
				}
				if string(code[i+1]) != "<" {
					break
				}
				temp++
				i++
			}
			program = append(program, []int{DecP, temp}...)
		// IncD
		case string(code[i]) == "+":
			temp := 1
			for {
				if len(code) <= i+1 {
					break
				}
				if string(code[i+1]) != "+" {
					break
				}
				temp++
				i++
			}
			program = append(program, []int{IncD, temp}...)
		// DecD
		case string(code[i]) == "-":
			temp := 1
			for {
				if len(code) <= i+1 {
					break
				}
				if string(code[i+1]) != "-" {
					break
				}
				temp++
				i++
			}
			program = append(program, []int{DecD, temp}...)
		// Put
		case string(code[i]) == ".":
			program = append(program, Put)
		// Get
		case string(code[i]) == ",":
			program = append(program, Get)
		// Start
		case string(code[i]) == "[":
			program = append(program, []int{Begin, 0}...)
			jmpStack = append(jmpStack, len(program))
		// End
		case string(code[i]) == "]":
			program = append(program, []int{End, jmpStack[len(jmpStack)-1]}...)
			// Set the opening ['s jump location
			program[jmpStack[len(jmpStack)-1]-1] = len(program)
			jmpStack = jmpStack[:len(jmpStack)-1]
		// Dbg
		case string(code[i]) == "d":
			program = append(program, Dbg)
		default:
			return nil, fmt.Errorf("character '%s' is not a valid instruction", string(code[i]))
		}
	}
	return program, nil
}
