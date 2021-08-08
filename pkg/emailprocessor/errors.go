package emailprocessor

type ImapError struct {
	//Mail   string
	//Action string
	Err    error
}

func (e ImapError) Error() string {
	return e.Err.Error()
}

type ParseMailError struct {
	//Mail   string
	//Action string
	Err    error
}

func (e ParseMailError) Error() string {
	return e.Err.Error()
}

type ProcessMailError struct {
	//Mail   string
	//Action string
	Err    error
}

func (e ProcessMailError) Error() string {
	return e.Err.Error()
}

type FireflyError struct {
	//Mail   string
	//Action string
	Err    error
}

func (e FireflyError) Error() string {
	return e.Err.Error()
}
