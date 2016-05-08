package lordcommander

var (
	// ErrAlreadyUsed is an error that a Commander returns when an attempt
	// it made to reuse it.
	ErrAlreadyUsed = NewError("commander already used")

	// ErrNonZero is an error that a Commander returns when the command
	// finishes but the process return code is not 0.
	ErrNonZero = NewError("command exited with non-zero status")
)

// Error is a strongly typed error for the Commander module. Having these sort
// of errors can make our Bugsnag look more readable.
type Error struct {
	message string
}

// NewError is an Error constructor.
func NewError(message string) error {
	return &Error{message: message}
}

func (e *Error) Error() string {
	return e.message
}
