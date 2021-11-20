package bridge

import (
	"errors"
	"fmt"
	"io"
)

type DataStreamWriter struct {
	Writer io.Writer
}

func (d *DataStreamWriter) PrintText(text string) error {
	if d.Writer == nil {
		return errors.New("io.Writer not set")
	}
	fmt.Fprintf(d.Writer, "%s", text)
	return nil
}
