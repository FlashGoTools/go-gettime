package gettime

import (
	"time"
)

var StartTime int64

func InitializeTimer() {
	StartTime = time.Now().UnixMicro()
}

func GetTimer() any {
	return time.Now().UnixMicro() - StartTime
}
