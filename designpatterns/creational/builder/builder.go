package builder

type IBuilder interface {
	WithChasis(chasis string) IBuilder
	WithGasEngine(isGas bool) IBuilder
	WithElectricalMotor(isElectrical bool) IBuilder
	WithMake(make string) IBuilder
	WithSeats(seat int) IBuilder
	Build() *Car
}

func NewCar() CarBuilder {
	return CarBuilder{
		chasis:          "undef",
		gasEngine:       false,
		electricalMotor: false,
	}
}

type CarBuilder struct {
	make            string
	chasis          string
	gasEngine       bool
	electricalMotor bool
}

func (c *CarBuilder) WithSeats(seat int) IBuilder {
	panic("implement me")
}

type Meta struct {
	Year  int
	Month string
}

type Specs struct {
	Wheels          int
	Seats           int
	Chasis          string
	GasEngine       bool
	ElectricalMotor bool
}

type Car struct {
	Meta     Meta
	Specs    Specs
	color    string
	maxSpeed string
	price    float64
	Make     string
}

func (c *CarBuilder) WithChasis(chasis string) IBuilder {
	c.chasis = chasis
	return c
}

func (c *CarBuilder) WithGasEngine(isGas bool) IBuilder {
	c.gasEngine = isGas
	return c
}

func (c *CarBuilder) WithElectricalMotor(isElectrical bool) IBuilder {
	c.electricalMotor = isElectrical
	return c
}

func (c *CarBuilder) WithMake(make string) IBuilder {
	c.make = make
	return c
}

func (c *CarBuilder) Build() *Car {
	return &Car{
		Meta: Meta{
			Year:  2022,
			Month: "January",
		},
		Specs: Specs{
			Chasis:          c.chasis,
			ElectricalMotor: c.electricalMotor,
			GasEngine:       c.gasEngine,
		},
		Make: c.make,
	}
}
