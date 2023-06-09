package sorted

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Int int

func (i Int) Less(j Int) bool { return i < j }

type IntSlice []Int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var _ = Describe("List", func() {

	var records []Int
	for i := 0; i < 1000; i++ {
		records = append(records, Int(rand.Int()))
	}

	Describe("Add", func() {
		When("filled with random items", func() {
			var l List[Int]
			for _, r := range records {
				l.Add(r)
			}

			It("has items in sorted order", func() {
				sort.IsSorted(IntSlice(l))
			})
		})
	})

	Describe("Index", func() {
		When("filled with random items", func() {
			var l List[Int]
			for _, r := range records {
				l.Add(r)
			}

			It("returns -1 if item does not exist", func() {
				Expect(l.Index(-3)).To(Equal(-1))
			})
			It("returns the index of an existing item", func() {
				r := rand.Intn(len(records))
				pos := l.Index(records[r])
				Expect(pos).To(BeNumerically(">=", 0))
				Expect(l[pos]).To(Equal(records[r]))
			})
		})
	})

	Describe("FirstIndex", func() {
		When("filled with non-unique items", func() {
			var l List[Int]
			for _, r := range []Int{1, 3, 5, 4, 2, 2, 6, 1, 2} {
				l.Add(r)
			}

			It("returns the first index of an existing item", func() {
				Expect(l.FirstIndex(2)).To(Equal(2))
			})
		})
	})

	Describe("LastIndex", func() {
		When("filled with non-unique items", func() {
			var l List[Int]
			for _, r := range []Int{1, 3, 5, 4, 2, 2, 6, 1, 2} {
				l.Add(r)
			}

			It("returns the last index of an existing item", func() {
				Expect(l.LastIndex(2)).To(Equal(4))
			})
		})
	})

})

func TestList(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sorted List Suite")
}
