package domain

import (
	"time"
)

// WFSHandler represents the domain handler for MAO wavefront sensor.
type WFSHandler interface {
	Initialize(args ...interface{}) error
	Finalize() error
	Start(beginTime time.Time) error
	Halt() error
}
