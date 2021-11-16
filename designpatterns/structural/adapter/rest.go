package adapter

type Rest struct{}

func (r *Rest) Get() string {
	return "making REST request"
}
