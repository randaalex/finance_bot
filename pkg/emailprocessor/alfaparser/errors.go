package alfaparser

type ParseMailError struct {
	Mail   string
	Line   string
	Action string
	Err    error
}

func (e ParseMailError) Error() string {
	return e.Err.Error()
}
