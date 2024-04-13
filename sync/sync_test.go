package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("simple counter", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertcount(t, counter, 3)
	})

	t.Run("Concurrent counter", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertcount(t, counter, wantedCount)
	})
}

func assertcount(t *testing.T, counter *Counter, want int) {
	t.Helper()

	if counter.Value() != want {
		t.Errorf("got %d want %d", counter.Value(), want)
	}
}

func NewCounter() *Counter {
	return &Counter{}
}
