package ring

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ring", func() {

	Describe("Head", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("returns an empty result", func() {
				_, ok := r.Head()
				Expect(ok).To(BeFalse())
			})
		})
		When("the ring is non-empty", func() {
			r := Ring[int]{
				items: []int{10, 20},
			}

			It("returns the head", func() {
				n, ok := r.Head()
				Expect(ok).To(BeTrue())
				Expect(n).To(Equal(10))
			})
		})
	})

	Describe("Tail", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("returns an empty result", func() {
				_, ok := r.Tail()
				Expect(ok).To(BeFalse())
			})
		})
		When("the ring is non-empty", func() {
			r := Ring[int]{
				items: []int{10, 20},
			}

			It("returns the tail", func() {
				n, ok := r.Tail()
				Expect(ok).To(BeTrue())
				Expect(n).To(Equal(20))
			})
		})
	})

	Describe("Push", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("appends the only element", func() {
				r.Push(10)
				Expect(r.Len()).To(Equal(1))
				Expect(r.MustTail()).To(Equal(10))
			})
		})
		When("the ring is non-empty", func() {
			r := Ring[int]{
				items: []int{10, 20},
			}

			It("appends the element at the tail", func() {
				r.Push(30)
				Expect(r.Len()).To(Equal(3))
				Expect(r.MustTail()).To(Equal(30))
			})
		})
		When("the tail precedes the head", func() {
			r := Ring[int]{
				items: []int{10, 0, 20},
				head:  2,
				tail:  1,
			}

			It("appends the element at the tail", func() {
				r.Push(30)
				Expect(r.Len()).To(Equal(3))
				Expect(r.MustTail()).To(Equal(30))
			})
		})
		When("the tail follows the head", func() {
			r := Ring[int]{
				items: []int{10, 20, 0},
				tail:  2,
			}

			It("appends the element at the tail", func() {
				r.Push(30)
				Expect(r.Len()).To(Equal(3))
				Expect(r.MustTail()).To(Equal(30))
			})
		})
		When("the tail collides with the head", func() {
			r := Ring[int]{
				items: []int{10, 20, 30},
				head:  1,
				tail:  1,
			}

			It("appends the element at the tail", func() {
				r.Push(40)
				Expect(r.Len()).To(Equal(4))
				Expect(r.MustTail()).To(Equal(40))
			})
		})
	})

	Describe("Pop", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("returns an empty result", func() {
				_, ok := r.Pop()
				Expect(ok).To(BeFalse())
			})
		})
		When("the ring is non-empty", func() {
			r := Ring[int]{
				items: []int{10, 20},
			}

			It("removes the head", func() {
				n, ok := r.Pop()
				Expect(ok).To(BeTrue())
				Expect(n).To(Equal(10))
				Expect(r.MustHead()).To(Equal(20))
			})
		})
		When("the fill-rate falls below the minimum", func() {
			r := Ring[int]{
				MinFillRate: 50,
				items:       []int{10, 20, 0},
				tail:        2,
			}

			It("triggers an automatic compaction", func() {
				r.Pop()
				Expect(r.Len()).To(Equal(1))
				Expect(r.MustHead()).To(Equal(20))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
	})

	Describe("Remove", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("does nothing", func() {
				Expect(r.Remove(10)).To(BeFalse())
			})
		})
		When("removing the head", func() {
			r := Ring[int]{
				items: []int{10, 20, 30},
			}

			It("removes the element", func() {
				Expect(r.Remove(10)).To(BeTrue())
				Expect(r.Len()).To(Equal(2))
				Expect(r.MustHead()).To(Equal(20))
			})
		})
		When("removing the tail", func() {
			r := Ring[int]{
				items: []int{10, 20, 30},
			}

			It("removes the element", func() {
				Expect(r.Remove(30)).To(BeTrue())
				Expect(r.Len()).To(Equal(2))
				Expect(r.MustTail()).To(Equal(20))
			})
		})
		When("removing the middle", func() {
			r := Ring[int]{
				items: []int{40, 50, 60, 10, 20, 30},
				head:  3,
				tail:  3,
			}

			It("removes the element", func() {
				Expect(r.Remove(20)).To(BeTrue())
				Expect(r.Len()).To(Equal(5))
				Expect(r.MustHead()).To(Equal(10))
				Expect(r.MustTail()).To(Equal(60))
				Expect(r.Remove(50)).To(BeTrue())
				Expect(r.Len()).To(Equal(4))
				Expect(r.MustHead()).To(Equal(10))
				Expect(r.MustTail()).To(Equal(60))
			})
		})
		When("the fill-rate falls below the minimum", func() {
			r := Ring[int]{
				MinFillRate: 50,
				items:       []int{10, 20, 0},
				tail:        2,
			}

			It("triggers an automatic compaction", func() {
				r.Remove(10)
				Expect(r.Len()).To(Equal(1))
				Expect(r.MustHead()).To(Equal(20))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
	})

	Describe("Len", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("returns 0", func() {
				Expect(r.Len()).To(BeZero())
			})
		})
		When("an element is pushed", func() {
			r := Ring[int]{
				items: []int{10},
			}

			It("increases the length by one", func() {
				r.Push(20)
				Expect(r.Len()).To(Equal(2))
			})
		})
		When("an element is popped", func() {
			r := Ring[int]{
				items: []int{10},
			}

			It("decreases the length by one", func() {
				r.Pop()
				Expect(r.Len()).To(BeZero())
			})
		})
	})

	Describe("FillRate", func() {
		When("the ring is empty", func() {
			r := Ring[int]{}

			It("returns 0%", func() {
				Expect(r.FillRate()).To(BeZero())
			})
		})
		When("2/3 is occupied", func() {
			r := Ring[int]{
				items: []int{10, 20, 0},
				tail:  2,
			}

			It("returns 66%", func() {
				Expect(r.FillRate()).To(Equal(66))
			})
		})
	})

	Describe("Compact", func() {
		When("the ring is full", func() {
			r := Ring[int]{
				items: []int{10, 20, 30},
			}

			It("does nothing", func() {
				r.Compact()
				Expect(r.Len()).To(Equal(3))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
		When("the right-side is empty", func() {
			r := Ring[int]{
				items: []int{0, 10, 20},
				head:  1,
			}

			It("compacts the ring", func() {
				r.Compact()
				Expect(r.Len()).To(Equal(2))
				Expect(r.MustHead()).To(Equal(10))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
		When("the left-side is empty", func() {
			r := Ring[int]{
				items: []int{10, 20, 0},
				tail:  2,
			}

			It("compacts the ring", func() {
				r.Compact()
				Expect(r.Len()).To(Equal(2))
				Expect(r.MustHead()).To(Equal(10))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
		When("the middle is empty", func() {
			r := Ring[int]{
				items: []int{20, 0, 10},
				head:  2,
				tail:  1,
			}

			It("compacts the ring", func() {
				r.Compact()
				Expect(r.Len()).To(Equal(2))
				Expect(r.MustHead()).To(Equal(10))
				Expect(r.FillRate()).To(Equal(100))
			})
		})
	})

})

func TestRing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ring Suite")
}
