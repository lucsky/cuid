package cuid

type DefaultCounter struct {
	counterChan chan int64
}

func NewDefaultCounter() *DefaultCounter {
	counter := &DefaultCounter{make(chan int64)}
	go counter.Loop()
	<-counter.counterChan
	return counter
}

func (c *DefaultCounter) Next() int64 {
	return <-c.counterChan
}

func (c *DefaultCounter) Loop() {
	var count int64 = -1
	for {
		c.counterChan <- count
		count = count + 1
		if count >= discreteValues {
			count = 0
		}
	}
}
