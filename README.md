go_basics
=========

Basic Go (golang) tools

urandcli
--------

Console interface for urand library.

  Generate random integers
  
  Usage: 
  
    urandcli [-h | intType | diceRolls ] [howMany] [dieFaces]
  
  Where:
  
      -h Print this information and exit
      intType Integer type - int, int8, int16, int32, int64, uint8, uint16, uint32, uint64. Defaults to 'uint' if not specified
      diceRolls Roll the dice
      
      howMany Number of integers to generate or number of times to roll the dice, Defaults to 1.
      dieFaces Number of die faces. Defaults to 6
  
