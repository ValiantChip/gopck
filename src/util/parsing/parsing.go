package parsing

type Parsable interface {
	String() string
}

type Unwrappable interface {
	Unwrap() any
}

type UnsupportedTypeError struct {
	Err string
}

func (e UnsupportedTypeError) Error() string {
	return e.Err
}
