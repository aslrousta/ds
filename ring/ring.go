package ring

// Ring implements a circular queue.
type Ring[T comparable] struct {

	// MinFillRate is the minimum fill-rate in percents. When the fill-rate
	// reaches below this threshold, it triggers an automatic compaction.
	// Setting a high value can reduce the memory usage significantly but may
	// cause too many re-allocations.
	MinFillRate int

	items      []T
	head, tail int
}

// Head returns the head element.
func (r *Ring[T]) Head() (t T, ok bool) {
	if len(r.items) == 0 {
		return t, false
	}
	return r.items[r.head], true
}

// MustHead returns the head element. It panics if the ring is empty.
func (r *Ring[T]) MustHead() T {
	if t, ok := r.Head(); ok {
		return t
	}
	panic("ring: head does not exist")
}

// Tail returns the tail element.
func (r *Ring[T]) Tail() (t T, ok bool) {
	if len(r.items) == 0 {
		return t, false
	}
	pos := r.tail - 1
	if pos < 0 {
		pos += len(r.items)
	}
	return r.items[pos], true
}

// MustTail returns the tail element. It panics if the ring is empty.
func (r *Ring[T]) MustTail() T {
	if t, ok := r.Tail(); ok {
		return t
	}
	panic("ring: tail does not exist")
}

// Push appends t at the tail.
func (r *Ring[T]) Push(t T) {
	switch {
	case r.tail != r.head:
		r.items[r.tail] = t
		r.tail = (r.tail + 1) % len(r.items)
	default:
		r.items = append(r.items, t)
		if r.tail > 0 {
			copy(r.items[r.head+1:], r.items[r.head:])
			r.head++
			r.items[r.tail] = t
			r.tail++
		}
	}
}

// Pop removes the head and returns it.
func (r *Ring[T]) Pop() (t T, ok bool) {
	if len(r.items) == 0 {
		return t, false
	}
	defer func() {
		r.head = (r.head + 1) % len(r.items)
		if r.head == r.tail {
			r.items = nil
		} else if r.FillRate() < r.MinFillRate {
			r.Compact()
		}
	}()
	return r.items[r.head], true
}

// Remove searches for t and returns true if t is found and deleted.
func (r *Ring[T]) Remove(t T) bool {
	pos := -1
	for i := 0; i < r.Len(); i++ {
		if cur := (r.head + i) % len(r.items); r.items[cur] == t {
			pos = cur
			break
		}
	}
	if pos < 0 {
		return false
	}
	defer func() {
		if r.head == r.tail {
			r.items = nil
		} else if r.FillRate() < r.MinFillRate {
			r.Compact()
		}
	}()
	lastPos := (r.tail - 1 + len(r.items)) % len(r.items)
	switch {
	case pos == r.head:
		r.head = (r.head + 1) % len(r.items)
	case pos == lastPos:
		r.tail = lastPos
	case pos > r.head:
		copy(r.items[r.head+1:], r.items[r.head:pos])
		r.head++
	default /* pos < r.tail */ :
		copy(r.items[pos:], r.items[pos+1:r.tail])
		r.tail--
	}
	return true
}

// Len returns the number of elements.
func (r *Ring[T]) Len() int {
	size := r.tail - r.head
	if size <= 0 {
		size += len(r.items)
	}
	return size
}

// FillRate returns the used space in percents.
func (r *Ring[T]) FillRate() int {
	if len(r.items) == 0 {
		return 0
	}
	return r.Len() * 100 / len(r.items)
}

// Compact releases the unused space.
func (r *Ring[T]) Compact() {
	if r.head == r.tail {
		return
	}
	tmp := make([]T, r.Len())
	switch {
	case r.tail == 0:
		copy(tmp, r.items[r.head:])
	case r.tail > r.head:
		copy(tmp, r.items[r.head:r.tail])
	default:
		n := copy(tmp, r.items[r.head:])
		copy(tmp[n:], r.items[:r.tail])
	}
	r.items = tmp
	r.head = 0
	r.tail = 0
}
