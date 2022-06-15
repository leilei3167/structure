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
	if l.length < 2 {
		panic("链表长度不足")
	}
	pre := l.head
	for pre.next != nil {
		pre = pre.next
	}
	pre.next = l.head.next.next.next //避免首位相连

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
		pre = pre.next
	}
	pre.next = pre.next.next
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
func (l *link) getCenterNode() *node {
	if l.length == 0 {
		return nil
	}
	if l.length == 1 {
		return l.head.next
	}
	//快慢指针,快指针一次移动2个
	fast, slow := l.head, l.head

	//让快指针到达或超过末尾
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

	}
	return slow
}

/* 2.获得链表倒数第n个节点 */
func (l *link) getRev(n int) *node {
	if n > l.length || n < 1 {
		panic("输入不合法")
	}

	//使快指针先移动n个位置,后慢指针和快指针同时移动
	fast, slow := l.head, l.head
	for i := 1; i <= n; i++ {
		fast = fast.next
	}
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	//此时慢指针就在目标前一个
	return slow.next
}

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

/* 4.是否有环 利用go的map */
func (l *link) isRing() bool {
	if l.length < 2 {
		return false
	}
	tem := make(map[*node]int)
	pre := l.head
	for pre.next != nil { //遍历所有节点
		//将每一个节点作为map的key放入
		if _, ok := tem[pre]; ok { //头节点也要验证
			return true //此处可返回入口节点
		} else {
			tem[pre] = 1
		}
		pre = pre.next
	}
	return false
}

func (l *link) isRing2() bool {
	//快慢指针
	if l.length < 2 {
		return false
	}
	fast, slow := l.head, l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			return true
		}
	}
	return false
}

/* 5.有环的话 求环的大小 */
//获得快慢指针相遇的节点
func getNode(l *link) *node {
	fast, slow := l.head, l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return slow
		}
	}
	return nil
}
func (l *link) getRingSize() int {
	if !l.isRing2() {
		return -1
	}
	//有环 求大小;快指针第一次和慢指针相遇后,第二次相遇时,刚好相隔一个环的大小
	//获得第一次相遇的节点
	n := getNode(l)
	len := 1
	temp := n.next
	for temp != n { //再相等时说明跑完一圈,计数器就是长度
		len++
		temp = temp.next
	}
	return len
}

/* 6.环的入口位置节点 */
//先获取环的长度n,之后让快指针走过n,之后和慢指针同时移动,当两者相遇处,即是入口节点
//为以上内容综合运用,判断是否是环,获取环的大小,之后在用快慢指针获取入口(类似获取倒数第几个节点)
func (l *link) ringEntry() *node {
	n := l.getRingSize()
	if n < 0 {
		return nil
	}
	ft, sl := l.head, l.head

	for i := 1; i <= n; i++ {
		ft = ft.next
	}
	for sl != ft {
		sl = sl.next
		ft = ft.next
	}
	return sl

}

/* 7.给定单向链表的头指针和一个节点指针(给定其某个节点地址)，定义一个函数在O（1）时间内删除该节点。 */
//直接删除某个节点一定是O(n),因为我们要先找到其前一个节点,必须遍历,但获取到待删除节点下一个节点就很容易
func (l *link) del(n *node) {
	//将待删节点和其下一个节点的value交换,实际可以达到删除的同样目的
	//待删除节点是尾节点时,复杂度为O(n),但平均时间复杂度仍为O(1)
}

/* 8.两个链表的第一个公共节点(交叉处) */

/* 9.合并两个有序的链表 */

/* 10.链表复制 */

/* 11.链表反转 */
//四种方法:http://c.biancheng.net/view/8105.html
func (l *link) reverse() *link {
	//创建新链表,头插法 将旧链表的节点迁入
	newList := &link{head: &node{value: "head node", next: nil}, length: 0}

	pre := l.head
	for pre.next != nil { //直到旧链表没有有效节点
		temp := pre.next
		pre.next = pre.next.next //将旧链表第一个节点取出

		temp.next = newList.head.next //头插法接入新链表
		newList.head.next = temp
		newList.length++
	}

	return newList
}
func main() {
	//创建一个链表,用于测试
	l := GetLinkedList().add(10)
	l.rangeList() //正反向遍历
	l.rangeList2()

	l.addNode(10000000)
	l.addNode(1021)
	l.removeNode(1) //删除第一个元素
	fmt.Println("---------------------增加2个删除1个后----------------------")
	l.rangeList() //检测顶部插入
	fmt.Println("链表长度:", l.getLen())
	fmt.Println("中间节点值:", l.getCenterNode().value) //获取链表的中间节点
	fmt.Println("倒数第1个值是:", l.getRev(1).value)
	fmt.Println("倒数第2个值是:", l.getRev(2).value)
	fmt.Println("倒数第3个值是:", l.getRev(3).value)
	fmt.Println("倒数第10个值是:", l.getRev(10).value)

	fmt.Println("------------------------连成环之后------------------------")
	//l.makeRing()                     //连成一个环
	fmt.Println("是否是环:", l.isRing()) //测试是否是环
	fmt.Println("是否是环:", l.isRing2())
	fmt.Println("环的大小:", l.getRingSize()) //因为头尾相连包含了头节点
	res := l.ringEntry()
	if res != nil {
		fmt.Println("环的入口节点值:", res.value)
	} else {
		fmt.Println("其不是一个有效的环!")
	}
	fmt.Println("------------------------反转链表------------------------")
	l = l.reverse()
	l.rangeList()
	fmt.Println("链表长度:", l.getLen())

}
