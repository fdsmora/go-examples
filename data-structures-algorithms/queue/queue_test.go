package queue

import (
	"testing"
)

type Node struct {
	Value int
}

func TestQueue(t *testing.T) {
	// Test: Enqueue and Dequeue on an empty queue
	q := Queue[*Node]{}

	if q.Dequeue() != nil {
		t.Errorf("Expected Dequeue() to return nil for an empty queue")
	}

	// Test: Enqueue a single element and Dequeue it
	node1 := &Node{Value: 1}
	q.Enqueue(node1)

	if len(q) != 1 {
		t.Errorf("Expected queue length to be 1, got %d", len(q))
	}

	dequeuedNode := q.Dequeue()
	if dequeuedNode != node1 {
		t.Errorf("Expected Dequeue() to return the enqueued node, got %+v", dequeuedNode)
	}

	if len(q) != 0 {
		t.Errorf("Expected queue length to be 0 after Dequeue, got %d", len(q))
	}

	// Test: Enqueue multiple elements and Dequeue them in order
	node2 := &Node{Value: 2}
	node3 := &Node{Value: 3}
	node4 := &Node{Value: 4}

	q.Enqueue(node2)
	q.Enqueue(node3)
	q.Enqueue(node4)

	if len(q) != 3 {
		t.Errorf("Expected queue length to be 3, got %d", len(q))
	}

	if q.Dequeue() != node2 {
		t.Errorf("Expected Dequeue() to return node2")
	}

	if q.Dequeue() != node3 {
		t.Errorf("Expected Dequeue() to return node3")
	}

	if q.Dequeue() != node4 {
		t.Errorf("Expected Dequeue() to return node4")
	}

	if len(q) != 0 {
		t.Errorf("Expected queue length to be 0 after all Dequeues, got %d", len(q))
	}

	// Test: Dequeue on an empty queue after using it
	if q.Dequeue() != nil {
		t.Errorf("Expected Dequeue() to return nil when queue is empty")
	}

	// Test: Enqueue after Dequeue
	node5 := &Node{Value: 5}
	q.Enqueue(node5)

	if len(q) != 1 {
		t.Errorf("Expected queue length to be 1 after Enqueue, got %d", len(q))
	}

	if q.Dequeue() != node5 {
		t.Errorf("Expected Dequeue() to return node5")
	}
}
