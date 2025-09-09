package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func ExtractBrowserUA(ua string) string {
	browserRegex := regexp.MustCompile(`(Chrome|Firefox|Safari|Edge)/[0-9.]+`)
	browser := browserRegex.FindString(ua)
	if browser == "" {
		browser = "Unknown"
	}

	os := "Other"

	switch {
	case strings.Contains(ua, "Mac OS X"):
		re := regexp.MustCompile(`Mac OS X [0-9_.]+`)
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
	case strings.Contains(ua, "Linux"):
		os = "Linux"
	}

	return fmt.Sprintf("%s (%s)", browser, os)
}
