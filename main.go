package main

import (
	"fmt"
	"strconv"
	"container/heap"
)


func main() {
	var input string

	fmt.Print("Enter N =")
	fmt.Scanf("%s",&input)
	N, _:= strconv.Atoi(input)

	uh := NewMinHeap(N)
	//heap.Init(uh)


	uh.Push(4)
	uh.Push(7)
	uh.Push(5)
	uh.Push(10)


	fmt.Printf("minimum: %d\n", uh.IntHeap[0])
	for uh.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(uh))
	}
        /**
	f, err := os.Open("number.txt")
	if (err != nil){
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0


	heap.Init(h)

	for scanner.Scan() {
		s := scanner.Text()
		num, _ := strconv.Atoi(s)



		if(count < N){
		//v = append(v, num)
		heap.Push(h, num)
		}else if(count == N){


		}

		if(count > N){

		}
		count++
	}**/

}
type Set interface {
	Add(value int)
	Contains(value int) (bool)
	Length() (int)
}
type MinHeap struct {
	IntHeap []int
}





type IntHeap []int

func (h MinHeap) Len() int           { return len(h.IntHeap) }
func (h MinHeap) Less(i, j int) bool { return h.IntHeap[i] < h.IntHeap[j] }
func (h MinHeap) Greater(i, j int) bool { return h.IntHeap[i] > h.IntHeap[j] }
func (h MinHeap) Swap(i, j int)      { h.IntHeap[i], h.IntHeap[j] = h.IntHeap[j], h.IntHeap[i] }
func (h *MinHeap) Peek()(x int){
	old := h.IntHeap
	n := len(old)
	x = old[n-1]
	return x
}
func (h *MinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	if(h.Greater(x.(int),h.Peek())){
	   h.Pop()
	   h.bubbleUp()
	   //*h = append(*h, x.(int))
	}
	for _, value := range values {
		h.list.Add(value)
	}
	size := h.Len()/2 + 1
	for i := size; i >= 0; i-- {
		h.bubbleDownIndex(i)
	}



}

func (h *MinHeap) Pop() interface{} {
	old := h.IntHeap
	n := len(old)
	x := old[n-1]
	h.IntHeap = old[0 : n-1]
	return x
}

// Performs the "bubble down" operation. This is to place the element that is at the index
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *MinHeap) bubbleDownIndex(index int) {
	size := heap.Len()
	for leftIndex := index<<1 + 1; leftIndex < size; leftIndex = index<<1 + 1 {
		rightIndex := index<<1 + 2
		smallerIndex := leftIndex
		leftValue := heap.IntHeap[leftIndex]
		rightValue := heap.IntHeap[rightIndex]
		if rightIndex < size && heap.Less(leftValue, rightValue){
			smallerIndex = rightIndex
		}
		indexValue := heap.IntHeap[index]
		smallerValue := heap.IntHeap[smallerIndex]
		if heap.Greater(indexValue, smallerValue)  {
			heap.Swap(index, smallerIndex)
		} else {
			break
		}
		index = smallerIndex
	}
}

// Performs the "bubble up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place so that
// the heap maintains the min/max-heap order property.
func (heap *MinHeap) bubbleUp() {
	index := heap.Len() - 1
	for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 {
		indexValue := heap.IntHeap[index]
		parentValue := heap.IntHeap[parentIndex]
		if heap.Less(parentValue, indexValue) {
			break
		}
		heap.Swap(index, parentIndex)
		index = parentIndex
	}
}
func NewMinHeap(size int) *MinHeap{
	return &MinHeap{ IntHeap: make([]int, size)}
}
