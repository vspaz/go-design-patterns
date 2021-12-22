package memento

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMementoOk(t *testing.T) {
	originator := &Originator{
		state: "idle",
	}
	originator.SetState("on")
	assert.Equal(t, "on", originator.GetState())

	memento := originator.GetMemento()
	originator.SetState("off")
	assert.Equal(t, "off", originator.GetState())

}
