package types

type Customer interface {
	Phone() string
	Email() string
}

type HasTokener interface {
	Token() string
}

type Jsoner interface {
	Json() ([]byte, error)
}
