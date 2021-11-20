package bridge

type ConsolePrinter struct{}

func (c *ConsolePrinter) PrintText(text string) string {
	return "consolePrinter: text"
}
