package stack

type node struct {
	next  *node
	value interface{}
}

type LinkedStack struct {
	head *node
}

func NewLinkedStack() Stack {
	return &LinkedStack{
		head: nil,
	}
}

func (l *LinkedStack) Push(value interface{}) int {
	l.head = &node{next: l.head, value: value}
	return 0
}

func (l *LinkedStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}

	value := l.head.value
	l.head = l.head.next
	return value
}

func (l *LinkedStack) Top() interface{} {
	if l.IsEmpty() {
		return nil
	}
	return l.head.value
}

func (l *LinkedStack) Flush() {
	l.head = nil
}

func (l *LinkedStack) IsEmpty() bool {
	return l.head == nil
}
