package main

type T struct {
	a []T
	b []*T
	c map[string]T
	d map[string]*T
	e chan T
	f chan *T
	h *T
	i func(T) T
	j func(*T) *T
	U
}

type U struct {
	k []T
	l []*T
	m map[string]T
	n map[string]*T
	o chan T
	p chan *T
	q *T
	r func(T) T
	s func(*T) *T
}

func main() {
	t := T{}
	u := U{}
	println(t)
	println(u)
}

// Output:
// struct{(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined)}
// struct{(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined),(undefined)}
