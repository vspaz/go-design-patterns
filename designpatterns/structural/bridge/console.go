package bridge

import "fmt"

type ConsolePrinter struct{}

func (c *ConsolePrinter) PrintText(text string) error {
	fmt.Printf("consolePrinter: %s", text)
	return nil
}
