package gettime

import (
	"time"
)

var StartTime time.Time

func InitializeTimer() {
	StartTime = time.Now()
}

func GetTimer() any {
	return time.Since(StartTime)
}
