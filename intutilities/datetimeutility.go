package intutilities

import (
	"strings"
	"time"
)

func GetCurrentISO8601() string {
	strNow := time.Now().Format(time.RFC3339)
	arrNow := strings.Split(strNow, "+")

	if len(arrNow) == 2 {
		strIso := arrNow[0] + ".000+" + arrNow[1]
		return strIso
	} else {
		return strNow
	}
}

func GetISO8601(t *time.Time) string {
	strTime := t.Format(time.RFC3339)
	arrTime := strings.Split(strTime, "+")

	if len(arrTime) == 2 {
		strIso := arrTime[0] + ".000+" + arrTime[1]
		return strIso
	} else {
		return strTime
	}
}
