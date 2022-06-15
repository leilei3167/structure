package main

import "fmt"

/*
一.栈
栈是限定仅在表尾进行插入和删除操作的线性表
允许插入删除的一端称为栈顶top,另一端称为栈底,栈简称(LIFO,last in first out)结构

注意:栈是一个线性表,具有线性关系(前驱后继),表尾指的是栈顶

栈的抽象数据类型:
ADT 栈(stack)
Data
	同线性表
Operation
	InitStack(*S)初始化,创建一个空栈
	清空栈
	是否是空栈
	返回栈顶元素
	Push入栈一个元素
	Pop弹出栈顶元素,并删除
	返回长度
endADT

由于栈本身就是线性表,因此栈可以使用顺序结构和链式结构来实现



*/

type stack interface {
	Pop() interface{}
	Push(interface{})
	GetLen() int
	GetTop() interface{}
	Range()
}

//数组栈
type arrayStack struct {
	array  []interface{}
	length int
}

func CreateArrayStack() *arrayStack {
	var arr arrayStack
	arr.array = make([]interface{}, 0)
	arr.length = 0
	return &arr
}
func (arr *arrayStack) Pop() interface{} {
	top := arr.GetTop()
	arr.array = arr.array[:arr.length-1]
	arr.length--
	return top
}

func (arr *arrayStack) Push(v interface{}) {
	arr.array = append(arr.array, v)
	arr.length++
}

func (arr *arrayStack) GetLen() int {
	return arr.length
}

func (arr *arrayStack) GetTop() interface{} {
	if arr.length < 1 {
		return nil
	}
	return arr.array[len(arr.array)-1]
}
func (arr *arrayStack) Range() {
	fmt.Println("数组栈成员:")
	for i := arr.length - 1; i >= 0; i-- {
		fmt.Println(arr.array[i])
	}
}

//------------------------链表栈-----------------------------------
type linkStack struct {
	Next   *Node //对于链表栈来说,头节点不需要也可以
	Length int
}

type Node struct {
	Data interface{}
	Next *Node
}

//创建空链表,头指针指向一个data为nil的头节点
func CreateLinkStack() *linkStack {
	return &linkStack{Next: nil, Length: 0} //不需要头节点,Next直接设为nil
}

func CreateNode(v interface{}) *Node {
	return &Node{Data: v, Next: nil}
}

func (l *linkStack) Pop() interface{} {
	res := l.Next
	l.Next = l.Next.Next
	l.Length--
	return res.Data
}

func (l *linkStack) Push(v interface{}) {
	//构建节点
	node := CreateNode(v)
	//需要将新元素加入到头部,前插
	node.Next = l.Next
	l.Next = node
	l.Length++
}

func (l *linkStack) GetLen() int {
	return l.Length
}

func (l *linkStack) GetTop() interface{} {
	return l.Next.Data
}

func (l *linkStack) Range() {
	fmt.Println("链表栈成员:")
	pre := l.Next
	for i := 1; i <= l.Length; i++ {
		fmt.Println(pre.Data)
		pre = pre.Next
	}
}

func main() {
	//数组栈实现
	arr := CreateArrayStack()

	arr.Push(1)
	arr.Push(2)
	arr.Push(3)
	arr.Push(4)
	arr.Push(5)

	fmt.Printf("栈顶元素值:%d 栈长度:%d\n", arr.GetTop(), arr.GetLen())
	arr.Range()
	arr.Pop()
	arr.Pop()
	fmt.Println("pop后")
	arr.Range()

	//链表栈
	link := CreateLinkStack()
	link.Push(1)
	link.Push(2)
	link.Push(3)
	link.Push(4)
	link.Push(5)
	fmt.Printf("链表栈顶元素值:%d 链表栈长度:%d\n", link.GetTop(), link.GetLen())
	link.Range()
	link.Pop()
	link.Pop()
	link.Pop()

	fmt.Println("pop后")
	link.Range()
}
