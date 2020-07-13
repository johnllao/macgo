package randlist

import (
	"math/rand"
	"time"
)

func RandomizeStrings(list []string) []string {

	var l = make([]string, len(list))
	copy(l, list)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(l), func(i, j int) {
		l[i], l[j] = l[j], l[i]
	})

	return l
}
