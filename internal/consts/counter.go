package consts

import "sync"

type Counter struct {
	value int
	mux   sync.Mutex
}

func (c *Counter) Inc() {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value++
}

func (c *Counter) Set(val int) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.value = val
}

func (c *Counter) Value() int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.value
}
