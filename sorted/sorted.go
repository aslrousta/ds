package sorted

import "github.com/aslrousta/ds"

// List implements a sorted list.
type List[T ds.Lesser[T]] []T

// Add inserts t into the sorted list.
func (l *List[T]) Add(t T) {
	*l = append(*l, t)
	for i, list := len(*l)-1, *l; i > 0; i-- {
		if !list[i].Less(list[i-1]) {
			break
		}
		list[i], list[i-1] = list[i-1], list[i]
	}
}

// IndexFunc returns the index of t, or returns -1 if no such item exists.
func (l List[T]) Index(t T) int {
	return l.IndexFunc(func(it T) int { return ds.Compare[T](t, it) })
}

// FirstIndex returns the index of the first occurrence of t, or returns -1 if
// no such item exists.
func (l List[T]) FirstIndex(t T) int {
	return l.FirstIndexFunc(func(it T) int { return ds.Compare[T](t, it) })
}

// LastIndex returns the index of the last occurrence of t, or returns -1 if no
// such item exists.
func (l List[T]) LastIndex(t T) int {
	return l.LastIndexFunc(func(it T) int { return ds.Compare[T](t, it) })
}

// IndexFunc returns the index of the item that cmp accepts, or returns -1 if no
// such item exists.
func (l List[T]) IndexFunc(cmp func(T) int) int {
	low, high := 0, len(l)-1
	for low <= high {
		mid := low + (high-low)/2
		switch r := cmp(l[mid]); {
		case r == 0:
			return mid
		case r < 0:
			high = mid - 1
		default /* res > 0 */ :
			low = mid + 1
		}
	}
	return -1
}

// FirstIndexFunc returns the index of the first item that cmp accepts, or
// returns -1 if no such item exists.
func (l List[T]) FirstIndexFunc(cmp func(T) int) int {
	p := l.IndexFunc(cmp)
	if p < 0 {
		return -1
	}
	for p > 0 {
		if cmp(l[p-1]) != 0 {
			break
		}
		p--
	}
	return p
}

// LastIndexFunc returns the index of the last item that cmp accepts, or returns
// -1 if no such item exists.
func (l List[T]) LastIndexFunc(cmp func(T) int) int {
	p := l.IndexFunc(cmp)
	if p < 0 {
		return -1
	}
	for p < len(l)-1 {
		if cmp(l[p+1]) != 0 {
			break
		}
		p++
	}
	return p
}
