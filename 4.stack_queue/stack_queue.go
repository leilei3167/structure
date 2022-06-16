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

栈的应用:
	1.递归:
		直接调用自己或间接调用自己的函数称之为递归函数
		每个递归定义必须至少有一个退出条件,使其不再引用自身从而退出递归
	2.四则运算表达式处理(加减乘除括号等)----后缀表示(逆波兰)法
	因为括号是成对出现的,碰到左括号就入栈,不管有多少重括号,当出现右括号时,让栈顶的左括号出栈,
	期间让数字运算

		平时所用的表达式如:9+(3-1)*3+10/2 叫中缀表达式,所有运算符号都在中间,转化为后缀表达式为:
		9 3 1 - 3 * + 10 2 / +


二.队列
队列时只允许在一端进行插入,另一端进行删除的线性表
先进先出,插入的一端叫队尾,既然是一种线性表,队列也有顺序结构和链式结构

顺序结构的队列(用切片表示),在追加至队尾时复杂度为O(1),但是出队时后续都必须的前移,解决方案是构成一个循环队列


链式结构:
	插入时采用尾插的形式,出队时直接出第一个元素即可,如果保存有tail指针 可以使入队也是O(1)




总的来说:可以确定队列长度最大值时,建议使用循环队列,无法预估时,使用链队列
*/
//定义栈的方法集
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
	var s arrayStack
	s.array = make([]interface{}, 0)
	s.length = 0
	return &s
}
func (s *arrayStack) Pop() interface{} {
	top := s.GetTop()
	s.array = s.array[:s.length-1]
	s.length--
	return top
}

func (s *arrayStack) Push(v interface{}) {
	s.array = append(s.array, v)
	s.length++
}

func (s *arrayStack) GetLen() int {
	return s.length
}

func (s *arrayStack) GetTop() interface{} {
	if s.length < 1 {
		return nil
	}
	return s.array[len(s.array)-1]
}
func (s *arrayStack) Range() {
	fmt.Println("数组栈成员:")
	for i := s.length - 1; i >= 0; i-- {
		fmt.Println(s.array[i])
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
	var s stack

	//数组栈实现
	s = CreateArrayStack()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	fmt.Printf("栈顶元素值:%d 栈长度:%d\n", s.GetTop(), s.GetLen())
	s.Range()
	s.Pop()
	s.Pop()
	fmt.Println("pop后")
	s.Range()

	//链表栈
	s = CreateLinkStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	fmt.Printf("链表栈顶元素值:%d 链表栈长度:%d\n", s.GetTop(), s.GetLen())
	s.Range()
	s.Pop()
	s.Pop()
	s.Pop()

	fmt.Println("pop后")
	s.Range()
	//接口并不关注具体实现方式,只关注其是否有规定的方法

	fmt.Println("-----------循环队列-----------------")
	queue := createLoopQueue(5)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	/* fmt.Println("队列第一个元素:", queue.Pop())
	fmt.Println("队列第二个元素:", queue.Pop())
	fmt.Println("队列第三个元素:", queue.Pop()) */
	queue.Range()
	fmt.Printf("队列长度:%d 容量:%d\n", queue.length, queue.cap)
	fmt.Println("队列第一个元素:", queue.Pop())
	fmt.Println("队列第一个元素:", queue.Pop())
	fmt.Printf("队列长度:%d 容量:%d\n", queue.length, queue.cap)
}

/*
循环队列实现思路：
1.循环队列须要几个參数来确定
  front，tail，length，capacity
  front指向队列的第一个元素，tail指向队列最后一个元素的下一个位置
  length表示当前队列的长度，capacity标示队列最多容纳的元素
2.循环队列各个參数的含义
（1）队列初始化时，front和tail值都为零
（2）当队列不为空时，front指向队列的第一个元素，tail指向队列最后一个元素的下一个位置；
（3）当队列为空时，front与tail的值相等，但不一定为零
（4）当（tail+1）% capacity == front ||  （length+1）== capacity 表示队列为满，
    因此循环队列默认浪费1个空间
3.循环队列算法实现
（1）把值存在tail所在的位置；
（2）每插入1个元素，length+1，tail=（tail+1）% capacity
（3）每取出1个元素，length-1，front=（front+1）% capacity
（4）扩容功能，当队列容量满，即length+1==capacity时，capacity扩大为原来的2倍
（5）缩容功能，当队列长度小于容量的1/4，即length<=capacity/4时，capacity缩短为原来的一半
*/
//定义队列的方法集,和栈一样
type queue interface {
	Pop() interface{}
	Push(interface{})
	GetLen() int
	GetTop() interface{}
	Range()
}

//引入length和cap后非常便于操作判断
type loopQueue struct {
	queues []interface{}
	front  int //队首,队列不为空时,指向第一个元素
	tail   int //队尾,代表的是最后一个元素的下一个位置
	length int //元素个数
	cap    int //队列容量
}

//创建一个指定容量的循环队列
func createLoopQueue(n int) *loopQueue {
	lp := &loopQueue{
		queues: make([]interface{}, n), //容量为n
		front:  0,
		tail:   0,
		length: 0,
		cap:    n,
	}

	return lp
}

//普通队列取出时后续位置都必须前移,循环队列则不需要,取出时头指针后移
//操作复杂度都为O(1)
func (q *loopQueue) Pop() interface{} {
	if q.length == 0 {
		panic("没有数据")
	}
	res := q.queues[q.front]
	q.queues[q.front] = 0           //清空
	q.front = (q.front + 1) % q.cap //取出时将头标志后移
	q.length--
	return res
}

func (q *loopQueue) Push(v interface{}) {
	if q.length == q.cap {
		panic("队列已满 无法加入")
	}
	//加入是加入到尾部
	q.queues[q.tail] = v
	q.tail = (q.tail + 1) % q.cap //取模循环,加入的时候将尾部后移
	q.length++
	return

}

func (q *loopQueue) GetLen() int {
	return q.length
}

func (q *loopQueue) GetTop() interface{} {
	if q.length == 0 {
		panic("队列为空")
	}
	return q.queues[q.front] //返回队首
}

func (q *loopQueue) Range() {
	if q.length == 0 {
		panic("没有数据")
	}
	fmt.Println("队列成员:")
	temp := q.length
	temp2 := q.front
	for temp != 0 {
		fmt.Println(q.queues[temp2])
		temp2 = (temp2 + 1) % q.cap //后移
		temp--
	}

}
