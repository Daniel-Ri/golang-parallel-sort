package utils

import (
	"math/rand"
	"time"
)

func Random(n int) []int {
	s := make([]int, n)

	src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(src)

	for i := 0; i < n; i++ {
		s[i] = rand.Intn(n)
	}

	return s
}
