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
	index uint64 // Index entry
	values []uint64 // Array of values
}

type RandomSet struct {
	entries []*RandomEntry // Array of pointers to entries
}

type unexportedInterface interface {
	Set(string, string, *RandomSet, time.Duration)
	Get(string, string) (*RandomSet, bool)
}

func secToDuration(sec int64) time.Duration {
	//fmt.Println("secToDuration sec ", sec)
	//fmt.Println("secToDuration time.Duration(sec) * time.Second ", time.Duration(sec) * time.Second)
	return time.Duration(sec) * time.Second
}

func durationToSec(d time.Duration) int64 {
	return int64(d.Nanoseconds() / time.Second.Nanoseconds())
}

func entryName(expiration int64) string {
	return fmt.Sprintf("%d.seconds",expiration)
}

// Add an item to the cache, replacing any existing item.
// The key is assumed to be some numbers of seconds i.e. 3 * time.Second
func Set(randomSet *RandomSet, durationInSecs int64) {
	randomSetCache.Set(entryName(durationInSecs),randomSet,secToDuration(durationInSecs))
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// The key is assumed to be some numbers of seconds i.e. 3 * time.Second
// The entry will be created if not present
func Get(durationInSecs int64) (*RandomSet, bool) {
	eName := entryName(durationInSecs)
	_, present := randomSetCache.Get(eName)
	if ! present {
		randomSet := generateRandomSet(durationInSecs)
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


func generateRandomEntry(durationInSecs int64) *RandomEntry {
	totalValues := uint64(calcTotalValues(durationInSecs) + 1) // Extra is index
	randomEntry := new(RandomEntry)
	randomValues := core.GetUInt64(totalValues,false)
	randomEntry.index = randomValues[0]
	randomEntry.values = randomValues[1:]
	return randomEntry
}

func generateRandomSet(durationInSecs int64) *RandomSet {
	totalEntries := calcTotalEntries(durationInSecs)
	randomEntries := make([]*RandomEntry,totalEntries, totalEntries)
	for i := int64(0) ; i < totalEntries ; i++ {
		randomEntries[i] = generateRandomEntry(durationInSecs)
	}
	randomSet := new(RandomSet)
	randomSet.entries = randomEntries
	return randomSet
}

// Determine how many values in an entry
func calcTotalValues(secs int64) int64 {
	//TODO Algorithm needed
	//return int64(1000) // 1000 random numbers per entry
	return int64(10) // 1000 random numbers per entry //DEBUG
}

// Determine how many entries in a set
func calcTotalEntries(secs int64) int64 {
	//TODO Algorithm needed
	return int64(1 * secs) // 1 set per  second
}
