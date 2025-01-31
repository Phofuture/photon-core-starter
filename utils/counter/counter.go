package counter

import "time"

func NewCounter(duration time.Duration) *Counter {
	return &Counter{
		ticker: time.NewTicker(duration),
		isStop: make(chan struct{}),
	}
}

type Counter struct {
	ticker *time.Ticker
	isStop chan struct{}
}

func (c *Counter) Stop() {
	c.ticker.Stop()
	close(c.isStop)
}

func (c *Counter) Touch() <-chan time.Time {
	return c.ticker.C
}

func (c *Counter) IsStop() chan struct{} {
	return c.isStop
}
