package main

import "fmt"

//链表的综合操作,查找某个中间节点(经典快慢指针方法)
//删除第n个节点,删除倒数第n个节点

type node struct {
	value interface{}
	next  *node
}

type link struct {
	head *node
}

func (l *link) add(n int) {
	l.head = &node{ //构建头节点
		value: "i am head node",
		next:  nil,
	}

	pre := l.head
	//添加n个节点
	for i := 1; i <= n; i++ {
		node := &node{value: i, next: nil}
		pre.next = node
		pre = pre.next

	}

}

//查找中间节点(快慢指针)
func (l *link) find() *node {
	/* 首先确保节点数大于等于2 */
	//如果链表为空
	if l.head.next == nil || l.head == nil {
		return nil
	}
	//如果只有1个节点
	if l.head.next.next == nil {
		return l.head.next
	}

	//2个或者2个节点以上
	//快慢指针
	fast, slow := l.head, l.head
	//遍历到最后一个节点,此时慢指针就是中间元素,当fast为最后一个元素或者超出范围终止循环
	for fast != nil && fast.next != nil {
		slow = slow.next      //慢指针每次后移一个
		fast = fast.next.next //快指针每次位移两个节点
	}
	return slow
}

//删除第n个节点
func (l *link) del(num int) *link {
	if num <= 0 {
		num = 1
	}
	//num不能超过链表的长度
	pre := l.head
	//移动到num所在的位置的前一个node
	for i := 1; i <= num-1; i++ {
		pre = pre.next
	}
	pre.next = pre.next.next
	return l

}

//删除倒数第n个节点(快慢指针)
func (l *link) delRev(num int) {
	if nil == l.head || nil == l.head.next {
		return
	}
	if num <= 0 {
		num = 1
	}
	//首先使快指针直接向前移动num
	fast := l.head
	for i := 1; i <= num && fast != nil; i++ {
		fast = fast.next
	}
	if fast == nil {
		return
	}

	//使得快指针领先num个位置
	slow := l.head
	for fast.next != nil {
		//两个指针再同时后移,直到fast到末尾,出循环时slow指向倒数第num+1的位置,即待删除元素前一个位置
		//刚好慢指针会慢快指针num个身位,即slow位于待删除元素前一个位置
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	return

}
func main() {
	l := new(link)
	l.add(11)
	//遍历
	begin := l.head
	for begin != nil {
		fmt.Println(begin.value)
		begin = begin.next
	}
	//获得中间节点,快指针一次移动2步,慢指针一次移动1步,当快指针到末尾时,慢指针所在地就是中间位置
	//要注意每次先移动慢指针
	fmt.Println("中间节点", l.find().value)

	l = l.del(10) //删除第10个节点
	l.delRev(1)   //删除倒数第1个节点

	begin = l.head
	for begin != nil {
		fmt.Println(begin.value)
		begin = begin.next
	}

}
