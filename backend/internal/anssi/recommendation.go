package anssi

import "encoding/json"

type Recommendation int

const (
	MinLength = iota
	UppercaseLetters
	LowercaseLetters
	NumericCharacters
	SpecialCharacters
)

// Implement the Stringer interface https://pkg.go.dev/fmt@go1.24.2#Stringer
func (r Recommendation) String() string {
	switch r {
	case MinLength:
		return "Minimum length of 12 characters"
	case UppercaseLetters:
		return "At least one uppercase letter"
	case LowercaseLetters:
		return "At least one lowercase letter"
	case NumericCharacters:
		return "At least one numeric character"
	case SpecialCharacters:
		return "At least one special character"
	default:
		return "Unknown recommendation"
	}
}

// Implement the Marshaller interface https://pkg.go.dev/encoding/json#Marshaler
func (r Recommendation) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
