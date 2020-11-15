package main

import (
	"time"
)


type TTLMap struct {
	m map[string] interface{}
}

func (m *TTLMap) Set(key string, value interface{}, expiry time.Duration) {
	m.m[key] = value
	var t = time.AfterFunc(expiry, func() {
		
	})
}

func main() {

}
