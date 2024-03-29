package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoker_RunCommand(t *testing.T) {
	receiver := Receiver{}
	undoCmd := &Undo{receiver: receiver}
	caller := Invoker{
		command: undoCmd,
	}
	assert.Equal(t, "'undo' run by receiver", caller.RunCommand())

	redoCmd := &Redo{receiver: receiver}
	caller.SetCommand(redoCmd)
	assert.Equal(t, "'redo' run by receiver", caller.RunCommand())
}
