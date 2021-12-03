package command

type Invoker struct {
	command ICommand
}

func (i *Invoker) SetCommand(command ICommand) {
	i.command = command
}
