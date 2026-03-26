package linked_list

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestDoublyLinkedList_PushFront(t *testing.T) {
	scenarios := []struct {
		name       string
		list       *DoublyLinkedList[int]
		nodes      []*Node[int]
		verifyList func(*DoublyLinkedList[int]) error
	}{
		{
			name: "empty list",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("1", 1),
			},
			verifyList: func(l *DoublyLinkedList[int]) error {
				var i int
				curr := l.head.next
				expectedValues := []int{1}

				for curr != l.tail {
					if curr.Value != expectedValues[i] {
						return fmt.Errorf("expected %v, got %v at idx %v", expectedValues[i], curr.Value, i)
					}
					curr = curr.next
					i++
				}

				return nil
			},
		},
		{
			name: "multiple nodes",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("1", 1),
				NewNode("2", 2),
				NewNode("3", 3),
			},
			verifyList: func(l *DoublyLinkedList[int]) error {
				var i int
				curr := l.head.next
				expectedValues := []int{3, 2, 1}

				for curr != l.tail {
					if curr.Value != expectedValues[i] {
						return fmt.Errorf("expected %v, got %v at idx %v", expectedValues[i], curr.Value, i)
					}
					curr = curr.next
					i++
				}

				return nil
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			for _, node := range scenario.nodes {
				scenario.list.PushFront(node)
			}

			if err := scenario.verifyList(scenario.list); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDoublyLinkedList_Back(t *testing.T) {
	scenarios := []struct {
		name       string
		list       *DoublyLinkedList[int]
		nodes      []*Node[int]
		verifyBack func(*DoublyLinkedList[int]) error
	}{
		{
			name:  "empty list",
			list:  New[int](),
			nodes: []*Node[int]{},
			verifyBack: func(l *DoublyLinkedList[int]) error {
				if l.Back() != nil {
					return fmt.Errorf("expected nil, got %v", l.Back().Value)
				}

				return nil
			},
		},
		{
			name: "one item",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("1", 1),
			},
			verifyBack: func(l *DoublyLinkedList[int]) error {
				back := l.Back()

				if back == nil {
					return errors.New("expected non-nil, got nil")
				}

				if back.Value != 1 {
					return fmt.Errorf("expected 1, got %v", back.Value)
				}

				return nil
			},
		},
		{
			name: "multiple items",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("3", 3),
				NewNode("1", 1),
				NewNode("2", 2),
			},
			verifyBack: func(l *DoublyLinkedList[int]) error {
				back := l.Back()

				if back == nil {
					return errors.New("expected non-nil, got nil")
				}

				if back.Value != 3 {
					return fmt.Errorf("expected 3, got %v", back.Value)
				}

				return nil
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			for _, node := range scenario.nodes {
				scenario.list.PushFront(node)
			}

			if err := scenario.verifyBack(scenario.list); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestDoublyLinkedList_Remove(t *testing.T) {
	scenarios := []struct {
		name          string
		list          *DoublyLinkedList[int]
		nodes         []*Node[int]
		removeIndices []int
		verifyRemove  func(*DoublyLinkedList[int]) error
	}{
		{
			name:          "only element",
			list:          New[int](),
			nodes:         []*Node[int]{NewNode("1", 1)},
			removeIndices: []int{0},
			verifyRemove: func(l *DoublyLinkedList[int]) error {
				if l.head.next != l.tail {
					return errors.New("list should be empty after removing only element")
				}
				return nil
			},
		},
		{
			name: "non empty list",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("1", 1),
				NewNode("2", 2),
				NewNode("3", 3),
			},
			removeIndices: []int{1},
			verifyRemove: func(l *DoublyLinkedList[int]) error {
				var i int
				curr := l.head.next
				expectedValues := []int{3, 1}

				for curr != l.tail {
					if curr.Value != expectedValues[i] {
						return fmt.Errorf("unexpected value %v at index %v", curr.Value, i)
					}
					curr = curr.next
					i++
				}

				return nil
			},
		},
		{
			name: "multiple removes",
			list: New[int](),
			nodes: []*Node[int]{
				NewNode("1", 1),
				NewNode("2", 2),
				NewNode("3", 3),
			},
			removeIndices: []int{1, 0},
			verifyRemove: func(l *DoublyLinkedList[int]) error {
				var i int
				curr := l.head.next
				expectedValues := []int{3}

				for curr != l.tail {
					if curr.Value != expectedValues[i] {
						return fmt.Errorf("unexpected value %v at index %v", curr.Value, i)
					}
					curr = curr.next
					i++
				}

				return nil
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			for _, node := range scenario.nodes {
				scenario.list.PushFront(node)
			}

			for _, index := range scenario.removeIndices {
				scenario.list.Remove(scenario.nodes[index])
			}

			if err := scenario.verifyRemove(scenario.list); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func BenchmarkDoublyLinkedList_MoveToFront(b *testing.B) {
	list := New[int]()

	for i := 0; i < 10_000; i++ {
		list.PushFront(NewNode[int](strconv.Itoa(i), i))
	}
	back := list.Back()

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		list.MoveToFront(back)
	}
}

func BenchmarkDoublyLinkedList_Back(b *testing.B) {
	list := New[int]()

	for i := 0; i < 10_000; i++ {
		list.PushFront(NewNode[int](strconv.Itoa(i), i))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		list.Back()
	}
}

func BenchmarkDoublyLinkedList_Remove(b *testing.B) {
	list := New[int]()
	for j := 0; j < 100_000; j++ {
		list.PushFront(NewNode[int](strconv.Itoa(j), j))
	}

	b.ReportAllocs()
	b.ResetTimer()

	for b.Loop() {
		list.Remove(list.Back())
	}
}
