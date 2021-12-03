package command

type Redo struct {
	Receiver Receiver
}

func (r *Redo) Redo(receiver Receiver) {
	r.Receiver = receiver
}

func (r *Redo) RunCommand() string {
	return r.Receiver.RunRedo()
}
