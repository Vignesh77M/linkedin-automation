package search

import "github.com/go-rod/rod"

func CollectProfiles(page *rod.Page) []string {
	links := page.MustElements(`a[href*="/in/"]`)
	results := []string{}

	for _, l := range links {
		href, _ := l.Attribute("href")
		if href != nil {
			results = append(results, *href)
		}
	}
	return results
}
