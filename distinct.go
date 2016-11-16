package gollection

import "reflect"

func (g *gollection) Distinct() *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.distinctStream()
	}

	return g.distinct()

}

func (g *gollection) DistinctBy(f /*func(v <T1>) <T2>*/ interface{}) *gollection {
	if g.err != nil {
		return &gollection{err: g.err}
	}

	if g.ch != nil {
		return g.distinctByStream(f)
	}

	return g.distinctBy(f)
}

func (g *gollection) distinct() *gollection {
	sv, err := g.validateSlice("Distinct")
	if err != nil {
		return &gollection{err: err}
	}

	ret := reflect.MakeSlice(sv.Type(), 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if processDistinct(v.Interface(), m) {
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) distinctStream() *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	go func() {
		var initialized bool
		m := make(map[interface{}]bool)
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- v.(reflect.Type)
						initialized = true
						continue
					}

					if processDistinct(v, m) {
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

func (g *gollection) distinctBy(f interface{}) *gollection {
	sv, err := g.validateSlice("DistinctBy")
	if err != nil {
		return &gollection{err: err}
	}

	funcValue, funcType, err := g.validateDistinctByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	resultSliceType := reflect.SliceOf(funcType.In(0))
	ret := reflect.MakeSlice(resultSliceType, 0, sv.Len())
	m := make(map[interface{}]bool)

	for i := 0; i < sv.Len(); i++ {
		v := sv.Index(i)
		if processDistinctBy(funcValue, v, m) {
			ret = reflect.Append(ret, v)
		}
	}

	return &gollection{
		slice: ret.Interface(),
	}
}

func (g *gollection) distinctByStream(f interface{}) *gollection {
	next := &gollection{
		ch: make(chan interface{}),
	}

	fv, _, err := g.validateDistinctByFunc(f)
	if err != nil {
		return &gollection{err: err}
	}

	go func(fv *reflect.Value) {
		var initialized bool
		m := make(map[interface{}]bool)
		for {
			select {
			case v, ok := <-g.ch:
				if ok {
					if !initialized {
						// initialize next stream type
						next.ch <- v.(reflect.Type)
						initialized = true
						continue
					}

					if processDistinctBy(*fv, reflect.ValueOf(v), m) {
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
	}(&fv)

	return next
}
