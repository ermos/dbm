package utils

import (
	"fmt"
	"time"
)

func FormatXTimeAgo(t time.Time, neverText string) string {
	if t == (time.Time{}) {
		return neverText
	}

	duration := time.Since(t)

	years := int(duration.Hours() / 24 / 365)
	months := int(duration.Hours()/24/30) % 12
	days := int(duration.Hours()/24) % 30
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60

	if years >= 1 {
		return fmt.Sprintf("%d year%s ago", years, pluralize(years))
	} else if months >= 1 {
		return fmt.Sprintf("%d month%s ago", months, pluralize(months))
	} else if days >= 1 {
		return fmt.Sprintf("%d day%s ago", days, pluralize(days))
	} else if hours >= 1 {
		return fmt.Sprintf("%d hour%s ago", hours, pluralize(hours))
	} else if minutes >= 1 {
		return fmt.Sprintf("%d minute%s ago", minutes, pluralize(minutes))
	}

	return "less than a minute ago"
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	} else {
		return "s"
	}
}
