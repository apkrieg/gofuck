package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	// Nil is 0x00
	Nil int = iota
	// Nop is No Op
	Nop
	// IncP increments the data pointer
	IncP
	// DecP decrements the data pointer
	DecP
	// IncD increments the byte at data pointer
	IncD
	// DecD decrements the byte at data pointer
	DecD
	// Put prints the byte at data pointer
	Put
	// Get stores input in the byte at data pointer
	Get
	// Begin of a loop
	Begin
	// End of a loop
	End
	// Dbg is a breakpoint that launches the debug console
	Dbg
)

type Runtime struct {
	DebugMode bool
	Pointer   int
	IP        int
	Code      []int
	// 32KiB Data section
	Data [1024 * 32]byte
}

func (r *Runtime) Run(code []int) error {
	r.Code = code
	for {
		if r.IP >= len(r.Code) {
			if r.DebugMode {
				r.Debug()
			} else {
				return nil
			}
		}
		switch i := r.Code[r.IP]; {
		case i == Nil:
			fmt.Println("Instruction Nil has no function")
			r.IP++
		case i == Nop:
			r.IP++
		case i == IncP:
			r.Pointer += r.Code[r.IP+1]
			r.IP += 2
		case i == DecP:
			r.Pointer -= r.Code[r.IP+1]
			r.IP += 2
		case i == IncD:
			r.Data[r.Pointer] += byte(r.Code[r.IP+1])
			r.IP += 2
		case i == DecD:
			r.Data[r.Pointer] -= byte(r.Code[r.IP+1])
			r.IP += 2
		case i == Put:
			fmt.Print(string(r.Data[r.Pointer]))
			r.IP++
		case i == Get:
			fmt.Scanf("%s", &r.Data[r.Pointer])
			r.IP++
		case i == Begin:
			if r.Data[r.Pointer] == 0 {
				r.IP = r.Code[r.IP+1]
				break
			}
			r.IP += 2
		case i == End:
			if r.Data[r.Pointer] != 0 {
				r.IP = r.Code[r.IP+1]
				break
			}
			r.IP += 2
		case i == Dbg:
			r.Debug()
			// We don't increment the IP because the user might want to set the IP to a specific value.
		default:
			return fmt.Errorf("instruction %x is not recognized", r.Code[r.IP])
		}
	}
}

func (r *Runtime) Debug() {
	fmt.Print("\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("dbg: ")
	for scanner.Scan() {
		switch command := strings.Split(scanner.Text(), " "); {
		case command[0] == "resume":
			r.IP++
			return
		case command[0] == "exit":
			os.Exit(0)
		case command[0] == "dump":
			// TODO: implement dumping the runtime data
		case command[0] == "set_cell":
			if len(command) != 3 {
				fmt.Printf("error: set_cell requires exactly two arguments, %d arguments provided\n", len(command)-1)
				break
			}
			var cell int
			var value int
			var err error
			if cell, err = strconv.Atoi(command[1]); err != nil {
				fmt.Printf("error: '%s' is not a number: %s\n", command[1], err)
				break
			}
			if value, err = strconv.Atoi(command[2]); err != nil {
				fmt.Printf("error: '%s' is not a number: %s\n", command[2], err)
				break
			}
			r.Data[cell] = byte(value)
		case command[0] == "get_cell":
			if len(command) != 2 {
				fmt.Printf("error: get_cell requires exactly one argument, %d arguments provided\n", len(command)-1)
				break
			}
			var cell int
			var err error
			if cell, err = strconv.Atoi(command[1]); err != nil {
				fmt.Printf("error: '%s' is not a number: %s\n", command[1], err)
				break
			}
			fmt.Printf("Cell: %d\n\tstr: %s\n\thex: %#x\n\tdec: %d\n", cell, string(r.Data[cell]), r.Data[cell], r.Data[cell])
		case command[0] == "set_pointer":
			if len(command) != 2 {
				fmt.Printf("error: set_pointer requires exactly one argument, %d arguments provided\n", len(command)-1)
				break
			}
			var value int
			var err error
			if value, err = strconv.Atoi(command[1]); err != nil {
				fmt.Printf("error: '%s' is not a number: %s\n", command[1], err)
				break
			}
			r.Pointer = value
		case command[0] == "get_pointer":
			if len(command) != 1 {
				fmt.Printf("error: get_pointer requires exactly zero arguments, %d arguments provided\n", len(command)-1)
				break
			}
			fmt.Printf("Pointer:\n\tstr: %s\n\thex: %#x\n\tdec: %d\n", string(r.Pointer), r.Pointer, r.Pointer)
		default:
			fmt.Printf("error: command '%s' not found\n", command[0])
		}
		fmt.Print("dbg: ")
	}
}
