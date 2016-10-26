package gollection

type gollection struct {
	slice interface{}
	val   interface{}
	err   error
}

func New(slice interface{}) *gollection {
	return &gollection{
		slice: slice,
	}
}

func (g *gollection) Result() (interface{}, error) {
	if g.val != nil {
		return g.val, g.err
	}
	return g.slice, g.err
}
