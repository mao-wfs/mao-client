package ports

// ConfigPort is the configure for MAO.
type ConfigPort interface {
	FormatConfig(
		integTime IntegrationTime,
		winFunc WindowFunction,
		swOrder SwitchOrder,
	) (*WFSConfig, error)
}

// WFSConfig represents the configure for MAO.
type WFSConfig struct {
	Correlator *CorrelatorConfig
	Switch     *SwitchConfig
}

// CorrelatorConfig represents the configure for the correlator of MAO.
type CorrelatorConfig struct {
	IntegrationTime IntegrationTime
	WindowFunction  WindowFunction
}

// NewCorrelatorConfig returns the configure of the correlator for MAO wavefront sensor.
func NewCorrelatorConfig(integTime IntegrationTime, winFunc WindowFunction) *CorrelatorConfig {
	return &CorrelatorConfig{
		IntegrationTime: integTime,
		WindowFunction:  winFunc,
	}
}

// SwitchConfig represents the configure for the switch of MAO.
type SwitchConfig struct {
	SwitchOrder SwitchOrder
}

// NewSwitchConfig returns the configure of the switch for MAO wavefront sensor.
func NewSwitchConfig(swOrder SwitchOrder) *SwitchConfig {
	return &SwitchConfig{
		SwitchOrder: swOrder,
	}
}

// IntegrationTime is the integration time in a correlation.
type IntegrationTime uint8

// WindowFunction is the window function for a FFT in a correlation.
type WindowFunction string

// SwitchOrder is the order to switch.
type SwitchOrder []string
