go_basics
=========

Basic Go (golang) tools

urandcli
--------

Console interface for urand library.

  Generate random integers, random dice throws.
  
  Usage: 
  
    urandcli [-h | intType | diceRolls ] [howMany] [dieFaces]
  
  Where:
  
      -h Print this information and exit

      intType Integer type - int, int8, int16, int32, int64, uint8, uint16, uint32, uint64. Defaults to 'uint' if not specified
      diceRolls Roll the dice
      
      howMany Number of integers to generate or number of times to roll the dice. Defaults to 1 if not specified.
      
      dieFaces Number of die faces. Defaults to 6 if not specified.
  
Installation

  go get https://github.com/tflynn/gobasics.git
  
  go install
  
  go install github.com/tflynn/gobasics/urand
  
  go install github.com/tflynn/gobasics/urandcli
  
