package stack

type Stack interface {
	Push(value interface{}) int
	Pop() interface{}
	Top() interface{}
	Flush()
	IsEmpty() bool
}
