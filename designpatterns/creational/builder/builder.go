package builder

type IBuilder interface {
	WithWheels() IBuilder
	WithSeats() IBuilder
	WithChasis() IBuilder
	WithGasEngine() IBuilder
	WithElectricalMotor() IBuilder
}

type CarBuilder struct {
	wheels          int
	seats           int
	chasis          string
	gasEngine       bool
	electricalMotor bool
}

func (c *CarBuilder) WithSeats() IBuilder {
	panic("implement me")
}

func (c *CarBuilder) WithChasis() IBuilder {
	panic("implement me")
}

func (c *CarBuilder) WithGasEngine() IBuilder {
	panic("implement me")
}

func (c *CarBuilder) WithElectricalMotor() IBuilder {
	panic("implement me")
}

func (c *CarBuilder) WithWheels() IBuilder {
	c.wheels = 4
	return c
}
