package auth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-rod/rod"
	"linkedin-automation/internal/stealth"
	"linkedin-automation/internal/utils"
)

func isSecurityCheckpoint(page *rod.Page) bool {
	url := page.MustInfo().URL
	if strings.Contains(url, "checkpoint") {
		return true
	}

	hasCaptcha, _, _ := page.Has(`iframe[src*="captcha"]`)
	if hasCaptcha {
		return true
	}

	return false
}


func Login(page *rod.Page) error {

	stealth.ApplyFingerprint(page)
	page.MustNavigate("https://www.linkedin.com/login")

	user := os.Getenv("LI_USERNAME")
	pass := os.Getenv("LI_PASSWORD")

	if user == "" || pass == "" {
		return errors.New("missing credentials: set LI_USERNAME and LI_PASSWORD")
	}

	stealth.TypeLikeHuman(page.MustElement("#username"), user)
	utils.HumanDelay(300, 700)

	stealth.TypeLikeHuman(page.MustElement("#password"), pass)
	utils.HumanDelay(300, 700)

	page.MustElement("button[type=submit]").MustClick()
	utils.HumanDelay(3000, 5000)

	if isSecurityCheckpoint(page) {
		log.Println("[WARN] Security checkpoint detected.")
		log.Println("Please solve the CAPTCHA manually in the browser.")
		log.Println("Press ENTER here once done...")
		fmt.Scanln()
	}

	if !strings.Contains(page.MustInfo().URL, "/feed") {
		return errors.New("login failed, captcha, or checkpoint detected")
	}

	return nil
}
