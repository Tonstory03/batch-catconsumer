package utils

import (
	"time"
)

const FORMAT_DATE_UTC = "2006-01-02T15:04:05Z"
const FORMAT_DATE_YYYYMMDD = "2006-01-02"
const FORMAT_DATE_ISO8601TIMEZONE = "2006-01-02T15:04:05.000-07:00"

func Time2StrFormatUTC(d time.Time) string {
	return d.UTC().Format(FORMAT_DATE_UTC)
}

func Str2TimeFormatYYYYMMDD(d string) (time.Time, error) {
	return time.Parse(FORMAT_DATE_YYYYMMDD, d)
}

func Str2TimeFormatUTC(d string) (time.Time, error) {
	return time.Parse(FORMAT_DATE_UTC, d)
}

func Str2TimeFormatISO8601Timezone(d string) (time.Time, error) {
	return time.Parse(FORMAT_DATE_ISO8601TIMEZONE, d)
}

func Time2StrFormatISO8601Timezone(d time.Time) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return d.In(loc).Format(FORMAT_DATE_ISO8601TIMEZONE)
}

func Unix2Time(t int64) time.Time {
	return time.Unix(t, 0)
}

func Unix13Digit2Time(t int64) time.Time {
	return time.UnixMilli(t)
}
