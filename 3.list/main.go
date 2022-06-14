package main

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

在顺序结构中,每个数据元素只需存放数据元素本身即可,现在在链式结构中,除了元素信息以外,还要存他后继元素的地址,我们把存数据元素信息的域称为数据域,存储直接后继位置的域称为指针域
指针域中的信息称作指针或链,这两部分组成数据元素的存储映像,称为节点(node)
n个节点链接形成一个链表,因为每一个节点中只包含一个指针域,所以叫单链表

链表的第一个节点的存储位置叫做头指针,整个链表的存取就必须从头指针开始,之后的每个节点,其实就是上一个后继指针指向的位置,最后一个节点的后继指针为nil

有时候为了方便操作,会在单链表第一个节点前在附设一个节点,称为头节点,头节点的数据域可以不存任何信息(但可以存链表长度等附加信息),头节点的指针域存着指向第一个节点的指针

3.1头指针和头节点的区别
头指针是指向链表第一个节点的指针,头指针是链表的必要元素(可以理解为有了头指针才能知道链表从哪里开始的),若链表有头节点,则是指向头节点的指针,头指针仅仅是一个指针即可,他也代表着链表的名字

而头节点是为了操作的统一方便设立的,放在第一元素节点之前,可用于存链表长度等信息,有了头节点,操作真正第一个节点时就和操作其他节点一样,头节点是非必须的,如果线性表为空,头节点的指针域值为nil

3.2单链表的读取
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

*/

func main() {

}

//TODO:实现单链表增删改查
//https://www.cnblogs.com/skzxc/p/11453679.html
