package gollection

import (
	"fmt"
	"reflect"
)

func (g *gollection) Skip(n int) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.skipStream(n)
	}

	return g.skip(n)
}

func (g *gollection) skip(n int) *gollection {
	sv, err := g.validateSlice("Take")
	if err != nil {
		return &gollection{err: err}
	}

	if n < 0 {
		return &gollection{err: fmt.Errorf("gollection.Skip called with invalid argument. should be larger than 0")}
	}

	limit := sv.Len()
	start := n
	if limit < start {
		start = limit
	}

	ret := reflect.MakeSlice(sv.Type(), 0, limit-start)

	for i := start; i < limit; i++ {
		ret = reflect.Append(ret, sv.Index(i))
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) skipStream(n int) *gollection {
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

				if ok {
					i++
					if n < i {
						next.ch <- v
					}
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
