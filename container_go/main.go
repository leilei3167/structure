package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

//go的container包实现了3个数据结构,可以学习其实现方式!

func main() {
	/* List:一个双向的链表 */
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)   //在链表最后添加
	e1 := l.PushFront(1)  //在链表最前添加
	l.InsertBefore(3, e4) //在e4之前添加
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	//Front获取表头元素,循环迭代
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	/* Ring:双向循环链表 */
	//创建一个双向循环链表(实际使用少)
	r := ring.New(5)

	// Get the length of the ring
	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next() //获取下一个元素
	}

	// Iterate through the ring and print its contents
	//Do会遍历元素,并且将每个元素作为参数执行函数
	r.Do(func(p any) {
		fmt.Println(p.(int))
	})
	//Link将两个环链表连接起来(拼接),Unlink将两个环链表断开(截取),Move跳到链表任意位置

	/*Heap:实现了一个堆,堆是一个特殊的树,每个节点都是其子树中的最小值节点(最小堆)
		堆是用于实现优先级队列的常用方法!
		使用此包时需先实现接口:
		type Interface interface {
		sort.Interface
		Push(x any) // add x as element Len()
		Pop() any   // remove and return element Len() - 1.
	}

	实现了以上接口的对象可以被作为堆来使用


	*/
	h := &IntHeap{2, 1, 5, 3, 4}
	heap.Init(h)

	heap.Push(h, 6)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
