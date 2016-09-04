// Package urand provides functions for obtaining (sets of) random integer values, dice throws
//
package urand

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

// Read n bytes from /dev/urandom
func RandomBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Get num random uint64 values
func GetUInt64(num uint64, print bool) []uint64 {
	var randValSize = 8
	randVals := make([]uint64, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random uint32 values
func GetUInt32(num uint64, print bool) []uint32 {
	var randValSize = 4
	randVals := make([]uint32, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random uint16 values
func GetUInt16(num uint64, print bool) []uint16 {
	var randValSize = 2
	randVals := make([]uint16, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random uint8 values
func GetUInt8(num uint64, print bool) []uint8 {
	var randValSize = 1
	randVals := make([]uint8, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random int64 values
func GetInt64(num uint64, print bool) []int64 {
	var randValSize = 8
	randVals := make([]int64, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random int32 values
func GetInt32(num uint64, print bool) []int32 {
	var randValSize = 4
	randVals := make([]int32, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random int16 values
func GetInt16(num uint64, print bool) []int16 {
	var randValSize = 2
	randVals := make([]int16, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Get num random int8 values
func GetInt8(num uint64, print bool) []int8 {
	var randValSize = 1
	randVals := make([]int8, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil
		}
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals
}

// Roll a dice with faces (default 6) howMany (default 1) time
func DiceRolls(faces uint64, howMany uint64, print bool) []uint64 {
	if faces == 0 {
		faces = 6
	}
	if howMany == 0 {
		howMany = 1
	}
	diceRolls := make([]uint64, howMany)
	var i uint64 = 0
	rawValues := GetUInt64(howMany, false)
	for ; i < howMany; i++ {
		rawValue := rawValues[i]
		dieCast := rawValue - ((rawValue / faces) * faces) + 1
		diceRolls[i] = dieCast
	}
	if print {
		for i = 0; i < howMany; i++ {
			fmt.Println(diceRolls[i])
		}
	}
	return diceRolls
}
