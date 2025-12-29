package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// Simple human-like mouse movement (POC-safe)
func MoveMouse(page *rod.Page, x, y int) {
	jitterX := float64(x + rand.Intn(5) - 2)
	jitterY := float64(y + rand.Intn(5) - 2)

	page.Mouse.MoveTo(proto.Point{
		X: jitterX,
		Y: jitterY,
	})
}
