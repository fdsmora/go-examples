package linkedlist

import "testing"

func TestLRUCache_BasicOperations(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	// This should evict key 2
	cache.Put(3, 3)

	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (evicted)", got)
	}

	// This should evict key 1
	cache.Put(4, 4)

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1 (evicted)", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLRUCache_SingleCapacity(t *testing.T) {
	cache := NewLRUCache(1)

	cache.Put(1, 1)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	cache.Put(2, 2)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1 (evicted)", got)
	}

	if got := cache.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}

func TestLRUCache_UpdateExistingKey(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(1, 10) // Update value for key 1

	if got := cache.Get(1); got != 10 {
		t.Errorf("Get(1) = %d, want 10", got)
	}

	// Key 2 should still exist
	if got := cache.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}
}

func TestLRUCache_GetNonExistentKey(t *testing.T) {
	cache := NewLRUCache(2)

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) from empty cache = %d, want -1", got)
	}

	cache.Put(1, 1)
	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) non-existent key = %d, want -1", got)
	}
}

func TestLRUCache_LRUEviction(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put(1, 1)
	cache.Put(2, 2)
	cache.Put(3, 3)

	// Access key 1 to make it recently used
	cache.Get(1)

	// Add new key, should evict key 2 (least recently used)
	cache.Put(4, 4)

	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (should be evicted)", got)
	}

	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1 (should still exist)", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3 (should still exist)", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4 (should still exist)", got)
	}
}

func TestLRUCache_ComplexScenario(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(2, 1)
	cache.Put(1, 1)
	cache.Put(2, 3) // Update key 2
	cache.Put(4, 1) // Evict key 1

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}

	if got := cache.Get(2); got != 3 {
		t.Errorf("Get(2) = %d, want 3", got)
	}
}

func TestLRUCache_GetPromotesToHead(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	// Access key 1 to make it most recently used
	cache.Get(1)

	// Add key 3, should evict key 2 (not key 1)
	cache.Put(3, 3)

	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1 (should not be evicted)", got)
	}

	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (should be evicted)", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}
}

func TestLRUCache_LargeCapacity(t *testing.T) {
	cache := NewLRUCache(100)

	// Fill cache
	for i := 1; i <= 100; i++ {
		cache.Put(i, i*10)
	}

	// Verify all values exist
	for i := 1; i <= 100; i++ {
		if got := cache.Get(i); got != i*10 {
			t.Errorf("Get(%d) = %d, want %d", i, got, i*10)
		}
	}

	// Add one more, should evict key 1
	cache.Put(101, 1010)

	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1 (should be evicted)", got)
	}

	if got := cache.Get(101); got != 1010 {
		t.Errorf("Get(101) = %d, want 1010", got)
	}
}

func TestLRUCache_ZeroAndNegativeValues(t *testing.T) {
	cache := NewLRUCache(3)

	cache.Put(1, 0)
	cache.Put(2, -5)
	cache.Put(3, -100)

	if got := cache.Get(1); got != 0 {
		t.Errorf("Get(1) = %d, want 0", got)
	}

	if got := cache.Get(2); got != -5 {
		t.Errorf("Get(2) = %d, want -5", got)
	}

	if got := cache.Get(3); got != -100 {
		t.Errorf("Get(3) = %d, want -100", got)
	}
}

func TestLRUCache_SequentialPutGet(t *testing.T) {
	cache := NewLRUCache(2)

	operations := []struct {
		op    string
		key   int
		value int
		want  int
	}{
		{"put", 1, 1, 0},
		{"put", 2, 2, 0},
		{"get", 1, 0, 1},
		{"put", 3, 3, 0},
		{"get", 2, 0, -1},
		{"put", 4, 4, 0},
		{"get", 1, 0, -1},
		{"get", 3, 0, 3},
		{"get", 4, 0, 4},
	}

	for i, op := range operations {
		if op.op == "put" {
			cache.Put(op.key, op.value)
		} else if op.op == "get" {
			got := cache.Get(op.key)
			if got != op.want {
				t.Errorf("Operation %d: Get(%d) = %d, want %d", i, op.key, got, op.want)
			}
		}
	}
}

func TestLRUCache_AlternatingPutGet(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	if got := cache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	cache.Put(2, 2)
	if got := cache.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2", got)
	}

	cache.Put(3, 3)
	if got := cache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1", got)
	}

	cache.Put(4, 4)
	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	if got := cache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}

	if got := cache.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}

func TestLRUCache_UpdateDoesNotEvict(t *testing.T) {
	cache := NewLRUCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)

	// Update key 1 multiple times - should not cause evictions
	cache.Put(1, 10)
	cache.Put(1, 100)
	cache.Put(1, 1000)

	if got := cache.Get(1); got != 1000 {
		t.Errorf("Get(1) = %d, want 1000", got)
	}

	if got := cache.Get(2); got != 2 {
		t.Errorf("Get(2) = %d, want 2 (should not be evicted)", got)
	}
}

func TestLRUCache_EdgeCaseCapacity10(t *testing.T) {
	cache := NewLRUCache(10)

	// Fill to capacity
	for i := 1; i <= 10; i++ {
		cache.Put(i, i*100)
	}

	// Access every other key
	for i := 2; i <= 10; i += 2 {
		cache.Get(i)
	}

	// Add 5 new keys - should evict odd keys (1, 3, 5, 7, 9)
	for i := 11; i <= 15; i++ {
		cache.Put(i, i*100)
	}

	// Odd keys should be evicted
	oddKeys := []int{1, 3, 5, 7, 9}
	for _, key := range oddKeys {
		if got := cache.Get(key); got != -1 {
			t.Errorf("Get(%d) = %d, want -1 (should be evicted)", key, got)
		}
	}

	// Even keys should still exist
	evenKeys := []int{2, 4, 6, 8, 10}
	for _, key := range evenKeys {
		if got := cache.Get(key); got != key*100 {
			t.Errorf("Get(%d) = %d, want %d", key, got, key*100)
		}
	}

	// New keys should exist
	for i := 11; i <= 15; i++ {
		if got := cache.Get(i); got != i*100 {
			t.Errorf("Get(%d) = %d, want %d\n", i, got, i*100)
		}
	}
}

func TestLRUCache_NilPointerException(t *testing.T) {
	cache := NewLRUCache(1)
	if got := cache.Get(6); got != -1 {
		t.Errorf("Get(6) = %d, want -1", got)
	}

	if got := cache.Get(8); got != -1 {
		t.Errorf("Get(8) = %d, want -1", got)
	}

	cache.Put(12, 1)

	if got := cache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	cache.Put(15, 11)
	cache.Put(5, 2)
	cache.Put(1, 15)
	cache.Put(4, 2)

	if got := cache.Get(4); got != 2 {
		t.Errorf("Get(4) = %d, want 2", got)
	}

	cache.Put(15, 15)

}
