package listops

type (
	IntList   []int
	unaryFunc func(int) int
	binFunc   func(int, int) int
	predFunc  func(int) bool
)

func (i IntList) Foldl(bf binFunc, x int) int {
	ol := i
	if len(ol) == 0 {
		return x
	}

	// "peel" the first list element off and start with it
	holder := bf(ol[0], x)
	for _, v := range ol[1:] {
		holder = bf(holder, v)
	}
	return holder
}

func (i IntList) Foldr(bf binFunc, x int) int {
	ol := i.Reverse()
	if len(ol) == 0 {
		return x
	}

	// "peel" the first list element off and start with it
	holder := bf(ol[0], x)
	for _, v := range ol[1:] {
		holder = bf(v, holder)
	}
	return holder
}

func (i IntList) Filter(f predFunc) IntList {
	ol := i
	nl := []int{}
	for _, v := range ol {
		if f(v) {
			nl = append(nl, v)
		}
	}
	return nl
}

func (i IntList) Length() int {
	count := 0
	for range i {
		count++
	}
	return count
}

func (i IntList) Map(f unaryFunc) IntList {
	ol := i
	nl := make(IntList, i.Length())
	for j, v := range ol {
		nl[j] = f(v)
	}
	return nl
}

func (i IntList) Reverse() IntList {
	l := i.Length()
	t := make(IntList, l)
	for index, v := range i {
		t[l-1-index] = v
	}
	return t
}

func (i IntList) Append(l IntList) IntList {
	size := i.Length()
	nl := make(IntList, size+l.Length())

	for index, v := range i {
		nl[index] = v
	}
	for index, v := range l {
		nl[size+index] = v
	}
	return nl

}

func (i IntList) Concat(list []IntList) IntList {
	nl := i
	for _, v := range list {
		nl = append(nl, v...)
	}
	return nl

}
