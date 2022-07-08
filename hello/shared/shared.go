package shared

import "time"

var SharedHandlerParam HandlerParam

type HandlerParam struct {
	SleepTime         int
	CpuBound          bool
	Target            int
	SleepTimeDuration time.Duration
}
