package parallel

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

func TestRunWorkers(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	t.Run("runs the specified number of workers", func(t *testing.T) {
		workers := rand.Intn(8) + 1 + 8

		workersRunCount := make([]int32, workers)

		RunWorkers(workers, func(workerNum, workerCount int) {
			atomic.AddInt32(&workersRunCount[workerNum], 1)
		})

		for i := range workersRunCount {
			if expected, actual := int32(1), workersRunCount[i]; expected != actual {
				t.Errorf("Expected worker %d to have run once but got %d times", i, actual)
			}
		}
	})

	t.Run("passes the correct worker count to workers", func(t *testing.T) {
		workers := rand.Intn(8) + 1 + 8

		workersCount := make([]int, workers)

		RunWorkers(workers, func(workerNum, workerCount int) {
			workersCount[workerNum] = workerCount
		})

		for i := range workersCount {
			if expected, actual := workers, workersCount[i]; expected != actual {
				t.Errorf("Expected worker %d to have total worker count %d but got %d", i, expected, actual)
			}
		}
	})

	t.Run("returns after all workers are done", func(t *testing.T) {
		workers := rand.Intn(4) + 1 + 4

		workerResults := make([]int, workers)
		workerResultOffset := int32(-1)

		RunWorkers(workers, func(workerNum, workerCount int) {
			time.Sleep(100 * time.Duration(workerNum) * time.Millisecond)

			offset := atomic.AddInt32(&workerResultOffset, 1)
			workerResults[offset] = workerNum
		})

		for i := range workerResults {
			if expected, actual := i, workerResults[i]; expected != actual {
				t.Errorf("Expected worker result %d to be %d but was %d", i, expected, actual)
			}
		}
	})
}
