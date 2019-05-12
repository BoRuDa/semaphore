package semaphore

import (
	"context"
	"testing"
	"time"
)

func TestOK(t *testing.T) {
	var (
		size    = 3
		sem     = New(size)
		ctx, _  = context.WithTimeout(context.TODO(), time.Millisecond*333)
		done    = make(chan struct{}, size)
		counter int
	)

	for i := 0; i < size; i++ {
		sem.Acquire()

		go func() {
			time.Sleep(time.Millisecond * 300)
			done <- struct{}{}
			sem.Release()
		}()
	}

	for {
		select {
		case <-done:
			counter++
			if counter == size {
				return
			}

		case <-ctx.Done():
			t.Fatal("timeout")
		}
	}
}
