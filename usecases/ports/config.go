package ports

// ConfigPort is the configure for MAO.
type ConfigPort interface {
	Correlator(args ...interface{})
	Switch(args ...interface{})
}
