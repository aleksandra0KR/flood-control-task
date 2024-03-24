package model

import "time"

type FloodControlParameters struct {
	Interval time.Duration
	Limit    int64
}
