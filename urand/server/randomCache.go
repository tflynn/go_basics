// Cache holds random number sets for various durations
package server

import (
	"github.com/robfig/go-cache"
	"time"
	"fmt"
	"github.com/tflynn/gobasics/urand/core"
)

var (
	//randomSetCache contains sets of random numbers  with specific expirations. Cleanup every 60 seconds
	randomSetCache = cache.New(-1, 60*time.Second)
)
type RandomEntry struct {
	indexUint64 uint64 // uint64 index entry
	valuesUint64 []uint64 // Array of uint64 values
	indexUint32 uint32 // uint32 index entry
	valuesUint32 []uint32 // Array of uint32 values
	indexUint16 uint16 // uint16 index entry
	valuesUint16 []uint16 // Array of uint16 values
	indexUint8 uint8 // uint8 index entry
	valuesUint8 []uint8 // Array of uint8 values
	indexUint uint // uint index entry
	valuesUint []uint // Array of uint values
	indexInt64 int64 // int64 index entry
	valuesInt64 []int64 // Array of int64 values
	indexInt32 int32 // int32 index entry
	valuesInt32 []int32 // Array of int32 values
	indexInt16 int16 // int16 index entry
	valuesInt16 []int16 // Array of int16 values
	indexInt8 int8 // int8 index entry
	valuesInt8 []int8 // Array of int8 values
	indexInt int // int index entry
	valuesInt []int // Array of int values
}

type RandomSet struct {
	entries []*RandomEntry // Array of pointers to entries
}

type unexportedInterface interface {
	Set(*RandomSet, time.Duration)
	Get(string, string) (*RandomSet, bool)
}

func secToDuration(sec uint64) time.Duration {
	//fmt.Println("secToDuration sec ", sec)
	//fmt.Println("secToDuration time.Duration(sec) * time.Second ", time.Duration(sec) * time.Second)
	return time.Duration(sec) * time.Second
}

func durationToSec(d time.Duration) int64 {
	return int64(d.Nanoseconds() / time.Second.Nanoseconds())
}

func entryName(expiration uint64, randomType string) string {
	return fmt.Sprintf("%s,%d.seconds",randomType, expiration)
}

// Add an item to the cache, replacing any existing item.
// The key is assumed to be some numbers of seconds i.e. 3 * time.Second
func Set(randomSet *RandomSet, durationInSecs uint64, randomType string) {
	randomSetCache.Set(entryName(durationInSecs, randomType),randomSet,secToDuration(durationInSecs))
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// The key is assumed to be some numbers of seconds i.e. 3 * time.Second
// The entry will be created if not present
func Get(durationInSecs uint64, randomType string, setSize uint64, totalSets uint64 ) (*RandomSet, bool) {
	eName := entryName(durationInSecs, randomType)
	_, present := randomSetCache.Get(eName)
	if ! present {
		randomSet := generateRandomSet(durationInSecs, randomType, setSize, totalSets)
		randomSetCache.Set(eName,randomSet,secToDuration(durationInSecs))
	}
	entry, present := randomSetCache.Get(eName)
	switch  v := entry.(type) {
	case (*RandomSet):
		return v, true
	default:
		fmt.Println("Unexpected value type found in cache. Expected *RandomSet , got ", v)
		return nil, false
	}
}


func generateRandomEntry(randomType string, setSize uint64) *RandomEntry {
	randomEntry := new(RandomEntry)
	switch randomType {
	case "uint64":
		setSize := setSize + 1 // Extra is index
		randomValuesUint64 := core.GetUInt64(setSize,false)
		randomEntry.indexUint64 = randomValuesUint64[0]
		randomEntry.valuesUint64 = randomValuesUint64[1:]
	case "uint32":
		setSize := setSize + 1 // Extra is index
		randomValuesUint32 := core.GetUInt32(setSize,false)
		randomEntry.indexUint32 = randomValuesUint32[0]
		randomEntry.valuesUint32 = randomValuesUint32[1:]
	case "uint16":
		setSize := setSize + 1 // Extra is index
		randomValuesUint16 := core.GetUInt16(setSize,false)
		randomEntry.indexUint16 = randomValuesUint16[0]
		randomEntry.valuesUint16 = randomValuesUint16[1:]
	case "uint8":
		setSize := setSize + 1 // Extra is index
		randomValuesUint8 := core.GetUInt8(setSize,false)
		randomEntry.indexUint8 = randomValuesUint8[0]
		randomEntry.valuesUint8 = randomValuesUint8[1:]
	case "int64":
		setSize := setSize + 1 // Extra is index
		randomValuesInt64 := core.GetInt64(setSize,false)
		randomEntry.indexInt64 = randomValuesInt64[0]
		randomEntry.valuesInt64 = randomValuesInt64[1:]
	case "int32":
		setSize := setSize + 1 // Extra is index
		randomValuesInt32 := core.GetInt32(setSize,false)
		randomEntry.indexInt32 = randomValuesInt32[0]
		randomEntry.valuesInt32 = randomValuesInt32[1:]
	case "int16":
		setSize := setSize + 1 // Extra is index
		randomValuesInt16 := core.GetInt16(setSize,false)
		randomEntry.indexInt16 = randomValuesInt16[0]
		randomEntry.valuesInt16 = randomValuesInt16[1:]
	case "int8":
		setSize := setSize + 1 // Extra is index
		randomValuesInt8 := core.GetInt8(setSize,false)
		randomEntry.indexInt8 = randomValuesInt8[0]
		randomEntry.valuesInt8 = randomValuesInt8[1:]
	default:
		fmt.Println("generateRandomEntry unknown random type found %s",randomType)
	}
	return randomEntry
}

func generateRandomSet(durationInSecs uint64, randomType string, setSize uint64, totalSets uint64) *RandomSet {
	var totalEntries uint64
	if totalSets == 0 {
		totalEntries = calcTotalEntries(durationInSecs)
	} else {
		totalEntries = totalSets
	}
	randomEntries := make([]*RandomEntry,totalEntries, totalEntries)
	for i := uint64(0) ; i < totalEntries ; i++ {
		randomEntries[i] = generateRandomEntry(randomType, setSize)
	}
	randomSet := new(RandomSet)
	randomSet.entries = randomEntries
	return randomSet
}

// Determine how many entries in a set
func calcTotalEntries(secs uint64) uint64 {
	//TODO Algorithm needed
	return uint64(1 * secs) // 1 set per  second
}
