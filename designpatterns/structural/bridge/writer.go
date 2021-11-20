package bridge

type IWriter interface {
	Flush(text string) error
}
