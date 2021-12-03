package command

type Receiver struct {
}

func (r *Receiver) RunUndo() string {
	return "'undo' run by receiver"
}

func (r *Receiver) RunRedo() string {
	return "'redo' run by receiver"
}
