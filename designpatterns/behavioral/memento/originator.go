package memento

type Originator struct {
	state string
}

type Memento struct {
	savedState string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

func (o *Originator) GetMemento() Memento {
	return Memento{o.state}
}

func (m *Memento) GetState() string {
	return m.savedState
}

func (o *Originator) Restore(memento Memento) {
	o.state = memento.GetState()
}
