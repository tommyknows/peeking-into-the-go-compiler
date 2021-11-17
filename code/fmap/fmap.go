package fmap

func fmap(f func(TypeA) TypeB, l []TypeA) []TypeB {
	lb := make([]TypeB, len(l))
	for i, a := range l {
		lb[i] = f(a)
	}
	return lb
}
