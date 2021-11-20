package bridge

type Writer struct {
	Text    string
	Printer IWriter
}

func (w *Writer) Flush() error {
	return w.Printer.Flush(w.Text)
}
