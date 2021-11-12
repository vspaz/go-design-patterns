package builder

type IBuilder interface {
	WithWheels() IBuilder
	WithSeats() IBuilder
	WithChasis() IBuilder
	WithGasEngine() IBuilder
	WithElectricMotor() IBuilder
}
