package server

import (
	"time"
	"fmt"
)

var (
	ErrKeyExists = fmt.Errorf("item already exists")
	ErrCacheMiss = fmt.Errorf("item not found")
)

type item struct {
	Object     interface{}
	Expiration *time.Time
}
