package bridge

import "errors"

type Writer struct {
	Text   string
	Printer IWriter
}

func (w *Writer) Print() error {
	return errors.New("not implemented yet")
}
