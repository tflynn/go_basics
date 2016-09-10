// urandcli - Console interface for urand library
package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unsafe"
	"github.com/tflynn/gobasics/urand/core"
)

// urandcli Main entry point
func main() {
	process_command_line()
}

// urandcli Process the command line
func process_command_line() {
	const usage = `Generate random integers

Usage: urandcli [-h | intType | diceRolls | diceWords ] [howMany] [dieFaces]

Where:

    -h Print this information and exit
    intType Integer type - int, int8, int16, int32, int64, uint8, uint16, uint32, uint64. Defaults to 'uint' if not specified.
    diceRolls Roll the dice.
    diceWords Words based on dice rolls.
    howMany
	  Number of integers to generate. Defaults to 1.
	  Number of times to roll the dice. Defaults to 1.
	  Number of diceWords. Defaults to 6.
    dieFaces Number of die faces. Defaults to 6
`
	var cmd string
	var howMany uint64 = 1
	var dieFaces uint64 = 6
	var validIntCmd = regexp.MustCompile(`^[u]*int`)
	var validDiceCmd = regexp.MustCompile(`^diceRolls$`)
	var validDiceWordsCmd = regexp.MustCompile(`^diceWords$`)

	args := os.Args[1:]
	if len(args) > 0 {
		cmd = args[0]

		if validDiceWordsCmd.MatchString(cmd) {
			howMany = 6
		}

		if len(args) > 1 && (validIntCmd.MatchString(cmd) || validDiceCmd.MatchString(cmd) || validDiceWordsCmd.MatchString(cmd)) {
			total, err := strconv.ParseInt(args[1], 10, 64) // use base 10 for sanity
			if err != nil {
				fmt.Println(err)
			}
			howMany = uint64(total)
		}
		if len(args) > 2 && validDiceCmd.MatchString(cmd) {
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
		core.GetUInt64(howMany, true)
	case "uint32":
		core.GetUInt32(howMany, true)
	case "uint16":
		core.GetUInt16(howMany, true)
	case "uint8":
		core.GetUInt8(howMany, true)
	case "int64":
		core.GetInt64(howMany, true)
	case "int32":
		core.GetInt32(howMany, true)
	case "int16":
		core.GetInt16(howMany, true)
	case "int8":
		core.GetInt8(howMany, true)
	case "diceRolls":
		core.DiceRolls(dieFaces, howMany, true)
	case "diceWords":
		core.GetRandDiceWords(howMany, true)
	default:
		core.GetUInt64(howMany, true)
	}
}
