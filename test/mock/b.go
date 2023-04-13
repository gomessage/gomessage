package mock

type Retrievers struct {
	Context string
}

func (r Retrievers) Get(cur string) string {
	return r.Context
}
