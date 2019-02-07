package ports

// Config is the configure for MAO.
type Config interface {
	Correlator(args ...interface{})
	Switch(args ...interface{})
}
