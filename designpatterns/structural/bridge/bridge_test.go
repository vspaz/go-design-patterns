package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsoleWriter_Flush(t *testing.T) {
	writer := Writer{
		Text:    "some text goes here",
		Printer: &ConsoleWriter{},
	}

	assert.Nil(t, writer.Flush())
}

func TestDataStreamWriter_Flush(t *testing.T) {
	writer := Writer{
		Text:    "some text goes here",
		Printer: &CustomWriter{},
	}

	assert.Nil(t, writer.Flush())
}