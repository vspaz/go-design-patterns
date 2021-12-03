package command

type Undo struct {
	receiver Receiver
}

func (u *Undo) Undo(receiver Receiver) {
	u.receiver = receiver
}

func (u *Undo) RunCommand() string {
	return u.receiver.RunUndo()
}
