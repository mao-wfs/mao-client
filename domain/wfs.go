package domain

// WFSHandler represents the domain handler for MAO wavefront sensor.
type WFSHandler interface {
	Initialize(args ...interface{}) error
	Finalize() error
	Start(t ...interface{}) error
	Halt() error
}
