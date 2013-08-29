package lane

import (
	"fmt"
	"sync"
)

// PQType represents a priority queue ordering kind (see MAXPQ and MINPQ)
type PQType int

const (
	MAXPQ PQType = iota
	MINPQ
)

type item struct {
	value    interface{}
	priority int
}

// PQueue is a heap priority queue implementation. It can be
// whether max or min ordered and is safe
// for concurrent read-write operations.
type PQueue struct {
	sync.RWMutex
	items      []*item
	elemsCount int
	comparator func(int, int) bool
}

func newItem(value interface{}, priority int) *item {
	return &item{
		value:    value,
		priority: priority,
	}
}

func (i *item) String() string {
	return fmt.Sprintf("<item value:%s priority:%d>", i.value, i.priority)
}

// NewPQueue creates a new priority queue with the provided pqtype
// ordering type
func NewPQueue(pqType PQType) *PQueue {
	var cmp func(int, int) bool

	if pqType == MAXPQ {
		cmp = max
	} else {
		cmp = min
	}

	items := make([]*item, 1)
	items[0] = nil // Heap queue first element should always be nil

	return &PQueue{
		items:      items,
		elemsCount: 0,
		comparator: cmp,
	}
}

// Size returns the elements present in the priority queue count
func (pq *PQueue) Size() int {
	return pq.elemsCount
}

// Push the value item into the priority queue with provided priority.
func (pq *PQueue) Push(value interface{}, priority int) {
	item := newItem(value, priority)

	pq.Lock()
	pq.items = append(pq.items, item)
	pq.elemsCount += 1
	pq.swim(pq.Size())
	pq.Unlock()
}

// Pop and returns the highest/lowest priority item (depending on whether
// you're using a MINPQ or MAXPQ) from the priority queue
func (pq *PQueue) Pop() (interface{}, int) {
	if pq.Size() < 1 {
		return nil, 0
	}

	pq.Lock()

	var max *item = pq.items[1]

	pq.exch(1, pq.Size())
	pq.items = pq.items[0:pq.Size()]
	pq.elemsCount -= 1
	pq.sink(1)

	pq.Unlock()

	return max.value, max.priority
}

// Head returns the highest/lowest priority item (depending on whether
// you're using a MINPQ or MAXPQ) from the priority queue
func (pq *PQueue) Head() (interface{}, int) {
	if pq.Size() < 1 {
		return nil, 0
	}

	pq.RLock()
	headValue := pq.items[1].value
	headPriority := pq.items[1].priority
	pq.RUnlock()

	return headValue, headPriority
}

func max(i, j int) bool {
	return i < j
}

func min(i, j int) bool {
	return i > j
}

func (pq *PQueue) less(i, j int) bool {
	return pq.comparator(pq.items[i].priority, pq.items[j].priority)
}

func (pq *PQueue) exch(i, j int) {
	var tmpItem *item = pq.items[i]

	pq.items[i] = pq.items[j]
	pq.items[j] = tmpItem
}

func (pq *PQueue) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.exch(k/2, k)
	}

	k = k / 2
}

func (pq *PQueue) sink(k int) {
	for 2*k <= pq.Size() {
		var j int = 2 * k

		if j < pq.Size() && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}

		pq.exch(k, j)
		k = j
	}
}