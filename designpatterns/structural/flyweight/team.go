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

type HistoricalData struct {
	Year         uint64
	LegueResults []Game
}

type Team struct {
	Id      uint64
	Name    string
	Shield  []byte
	Players []Player
}
