package reducer

// Error represents an error condition caused by the reducer.
type Error struct {
	message string
	cause   error
}

func (e *Error) Error() string {
	if e.cause != nil {
		return e.message + ": " + e.cause.Error()
	}
	return e.message
}
