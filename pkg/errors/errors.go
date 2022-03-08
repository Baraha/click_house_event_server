package errors

type DBNotAvailable struct {
	s string
}

func NewDBNotAvailable(text string) error {
	return &DBNotAvailable{text}
}

func (e *DBNotAvailable) Error() string {
	return e.s
}
