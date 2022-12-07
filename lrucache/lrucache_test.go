package lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var lrucache *LRUCache

func setup() {
	lrucache = NewLRUCache(3)
	lrucache.Add(1)
	lrucache.Add(2)
	lrucache.Add(3)
}

func TestAddToCache(t *testing.T) {
	setup()
	val, err := lrucache.Get(1)
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	require.Equal(t, 1, val, "The value should be same")
}

func TestMaxCapacity(t *testing.T) {
	setup()
	lrucache.Add(4)
	require.Equal(t, 3, lrucache.CurrentCapacity(), "Capacity should not exceed max capacity")
}

func TestMostRecent(t *testing.T) {
	setup()
	lrucache.Get(1)
	lrucache.Add(4)
	lrucache.Add(5)
	lrucache.Print()
	got, err := lrucache.Get(1)
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	require.Equal(t, 1, got, "Most recent cache should not be evicted")
}
