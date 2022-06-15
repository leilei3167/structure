package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
实现单链表:
基本操作
	插入数据、删除数据、查找数据、求链表长度
辅助操作
	创建结点、创建（初始化 ）链表、判断是否为空链表

go中有GC 所以不用手动销毁指针或内存

*/

//创建节点的结构(真正存储数据的节点)
type Node struct {
	Data interface{}
	Next *Node
}

//链表的结构,代表着一个链表,其中包含一个Head头指针,指向头节点(如有),头节点不存放主要数据,主要是便于操作
type LList struct {
	Head   *Node //头指针:指向头节点,头节点仅便于操作,不算真正的节点
	Length int   //链表长度,不包含头节点
}

/* 辅助操作 */
//创建节点
func CreateNode(v interface{}) *Node {
	return &Node{Data: v, Next: nil}
}

//创建空链表(带一个头节点,头节点值为nil,链表头指针Head指向头节点)
func CreateList() *LList {
	return &LList{Head: CreateNode(nil), Length: 0} //以Data为nil创建一个头节点,不计入长度
}

//创建一个指定长度的随机链表,后插,类似于排队,后创建的在后面
func RandList(length int) *LList {
	rand.Seed(time.Now().UnixNano())

	list := CreateList()

	pre := list.Head

	for i := 0; i < length; i++ { //尾插
		nodes := CreateNode(rand.Intn(1000))
		fmt.Printf("第%d个插入的元素:%d\n", i, nodes.Data)
		pre.Next = nodes
		pre = pre.Next
		list.Length++
	}

	return list
}

//前插,新加入的node始终在第一位
func RandList2(length int) *LList {
	rand.Seed(time.Now().UnixNano())

	list := CreateList()

	pre := list.Head.Next //代表上一个创建的节点

	for i := 0; i < length; i++ { //前插
		nodes := CreateNode(rand.Intn(1000))
		//先指向上一个节点
		nodes.Next = pre
		fmt.Printf("第%d个插入的元素:%d\n", i, nodes.Data)
		//在让头节点指向待插入节点
		list.Head.Next = nodes

		//上一个节点更新为这个已插入的节点
		pre = nodes
		list.Length++
	}

	return list
}

/* 基本操作,基于链表结构 */
//增删第i个元素,i从1开始
func (list *LList) Insert(i int, v interface{}) {
	if i > list.Length+1 {
		panic("i大于链表长度")
	}
	//根据数据,构建待节点
	s := CreateNode(v)
	pre := list.Head //先获取到当前链表的第一个节点
	for count := 0; count <= i; count++ {
		if count == i-1 { //到第i个元素前面一个时
			s.Next = pre.Next //让待插入节点指向目标节点的下一个
			pre.Next = s      //再让目标节点只想待插入节点
			list.Length++
		}
		pre = pre.Next
	}

}

//删除第i处节点,当某个节点不在被另一些指针指向时,就会被GC回收,所以删除节点仅需要接触引用
func (list *LList) Delete(i int) {
	pre := list.Head //创建一个工作指针pre,指向链表头部第一个节点
	for count := 0; count < i; count++ {
		s := pre.Next     //临时工作变量
		if count == i-1 { //到达目标位置
			pre.Next = s.Next //跳过目标节点,直接连到下一个(被跳过的节点会被GC)
			list.Length--
		}
		pre = pre.Next
	}
}

func (list *LList) GetLen() int {
	return list.Length
}

//查询值v所在的位置,第i个,i从1开始
func (list *LList) Search(v interface{}) int {
	//从头开始,遍历整个链表
	pre := list.Head
	for i := 1; i <= list.Length; i++ {
		if pre.Data == v {
			return i
		}
		pre = pre.Next
	}
	return -1
}

//判断链表是否为空
func (list *LList) isNil() bool {
	pre := list.Head.Next //Head为头指针,指向头节点,头节点不算入节点,其Next才是真正的节点
	if pre != nil {
		return true
	}
	return false
}

//打印链表
func PrintList(list *LList) {
	//遍历打印
	pre := list.Head.Next //第一个真正存数据的节点
	fmt.Println("链表成员:")
	for i := 1; i <= list.Length; i++ {
		fmt.Printf("%v\n", pre.Data)
		pre = pre.Next
	}

}
func main() {

	/* 	list := CreateList()
	   	fmt.Println("List is null: ", list.isNil())

	   	list.Insert(1, 1000)
	   	list.Insert(1, 100)
	   	list.Insert(1, 10)
	   	list.Insert(1, 567)
	   	list.Insert(1, 9090)
	   	list.Insert(1, 890)
	   	list.Insert(1, 1000000000)

	   	PrintList(list)

	   	list.Delete(4)

	   	PrintList(list) */

	/* 	list2 := RandList(10)
	   	fmt.Println("list2的长度", list2.GetLen())
	   	PrintList(list2)
	   	list2.Delete(10)
	   	list2.Delete(1)
	   	list2.Delete(1)
	   	list2.Delete(1)

	   	PrintList(list2)
	   	fmt.Println("删除元素后list2的长度", list2.GetLen()) */

	list3 := RandList2(5)
	fmt.Println("list3的长度", list3.GetLen())
	PrintList(list3)

	list2 := RandList(5)
	fmt.Println("list2的长度", list2.GetLen())
	PrintList(list2)
}
