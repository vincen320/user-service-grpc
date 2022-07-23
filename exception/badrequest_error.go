package exception

type BadRequestError struct {
	ErrorString string
}

func NewBadRequestError(errString string) *BadRequestError {
	return &BadRequestError{
		ErrorString: errString,
	}
}

//implement error tidak menerima pointer
func (b BadRequestError) Error() string {
	return b.ErrorString
}
