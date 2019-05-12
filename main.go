package semaphore

type Semaphore interface {
	Acquire()
	Release()
}

type semaphore struct {
	res chan struct{}
}

func New(size int) Semaphore {
	return &semaphore{
		res: make(chan struct{}, size),
	}
}

func (s *semaphore) Acquire() {
	s.res <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.res
}
