package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
一.基本
线性表定义:0个或多个数据元素的有限序列
元素之间是有顺序的,并且强调是有限的

假想我们在使用线性表时,所需要用到的基本操作(如增删改查),就可以得到线性表的抽象数据类型表示:
ADT List
Data
	定义...描述...
Operation
	ListEmpty(L)
        初始条件：线性表已存在
        操作结果：若线性表L为空表，则返回TRUE，否则返回FALSE
	 GetElem(L, i, &e)
        初始条件：线性表已存在（1≥i≥ListLenght(L)）
        操作结果：用e返回线性表L中第i个数据元素的值
		...
endADT

以上的只是线性表的最基本的操作,对于实际问题可能还需更复杂的操作,旧可以使用基本操作的组合来实现


二.线性表的顺序存储结构
指的是用一段地址连续的存储单元一次存储线性表的数据元素,比如go语言中的数组,线性表长度要小于等于数组的长度

存储器中每个存储单元都有自己的编号,这个编号被称为地址,当获取到起始位置时,后面的编号都可以计算获得,而计算
的成本都是一致的,不会因为线性表的大小而改变,因此我们对线性表每个位置的存入或取出,对于计算机来说都是相等的时间,
存储性能为O(1)

2.1线性表的操作
查:
	根据索引返会对应的值即可
增:
	思路:
		插入位置不合理抛出异常
		线性表长度大于等于数组长度,扩容或者抛异常
		从最后一个元素向前遍历,到第i个位置,分别将他们全部向后移动一个位置
		将要加入的元素放入i处
		表长加1

2.2线性表的缺点:
	增删会造成大量元素位移



三.线性表的链式存储结构
为什么顺序存储时增加删除会造成大量数据位移?因为内存分布是紧挨着的,地址是顺序的,删除后留下的空白需要填充

3.1.单链表
在顺序结构中,每个数据元素只需存放数据元素本身即可,现在在链式结构中,除了元素信息以外,还要存他后继元素的地址,我们把存数据元素信息的域称为数据域,存储直接后继位置的域称为指针域
指针域中的信息称作指针或链,这两部分组成数据元素的存储映像,称为节点(node)
n个节点链接形成一个链表,因为每一个节点中只包含一个指针域,所以叫单链表

链表的第一个节点的存储位置叫做头指针,整个链表的存取就必须从头指针开始,之后的每个节点,其实就是上一个后继指针指向的位置,最后一个节点的后继指针为nil

有时候为了方便操作,会在单链表第一个节点前在附设一个节点,称为头节点,头节点的数据域可以不存任何信息(但可以存链表长度等附加信息),头节点的指针域存着指向第一个节点的指针

3.1.1头指针和头节点的区别
头指针是指向链表第一个节点的指针,头指针是链表的必要元素(可以理解为有了头指针才能知道链表从哪里开始的),若链表有头节点,则是指向头节点的指针,头指针仅仅是一个指针即可,他也代表着链表的名字

而头节点是为了操作的统一方便设立的,放在第一元素节点之前,可用于存链表长度等信息,有了头节点,操作真正第一个节点时就和操作其他节点一样,头节点是非必须的,如果线性表为空,头节点的指针域值为nil

3.1.2单链表的读取
在顺序结构中,得到任意一个元素的存储位置是非常容易的,但是在链表中
由于某个元素的位置一开始是没法知道的,就得从头开始遍历寻找,要相对麻烦一些

思路:
	1.声明一个节点p指向链表得第一个节点,初始化j从1开始
	2.当j<i时,就依次遍历链表,让p得指针向后移动,不断指向下一节点,并同时j累加1
	3.若直到p指向nil,说明第i个元素不存在
	4.否则查找成功,返回p得数据

时间复杂度取决于i得位置,最坏为O(n),在查找方面链表不占优势
单链表中没有定义表长,所以不能事先知道要循环多少次,所以无法使用for循环来控制,其主要核心思想就是工作指针后移,很多算法常用得技术

3.3插入和删除
链表得插入删除就是其主要优势,只需让待插入得节点得指针域存储前一节点原指向的后继节点,并使前一节点的指针域存待插入节点的地址即可,其他所有节点都不需要修改(顺序不能颠倒,一定要先将后继节点存入待插入元素,否则将丢失其地址)

而删除某个节点则是将其上一个节点的next指向待删除的next(直接绕过待删除节点,待删除节点next置为nil即可)

对于单链表的增删查,时间复杂度都是o(n),单一操作和顺序结构没有太大区别,但是如果要一次插入大量数据,对于顺序结构来说,每一次插入都需要后移n-i个元素,每次都是O(n),而单链表仅在第一次找i的位置时为O(n),一旦找到后续插入仅仅是简单的通过赋值来改变指针而已,为O(1)复杂度;显然,插入删除数据操作越频繁,链表越有优势,而顺序结构适用于增删不频繁但查频繁的场景


3.2.单循环链表
对于单链表,每个节点只存了向后的指针,只能单向遍历,对于后继节点,其无法知道前驱节点的情况

将单链表中尾节点的指针由nil改为指向头节点,就能够使单链表形成一个环,就成了单循环链表!
循环链表解决了从其中任意节点开始无法访问到全部节点的问题

为了使空链表和非空链表操作一致,通常也会设置一个头节点;当有了头节点之后,访问第一个元素时间复杂度为
O(1),访问最后一个元素确实O(n),如何优化??
解决方法是不使用头指针指向头节点,而是单独创建一个尾指针指向尾节点(rear),这样头节点就是rear.next,
第一个节点就是rear.next.next,这样就使得获取尾节点和起始处都为O(1)

有了尾指针使得合并两个循环链表也很容易
假设循环链表A,有尾节点rearA,循环链表B有尾节点rearB
即A的头节点pA为rearA.Next,B的头节点pB为rear.Next
将A的尾指针由原指向pA改为指向B的第一个节点,即pB.next,再将rearB指向pA(合并后保留一个头节点)

TODO:循环链表经典题:如何判断一个链表是否有环?


3.3.双向链表
在单链表的每个节点中,再设置一个指向其前驱节点的指针域,相当于双向链表中,每一个节点同时有2个指针域,
分别保存其前继节点和后继节点
双向链表也可以构成循环链表,双向链表的很多操作和单链表都是相同的,多的功能就是能够反向遍历查找等,但是
需要付出一些代价:在插入删除时,需要修改两个指针变量,要尤其注意顺序!!

假设待储存节点s,要将其插入到节点p和p.next之间:
	1.首先将p赋值给s的前驱,s.pre=p
	2.把p.next赋值给s的后继,s.next=p.next
	3.把s赋值给p.next的前驱
	4.把s赋值给p的后继即p.next=s
先搞定s的前驱后继,再搞定后继节点的前驱,最后解决前节点的后继

删除的话,使待删除节点的前驱后继跳过 不再指向它即可
s.pre.next=s.next
s.next.pre=s.pre

双向链表中带你在于增删操作,由于双向的特性可以有效提高算法的时间性能



线性表的两种存储结构,是后面其他数据结构的基础!


*/

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
