package main

import "fmt"

//--------------------------------初始化----------------------------------------------
type node struct {
	value interface{}
	next  *node
}

type link struct { //代表一个链表,设置头指针和长度
	head   *node
	length int
}

//获得一个空链表
func GetLinkedList() *link {
	return &link{head: nil, length: 0}
}

//向空链表中加入元素
func (l *link) add(n int) *link {
	l.head = &node{ //构建头节点
		value: "i am head node",
		next:  nil,
	}

	pre := l.head
	//添加n个节点
	for i := 1; i <= n; i++ {
		node := &node{value: i, next: nil} //尾插
		pre.next = node
		l.length++
		pre = pre.next
	}
	return l
}

//使一个链表成环
func (l *link) makeRing() *link {
	//找到末尾,连接到头节点即可
	if l.length < 1 {
		panic("链表长度不足")
	}
	pre := l.head
	for pre.next != nil {
		pre = pre.next
	}
	pre.next = l.head

	return l
}

func (l *link) getLen() int {
	return l.length
}

func (l *link) addNode(v interface{}) { //头插一个元素
	pre := l.head
	node := &node{value: v, next: nil}
	node.next = pre.next
	pre.next = node
	l.length++
}

func (l *link) removeNode(num int) { //删除第i个节点

	if num > l.length || num < 1 {
		panic("输入不合法")
	}

	pre := l.head

	for i := 1; i <= num-1; i++ {

	}

	l.length--

}

//-------------------------------------------------------------------------------

//正向遍历链表
func (l *link) rangeList() {
	fmt.Println("链表的成员:")
	pre := l.head
	for i := 1; i <= l.length; i++ {
		fmt.Println(pre.next.value) //(忽略头节点)
		pre = pre.next
	}
}

/* 1.获得链表的中间节点 */

/* 2.获得链表倒数第n个节点 */

/* 3.从尾到头遍历列表 */
func (l *link) rangeList2() {
	fmt.Println("链表的成员(倒序):")
	temp := make([]*node, 0)
	pre := l.head
	for i := 1; i <= l.length; i++ {
		temp = append(temp, pre.next)
		pre = pre.next
	}
	for i := len(temp) - 1; i >= 0; i-- {
		fmt.Println(temp[i].value)
	}
}

/* 4.是否有环 */

/* 5.有环的话 求环的大小 */

/* 6.环的入口位置节点 */

/* 7.删除节点,要求时间复杂度为O(1) */

/* 8.两个链表的第一个公共节点(交叉处) */

/* 9.合并两个有序的链表 */

/* 10.链表复制 */

/* 11.链表反转 */
func main() {
	//创建一个链表,用于测试
	l := GetLinkedList().add(10)
	l.rangeList()
	l.rangeList2()

	l.addNode(10000000)
	l.addNode(1021)
	l.rangeList() //检测顶部插入
	fmt.Println("链表长度:", l.getLen())

}
