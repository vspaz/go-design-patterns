package adapter

type IClient interface {
	Get() string
}

func Get(client IClient) string {
	return client.Get()
}
