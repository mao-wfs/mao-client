package device

// defaultBufSize is the default size of buffer.
const defaultBufSize = 1024

// Client is the client for MAO-WFS devices.
type Client interface {
	Write(msg string) error
	Read(bufSize int) ([]byte, error)
	Query(msg string, bufSize int) ([]byte, error)
}
