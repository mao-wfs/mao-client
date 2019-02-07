package domain

// WFSHandler represents the domain handler for MAO wavefront sensor.
type WFSHandler interface {
	Initialize(args ...interface{}) error
	Finalize() error
	Start() error
	Halt() error
}
