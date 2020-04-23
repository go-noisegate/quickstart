package quickstart

import "time"

func SlowAdd(a, b int64) int64 {
	time.Sleep(time.Second)
	return a + b
}

func SlowSub(a, b int64) int64 {
	time.Sleep(time.Second)
	return a + b
}
