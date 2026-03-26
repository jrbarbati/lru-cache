package linked_list

type Node[T any] struct {
	Key   string
	Value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func NewNode[T any](key string, value T) *Node[T] {
	return &Node[T]{
		Key:   key,
		Value: value,
	}
}

func New[T any]() *DoublyLinkedList[T] {
	head := &Node[T]{}
	tail := &Node[T]{}

	head.next = tail
	tail.prev = head

	return &DoublyLinkedList[T]{
		head: head,
		tail: tail,
	}
}

func (l *DoublyLinkedList[T]) PushFront(node *Node[T]) {
	node.next = l.head.next
	node.prev = l.head
	l.head.next.prev = node
	l.head.next = node
}

func (l *DoublyLinkedList[T]) MoveToFront(node *Node[T]) {
	l.Remove(node)
	l.PushFront(node)
}

func (l *DoublyLinkedList[T]) Back() *Node[T] {
	if l.head.next == l.tail {
		return nil
	}

	return l.tail.prev
}

func (l *DoublyLinkedList[T]) Remove(node *Node[T]) {
	if l.head.next == l.tail {
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}
