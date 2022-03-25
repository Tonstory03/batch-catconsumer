package utils

import "time"

const FORMAT_DATE_UTC = "2006-01-02T15:04:05Z"

func Time2StrFormatUTC(d time.Time) string {
	return d.Format(FORMAT_DATE_UTC)
}
