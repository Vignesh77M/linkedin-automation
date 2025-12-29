package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func NewBrowser() *rod.Browser {
	l := launcher.New().
		Headless(false).
		Set("disable-blink-features", "AutomationControlled")

	return rod.New().
		ControlURL(l.MustLaunch()).
		MustConnect()
}
