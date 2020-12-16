package types

type Customer interface {
	Phone() string
	Email() string
}

type HasToken interface {
	Token() string
}
