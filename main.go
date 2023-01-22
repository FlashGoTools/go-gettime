package gettime

import (
	"time"
)

var StartTime int64

func InitializeTimer() {
	StartTime = time.Now().UTC().UnixMilli()
}

func GetTimer() any {
	return time.Now().UTC().UnixMilli() - StartTime
}
