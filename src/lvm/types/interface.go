package types

type Customer interface {
	Phone() string
	Email() string
}

type HasTokener interface {
	Token() (string, error)
	GetId() string
}

type Jsoner interface {
	Json() ([]byte, error)
}
