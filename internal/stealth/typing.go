package stealth

import (
	"math/rand"
	"github.com/go-rod/rod"
	"time"
)

func TypeLikeHuman(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(rand.Intn(120)+80) * time.Millisecond)
	}
}
