package main

import (
	"fmt"
	"strconv"
	"container/heap"
	"os"
	"bufio"
)


func main() {
	var input string

	fmt.Print("Enter N for top N numbers =")
	fmt.Scanf("%s",&input)
	N, _:= strconv.Atoi(input)

	f, err := os.Open("number.txt")
	if (err != nil){
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	buff := make([]byte, 0, 1024*1024)
	scanner.Buffer(buff,10*1024*1024)
	count := 0

	v := &IntHeap{}

	for scanner.Scan() {
		s := scanner.Text()
		num, _ := strconv.Atoi(s)



		if(count < N && v.notADup(num)){
		  *v = append(*v, num)
		}else if(count == N){
		  heap.Init(v)
		}

		if(count >= N && v.notADup(num)){
		  heap.Push(v,num)
		}
		count++
	}

	//Cases when user input is greater than file length
	if(N > count){
	   heap.Init(v)
	}
	for v.Len() >= 1 {

		fmt.Printf("%d ", heap.Pop(v))

	}
}


type IntHeap []int

// Peek allows for peek at the smallest number in the heap
func (h *IntHeap) Peek()(x int){
	old := *h
	x = old[0]
	return x
}

func (h IntHeap) Len() int           { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


// Push before inserts checks for duplicates and removes the smallest number to input the largest,
func (h *IntHeap) Push(x interface{}) {


	if(x.(int) > h.Peek()){
	   heap.Pop(h)
	   *h = append(*h, x.(int))
	}

}
func (h *IntHeap)notADup(x int) bool{
	for i := 0; i < h.Len(); i++ {
		if((*h)[i] == x){
		   return false
		}
	}
	return true
}


func (h *IntHeap) Pop() interface{} {

	old := *h

	n := len(old)

	x := old[n-1]

	*h = old[0 : n-1]

	return x

}