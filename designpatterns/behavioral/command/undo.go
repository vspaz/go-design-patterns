package command

type Undo struct {
	Receiver Receiver
}

func (u *Undo) Undo(receiver Receiver) {
	u.Receiver = receiver
}

func (u *Undo) RunCommand() string {
	return u.Receiver.RunUndo()
}
