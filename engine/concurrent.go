package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
}

func (c *ConcurrentEngine)Run(seeds ...Request)  {
	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
}

