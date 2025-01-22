package utils

import (
	"math/rand"
	"time"
)

func SleepRandom() {
	rand.Seed(time.Now().UnixNano())
	delay := time.Duration(rand.Intn(300)+100) * time.Millisecond
	time.Sleep(delay)
}
