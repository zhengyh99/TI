package utils

import "time"

func TimeFormat(t time.Time) string {
	return t.Format("2006年1月2日 15:04:05")
}
