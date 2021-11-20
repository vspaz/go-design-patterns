package bridge

import "fmt"

type CustomWriter struct {
}

func (c *CustomWriter) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	return len(p), nil
}
