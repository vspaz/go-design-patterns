package builder

type IBuilder interface {
	WithChasis(chasis string) IBuilder
	WithGasEngine(isGas bool) IBuilder
	WithElectricalMotor(isElectrical bool) IBuilder
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
	chasis          string
	gasEngine       bool
	electricalMotor bool
}

type Specs struct {
	Wheels          int
	Seats           int
	Chasis          string
	GasEngine       bool
	ElectricalMotor bool
}

type Car struct {
	Escuderia string
	Brand     string
	Make      string
	Specs     Specs
	color     string
	maxSpeed  string
	price     float64
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

func (c *CarBuilder) Build() *Car {
	return &Car{
		Escuderia: "Ferrari",
		Make:      "Portofino",
		Brand:     "Ferrari",
		Specs: Specs{
			Chasis:          c.chasis,
			ElectricalMotor: c.electricalMotor,
			GasEngine: c.gasEngine,
		},
		color: "red",
	}
}
