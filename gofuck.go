package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	flag.Parse()
	r := Runtime{}
	switch tool := flag.Arg(0); {
	// Build command
	case tool == "build":
		fileName := flag.Arg(1)
		if fileName == "" {
			log.Println("gofuck: error: please provide a filename")
			os.Exit(1)
		}
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("gofuck: error: unable to open %s: %s\n", fileName, err)
			os.Exit(1)
		}
		defer file.Close()
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("gofuck: error: unable to read %s: %s\n", fileName, err)
			os.Exit(1)
		}
		bytecode, err := Compile(string(fileData))
		if err != nil {
			log.Println("gofuck: error: compilation error:", err)
			os.Exit(1)
		}
		outFile, err := os.Create("a.bfc")
		if err != nil {
			log.Println("gofuck: error: unable to create output file:", err)
			os.Exit(1)
		}
		buf := []byte{0xbf, 0xbf}
		for i := range bytecode {
			temp := make([]byte, 8)
			binary.LittleEndian.PutUint64(temp, uint64(bytecode[i]))
			buf = append(buf, temp...)
		}
		if _, err := outFile.Write(buf); err != nil {
			log.Println("gofuck: error: unable to write to out file:", err)
			os.Exit(1)
		}
		log.Println("gofuck: compilation completed successfully")
		return
	// Debug command
	case tool == "debug":
		r.DebugMode = true
		fallthrough
	// Run command
	case tool == "run":
		fileName := flag.Arg(1)
		if fileName == "" {
			log.Println("gofuck: error: please provide a filename")
			os.Exit(1)
		}
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("gofuck: error: unable to open %s: %s\n", fileName, err)
			os.Exit(1)
		}
		defer file.Close()
		fileData, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("gofuck: error: unable to read %s: %s\n", fileName, err)
			os.Exit(1)
		}
		if fileData[0] == 0xBF && fileData[1] == 0xBF {
			rawBytecode := fileData[2:]
			bytecode := make([]int, len(rawBytecode)/8)
			for i := range bytecode {
				bytecode[i] = int(binary.LittleEndian.Uint64(rawBytecode[i*8 : (i+1)*8]))
			}
			if err := r.Run(bytecode); err != nil {
				log.Println("gofuck: error: runtime error:", err)
				return
			}
		} else {
			bytecode, err := Compile(string(fileData))
			if err != nil {
				log.Println("gofuck: error: compilation error:", err)
				os.Exit(1)
			}
			if err := r.Run(bytecode); err != nil {
				log.Println("gofuck: error: runtime error:", err)
				return
			}
		}
	case tool == "":
		fmt.Print(">>>: ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "quit" {
				break
			}
			var err error
			var code []int
			if code, err = Compile(scanner.Text()); err != nil {
				fmt.Println("error: invalid brainfuck code:", err)
				fmt.Print(">>>: ")
				continue
			}
			r.IP = 0
			if err = r.Run(code); err != nil {
				fmt.Println("error: runtime error: unable to run compiled bytecode:", err)
				fmt.Print(">>>: ")
				continue
			}
			fmt.Print("\n>>>: ")
		}
	default:
		log.Println("gofuck: error: tool not found")
	}
}
