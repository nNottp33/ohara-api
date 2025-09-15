package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func GetBrowser(ua string) (browser string) {
	browserRegex := regexp.MustCompile(`(Chrome|Firefox|Safari|Edge|Opera|OPR|Brave|Vivaldi|SamsungBrowser|DuckDuckGo)/[0-9.]+`)
	browser = browserRegex.FindString(ua)

	if browser == "" {
		switch {
		case strings.Contains(ua, "PostmanRuntime"):
			re := regexp.MustCompile(`PostmanRuntime/[0-9.]+`)
			match := re.FindString(ua)
			if match != "" {
				browser = match
			} else {
				browser = "Postman"
			}
		case strings.Contains(ua, "Apidog"):
			re := regexp.MustCompile(`Apidog/[0-9.]+`)
			match := re.FindString(ua)
			if match != "" {
				browser = match
			} else {
				browser = "Apidog"
			}
		default:
			browser = "Unknown"
		}
	} else {
		if strings.HasPrefix(browser, "OPR/") {
			browser = strings.Replace(browser, "OPR", "Opera", 1)
		}
	}

	return browser
}

func ExtractBrowserUA(ua string) string {
	browser := GetBrowser(ua)
	os := "Other"
	fmt.Printf("%s\n", ua)
	switch {
	case strings.Contains(ua, "Mac OS X"):
		re := regexp.MustCompile(`Mac OS X [0-9_]+`)
		match := re.FindString(ua)
		os = strings.ReplaceAll(match, "_", ".")
	case strings.Contains(ua, "Windows NT"):
		re := regexp.MustCompile(`Windows NT [0-9.]+`)
		match := re.FindString(ua)
		os = match
	case strings.Contains(ua, "Android"):
		re := regexp.MustCompile(`Android [0-9.]+`)
		match := re.FindString(ua)
		os = match
	case strings.Contains(ua, "iPhone OS"):
		re := regexp.MustCompile(`iPhone OS [0-9_]+`)
		match := re.FindString(ua)
		os = strings.ReplaceAll(match, "_", ".")
	case strings.Contains(ua, "iPad"):
		os = "iPadOS"
	case strings.Contains(ua, "CrOS"):
		os = "ChromeOS"
	case strings.Contains(ua, "Linux"):
		os = "Linux"
	}

	return fmt.Sprintf("%s (%s)", browser, os)
}
