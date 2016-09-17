// Package urand provides functions for obtaining (sets of) random integer values, random dice throws, random dice words
package core

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
func GetUInt64(num uint64, print bool) ([]uint64, []interface{}) {
	var randValSize = 8
	randVals := make([]uint64, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random uint32 values
func GetUInt32(num uint64, print bool) ([]uint32, []interface{}) {
	var randValSize = 4
	randVals := make([]uint32, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random uint16 values
func GetUInt16(num uint64, print bool) ([]uint16, []interface{}) {
	var randValSize = 2
	randVals := make([]uint16, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random uint8 values
func GetUInt8(num uint64, print bool) ([]uint8, []interface{}) {
	var randValSize = 1
	randVals := make([]uint8, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random int64 values
func GetInt64(num uint64, print bool) ([]int64, []interface{}) {
	var randValSize = 8
	randVals := make([]int64, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random int32 values
func GetInt32(num uint64, print bool) ([]int32, []interface{}) {
	var randValSize = 4
	randVals := make([]int32, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random int16 values
func GetInt16(num uint64, print bool) ([]int16, []interface{}) {
	var randValSize = 2
	randVals := make([]int16, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
}

// Get num random int8 values
func GetInt8(num uint64, print bool) ([]int8, []interface{}) {
	var randValSize = 1
	randVals := make([]int8, num)
	randUntypedVals := make([]interface{}, num)
	var i uint64 = 0
	for ; i < num; i++ {
		randomBytes, err := RandomBytes(randValSize)
		buffer := bytes.NewBuffer(randomBytes)
		binary.Read(buffer, binary.BigEndian, &randVals[i])
		if err != nil {
			return nil, nil
		}
		randUntypedVals[i] = randVals[i]
	}
	if print {
		for i = 0; i < num; i++ {
			fmt.Println(randVals[i])
		}
	}
	return randVals, randUntypedVals
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
	rawValues, _ := GetUInt64(howMany, false)
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
