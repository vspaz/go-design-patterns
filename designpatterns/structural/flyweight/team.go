package flyweight

import "time"

type Player struct {
	Name         string
	LastName     string
	PreviousTeam uint64
	Pic          []byte
}

type Game struct {
	Date          time.Time
	VisitorId     uint64
	LocalId       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint64
	VisitorShoots uint64
}
