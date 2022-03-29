package locallogging

import (
	"time"
)

func calculateElapsedTime(startTime *time.Time, endTime *time.Time) int64 {
	return (endTime.UnixNano() - startTime.UnixNano()) / int64(time.Millisecond)
}
