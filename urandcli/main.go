package main

import (
	"fmt"
	"github.com/tflynn/gobasics/urand"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	process_command_line()
}

func process_command_line() {
	var usage = `Generate random integers

Usage: urandcli [-h | intType | diceRolls ] [howMany] [dieFaces]

Where:

    -h Print this information and exit
    intType Integer type - int, int8, int16, int32, int64, uint8, uint16, uint32, uint64. Defaults to 'uint' if not specified
    diceRolls Roll the dice
    howMany Number of integers to generate or number of times to roll the dice, Defaults to 1.
    dieFaces Number of die faces. Defaults to 6
`
	var cmd string
	var howMany uint64 = 1
	var dieFaces uint64 = 6

	args := os.Args[1:]
	if len(args) > 0 {
		cmd = args[0]
		if len(args) > 1 {
			total, err := strconv.ParseInt(args[1], 10, 64) // use base 10 for sanity
			if err != nil {
				fmt.Println(err)
			}
			howMany = uint64(total)
		}
		if len(args) > 2 {
			total, err := strconv.ParseInt(args[2], 10, 64) // use base 10 for sanity
			if err != nil {
				fmt.Println(err)
			}
			dieFaces = uint64(total)
		}
	} else {
		var ignore uint = 0
		var intSize = unsafe.Sizeof(ignore)
		switch intSize {
		case 8:
			cmd = "uint64"
		case 4:
			cmd = "uint32"
		case 2:
			cmd = "uint16"
		case 1:
			cmd = "uint8"
		default:
			fmt.Println(usage)
		}
	}

	switch cmd {
	case "-h":
		fmt.Println(usage)
	case "uint64":
		urand.GetUInt64(howMany, true)
	case "uint32":
		urand.GetUInt32(howMany, true)
	case "uint16":
		urand.GetUInt16(howMany, true)
	case "uint8":
		urand.GetUInt8(howMany, true)
	case "int64":
		urand.GetInt64(howMany, true)
	case "int32":
		urand.GetInt32(howMany, true)
	case "int16":
		urand.GetInt16(howMany, true)
	case "int8":
		urand.GetInt8(howMany, true)
	case "diceRolls":
		urand.DiceRolls(dieFaces, howMany, true)
	default:
		urand.GetUInt64(howMany, true)
	}
}