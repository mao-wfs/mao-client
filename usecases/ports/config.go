package ports

// ConfigPort is the configure for MAO.
type ConfigPort interface {
	FormatConfig(corrConf *CorrelatorConfig, swConf *SwitchConfig) (*WFSConfig, error)
}

// WFSConfig represents the configure for MAO.
type WFSConfig struct {
	Correlator CorrelatorConfig
	Switch     SwitchConfig
}

// CorrelatorConfig represents the configure for the correlator of MAO.
// NOTE: TODO
type CorrelatorConfig struct {
}

// SwitchConfig represents the configure for the switch of MAO.
// NOTE: TODO
type SwitchConfig struct {
}
