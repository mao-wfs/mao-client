package ports

// WFSConfig represents the configure for MAO wavefront sensor.
type WFSConfig struct {
	IntegrationTime IntegrationTime
	WindowFunction  WindowFunction
	SwitchOrder     SwitchOrder
}

// IntegrationTime is the integration time in a correlation.
type IntegrationTime uint8

// WindowFunction is the window function for a FFT in a correlation.
type WindowFunction string

// SwitchOrder is the order to switch.
type SwitchOrder []string
