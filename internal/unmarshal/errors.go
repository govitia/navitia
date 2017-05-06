package unmarshal

import "fmt"

// Error is returned when unmarshalling fails
// It implements both error and github.com/pkg/errors's causer
type Error struct {
	// Type on which the unmarshaller where the error occurred works
	Type string

	// JSON Key where failure occurred
	Key string

	// Name of the key in package
	Name string

	// Value associated with the key
	Value interface{}

	// Message of the error
	Message string

	// Underlying error
	Underlying error

	// Full JSON data being unmarshalled
	JSON *[]byte
}

// Cause implements github.com/pkg/error's causer
func (err Error) Cause() error {
	return err.Underlying
}

// Error implements error
func (err Error) Error() string {
	msg := fmt.Sprintf("(*%s).UnmarshalJSON: Unmarshalling %s (json: \"%s\") with value \"%v\" failed", err.Type, err.Name, err.Key, err.Value)
	if err.Message != "" {
		msg += ": " + err.Message
	}
	if err.Underlying != nil {
		msg += " [" + err.Cause().Error() + "]"
	}
	return msg
}

// Generator allows us to make better error messages
type Generator struct {
	Type string
	JSON *[]byte
}

// NewGenerator creates a new Generator
func NewGenerator(curType string, json *[]byte) Generator {
	return Generator{
		Type: curType,
		JSON: json,
	}
}

// Gen creates a new Error
func (gen Generator) Gen(underlyingErr error, name string, key string, value interface{}, message string) error {
	return Error{
		Type:       gen.Type,
		Key:        key,
		Name:       name,
		Value:      value,
		Message:    message,
		Underlying: underlyingErr,
		JSON:       gen.JSON,
	}
}

// Close closes a Generator
func (gen Generator) Close() {
	gen.JSON = nil
}
