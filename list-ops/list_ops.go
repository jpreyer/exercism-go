package listops

type (
	IntList   []int
	unaryFunc func(int) int
	binFunc   func(int, int) int
	predFunc  func(int) bool
)

var (
	list IntList
)

func Foldr() {}

func Foldl() {}

func Filter() {}

func Length(i IntList) int {
	return len(i)

}

func Map() {}

func Reverse() {}

func Append() {}

func Concat() {}
