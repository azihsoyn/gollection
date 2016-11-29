package gollection

import "reflect"

func (g *gollection) Take(n int) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.takeStream(n)
	}

	return g.take(n)
}

func (g *gollection) take(n int) *gollection {
	sv, err := g.validateSlice("Take")
	if err != nil {
		return &gollection{err: err}
	}

	limit := sv.Len()
	if n < limit {
		limit = n
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())

	for i := 0; i < limit; i++ {
		ret = reflect.Append(ret, sv.Index(i))
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) takeStream(n int) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	var initialized bool
	go func() {
		i := 0
		for {
			select {
			case v, ok := <-g.ch:
				// initialize next stream type
				if ok && !initialized {
					next.ch <- v
					initialized = true
					continue
				}

				if ok && i < n {
					next.ch <- v
					i++
				} else {
					close(next.ch)
					return
				}
			default:
				continue
			}
		}
	}()
	return next
}
