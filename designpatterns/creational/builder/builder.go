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
	c.seats = 4
	return c
}

func (c *CarBuilder) WithChasis() IBuilder {
	c.chasis = "aluminium"
	return c
}

func (c *CarBuilder) WithGasEngine() IBuilder {
	c.gasEngine = false
	return c
}

func (c *CarBuilder) WithElectricalMotor() IBuilder {
	c.electricalMotor = true
	return c
}

func (c *CarBuilder) WithWheels() IBuilder {
	c.wheels = 4
	return c
}
