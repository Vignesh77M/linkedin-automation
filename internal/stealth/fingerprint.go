package stealth

import "github.com/go-rod/rod"

// Minimal & SAFE stealth (LinkedIn-compatible)
func ApplyFingerprint(page *rod.Page) {

	// Run before page load
	page.MustEvalOnNewDocument(`() => {

		// webdriver
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		// languages
		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});

		// plugins
		Object.defineProperty(navigator, 'plugins', {
			get: () => [1, 2, 3, 4, 5]
		});

		// chrome runtime
		window.chrome = {
			runtime: {}
		};

	}`)
}
