package bridge

import "fmt"

type CustomWriter struct {
}

func (c *CustomWriter) Flush(text string) error {
	_, err := c.Write([]byte(text))
	return err
}

// interface to match io.Writer
func (c *CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	return len(p), nil
}
