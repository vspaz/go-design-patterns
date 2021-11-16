package adapter

type Soap interface {
	Decode(body string) string
}

