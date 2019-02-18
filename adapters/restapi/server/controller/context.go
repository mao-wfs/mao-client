package controller

// Context represents the context of the current HTTP request.
type Context interface {
	Bind(i interface{}) error
	String(code int, s string) error
	JSON(code int, i interface{}) error
}
