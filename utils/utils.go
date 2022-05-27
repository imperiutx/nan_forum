package utils

import (
	"fmt"
	"time"
)

func TimeSince(t time.Time) string {
	data := fmt.Sprintf("%v", time.Since(t))
	minute := 1 * time.Minute
	hour := 1 * time.Hour
	if time.Since(t) > 24*hour {
		h, _ := time.ParseDuration(data)
		return fmt.Sprintf("%.0f days ago ", h.Hours()/24)
	} else if time.Since(t) > hour {
		h, _ := time.ParseDuration(data)
		return fmt.Sprintf("%.0f hours ago", h.Hours())
	} else if time.Since(t) > minute {
		h, _ := time.ParseDuration(data)
		return fmt.Sprintf("%.0f minutes ago", h.Minutes())
	}
	h, _ := time.ParseDuration(data)

	return fmt.Sprintf("%.0f seconds ago", h.Seconds())
}
