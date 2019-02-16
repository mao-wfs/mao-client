package ports

import (
	"fmt"
)

// WFSConfig represents the configure for MAO wavefront sensor.
type WFSConfig struct {
	IntegrationTime IntegrationTime
	WindowFunction  WindowFunction
	SwitchOrder     SwitchOrder
}

// Validate validate the configure of MAO wavefront sensor.
// NOTE: TODO (SwitchOrder)
func (conf *WFSConfig) Validate() error {
	if err := conf.validateIntegrationTime(); err != nil {
		return err
	}
	if err := conf.validateWindowFunction(); err != nil {
		return err
	}
	return nil
}

// validateIntegrationTime validate the integration time of MAO wavefront sensor.
func (conf *WFSConfig) validateIntegrationTime() error {
	switch conf.IntegrationTime {
	case Int5ms, Int10ms:
		return nil
	default:
		return fmt.Errorf("Invalid integration time: %d", conf.IntegrationTime)
	}
}

// validateWindowFunction validate the window function of MAO wavefront sensor.
func (conf *WFSConfig) validateWindowFunction() error {
	switch conf.WindowFunction {
	case None, Hamming, Hanning, Blackman:
		return nil
	default:
		return fmt.Errorf("Invalid window function: %s", conf.WindowFunction)
	}
}

// IntegrationTime is the integration time in a correlation.
type IntegrationTime uint8

const (
	// Int5ms : 5 ms integration.
	Int5ms IntegrationTime = 5
	// Int10ms : 10 ms integration.
	Int10ms = 10
)

// WindowFunction is the window function for a FFT in a correlation.
type WindowFunction string

const (
	// None : FFT with no window function.
	None WindowFunction = "none"
	// Hamming : FFT with "Hamming" function.
	Hamming = "hamming"
	// Hanning : FFT with "Hanning" function.
	Hanning = "hanning"
	// Blackman : FFT with "Blackman" function.
	Blackman = "blackman"
)

// SwitchOrder is the order to switch.
type SwitchOrder []string
