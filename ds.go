package ds

// Lesser defines types that can be compared to check for equality or to find
// the least. It is essential for ordered data-structures such as heaps, search
// trees and sorted sets.
type Lesser[T any] interface {
	comparable

	// Less returns true if this value is less than the other.
	Less(other T) bool
}

// Compare compares a and b and returns -1, 0 or 1 if a is less-than, equal-to
// or greater-than b, respectively.
func Compare[T Lesser[T]](a, b T) int {
	switch {
	case a == b:
		return 0
	case a.Less(b):
		return -1
	default:
		return 1
	}
}
