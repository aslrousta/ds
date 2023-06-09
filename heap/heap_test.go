package heap

import (
	"math/rand"
	"strconv"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Int int

func (i Int) Less(j Int) bool { return i < j }

var _ = Describe("Heap", func() {

	records := make(map[string]Int)
	for i := 0; i < 1000; i++ {
		records[strconv.Itoa(i)] = Int(rand.Int())
	}

	Describe("Push", func() {
		When("filled with random records", func() {
			h := New[string, Int]()
			for k, v := range records {
				h.Push(k, Int(v))
			}

			It("has the correct length", func() {
				Expect(h.Len()).To(Equal(len(records)))
			})
		})
	})

	Describe("Pop", func() {
		When("filled with random records", func() {
			h := New[string, Int]()
			for k, v := range records {
				h.Push(k, Int(v))
			}

			It("has the same items", func() {
				popped := 0
				for h.Len() > 0 {
					key, value, ok := h.Pop()
					Expect(ok).To(BeTrue())
					Expect(value).To(Equal(records[key]))
					popped++
				}
				Expect(popped).To(Equal(len(records)))
			})
		})
	})

	Describe("Remove", func() {
		When("filled with random records", func() {
			h := New[string, Int]()
			for k, v := range records {
				h.Push(k, Int(v))
			}

			It("has the same items", func() {
				for key := range records {
					value, ok := h.Remove(key)
					Expect(ok).To(BeTrue())
					Expect(value).To(Equal(records[key]))
				}
				Expect(h.Len()).To(BeZero())
			})
		})
	})

})

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Heap Suite")
}
