package adapter

func Get(client Rest) string {
	return client.Get()
}