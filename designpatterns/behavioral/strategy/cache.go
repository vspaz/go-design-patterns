package strategy

type Cache struct {
	Storage        map[string]string
	EvictionPolicy evictionPolicy
	Capacity       int
	MaxCapacity    int
}
