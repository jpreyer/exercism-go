package listops

type (
	IntList   []int
	unaryFunc func(int) int
	binFunc   func(int, int) int
	predFunc  func(int) bool
)

//var (
//	list IntList
//)

func Foldr() {}

func Foldl() {}

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
	nl := []int{}
	for _, v := range ol {
		nl = append(nl, f(v))
	}
	return nl
}

func (i IntList) Reverse() IntList {
	l := i.Length()
	//better: t := make(IntList, l)
	var t IntList
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
