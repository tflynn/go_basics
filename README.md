go_basics
=========

Basic Go (golang) tools

Installation
------------

    go get https://github.com/tflynn/gobasics.git
    
    go install
    
    go install github.com/tflynn/gobasics/urand
    
    go install github.com/tflynn/gobasics/urandcli
    

urand
-----

Library to generate random integers, random dice throws, words based on random dice throws.

Documentation: https://godoc.org/github.com/tflynn/gobasics/urand

'Words based on random dice throws' is a direct implementation of Diceware http://world.std.com/~reinhold/diceware.html

urandcli
--------

Console interface for urand library.

  Generate random integers, random dice throws, words based on random dice throws.
  
  Usage: 
  
    urandcli [-h | intType | diceRolls | diceWords] [howMany] [dieFaces]
  
  Where:
  
      -h Print this information and exit

      intType Integer type - int, int8, int16, int32, int64, uint8, uint16, uint32, uint64. Defaults to 'uint' if not specified
      diceRolls Roll the dice.
      diceWords Words based on dice rolls.
      
      howMany 
        Number of integers to generate. Defaults to 1.
        Number of times to roll the dice. Defaults to 1.
        Number of diceWords. Defaults to 6.
      
      dieFaces Number of die faces. Defaults to 6 if not specified.
  
