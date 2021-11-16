package adapter

type SoapAdapter struct {
	Soap *Soap
}

func (a *SoapAdapter) Get() string {
	return a.Soap.QuerySoap()
}
