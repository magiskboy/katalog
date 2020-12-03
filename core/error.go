package core

const (
	RESOURCE_NOT_FOUND = 1
)

type errorCatalog struct {
	error
	Message string
	Code    string
}

func New(code uint32, message string) error {
	return &errorCatalog{Code: code, Message: message}
}

// Error compatible error golang
func (self *errorCatalog) Error() string {
	return self.Message
}
