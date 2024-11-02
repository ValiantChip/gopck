package parsing

type Parsable interface {
	String() string
}

type UnsupportedTypeError struct {
	Err string
}

func (e UnsupportedTypeError) Error() string {
	return e.Err
}
