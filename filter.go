package gollection

import "reflect"

func (g *gollection) Filter(f /* func(v <T>) bool */ interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.filterStream(f)
	}
	return g.filter(f)
}

func (g *gollection) filter(f interface{}) *gollection {
	sv, err := g.validateSlice("Filter")
	if err != nil {
		return &gollection{err: err}
	}

	funcValue, funcType, err := g.validateFilterFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if processFilter(funcValue, v) {
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) filterStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	funcValue, funcType, err := g.validateFilterFunc(f)
	if err != nil {
		return &gollection{
			err: err,
		}
	}

	var initialized bool
	go func() {
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- reflect.SliceOf(funcType.In(0))
						initialized = true
						continue
					}

					if processFilter(funcValue, reflect.ValueOf(v)) {
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
