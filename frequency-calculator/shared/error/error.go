package errors

type ConstError string

func (e ConstError) Error() string {
	return string(e)
}

const (
	ErrInputNotEmpty = ConstError("Input Must not be empty")
)
