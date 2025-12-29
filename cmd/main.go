package main

import (
	"log"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/auth"
)

func main() {
	b := browser.NewBrowser()
	page := b.MustPage()

	if err := auth.Login(page); err != nil {
		log.Fatal(err)
	}

	log.Println("Logged in successfully")
}
