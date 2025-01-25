package custom_types

import "fmt"

type HeaderName string
type HeaderValue string
type Headers map[HeaderName]HeaderValue

type NoSuchHeader struct {
	Code    int
	Message string
}

func (e *NoSuchHeader) Error() string {
	return fmt.Sprintf("Error Code %d: %s", e.Code, e.Message)
}

func (e *NoSuchHeader) CodeInfo() int {
	return e.Code
}

func (e *NoSuchHeader) MessageInfo() string {
	return e.Message
}
