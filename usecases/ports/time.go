package ports

import (
	"time"
)

// WFSTime represents the sensing time of MAO wavefront sensor.
type WFSTime struct {
	BeginTime   time.Time
	SensingTime time.Duration
}
