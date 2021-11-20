package bridge

import "fmt"

type ConsoleWriter struct{}

func (c *ConsoleWriter) Flush(text string) error {
	fmt.Printf("consolePrinter: %s", text)
	return nil
}
