package jwt

// Error represents a JWT error.
type Error string

func (e Error) Error() string {
	return string(e)
}

var _ error = (Error)("")

const (
	// ErrPartMissed indicates that token format is invalid.
	ErrPartMissed = Error("token format is invalid")

	// ErrInvalidKey when provided an incorrect key.
	ErrInvalidKey = Error("key is invalid")

	// ErrInvalidSignature signature wasn't correct.
	ErrInvalidSignature = Error("signature is invalid")
)
