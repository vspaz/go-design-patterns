package facade

type SecurityCode struct {
	Code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		Code: code,
	}
}
