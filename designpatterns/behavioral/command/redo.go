package command

type Redo struct {
	receiver Receiver
}

func (r *Redo) Redo(receiver Receiver) {
	r.receiver = receiver
}

func (r *Redo) RunCommand() string {
	return r.receiver.RunRedo()
}
