package utils

import (
	"math/rand"
	"time"
)

func HumanDelay(minMs, maxMs int) {
	time.Sleep(time.Duration(rand.Intn(maxMs-minMs)+minMs) * time.Millisecond)
}
