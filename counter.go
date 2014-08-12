package cuid

type DefaultCounter struct {
	counterChan chan int64
}

func NewDefaultCounter() *DefaultCounter {
	counter := &DefaultCounter{make(chan int64)}
	go counter.Loop()
	return counter
}

func (c *DefaultCounter) Next() int64 {
	return <-c.counterChan
}

func (c *DefaultCounter) Loop() {
	var count int64 = 0
	for {
		c.counterChan <- count
		count = count + 1
		if count >= discreteValues {
			count = 0
		}
	}
}
