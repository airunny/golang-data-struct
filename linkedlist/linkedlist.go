package linkedlist

import (
	"fmt"
	"sync"
)

type Node struct {
	Value float64
	Next *Node
}

func (n* Node) String()string{
	if n == nil {
		return ""
	}

	ret := ""
	current := n
	for   {
		if current  == nil{
			ret += fmt.Sprint("nil")
			return ret
		}

		ret += fmt.Sprintf("%v->",current.Value)
		current = current.Next
	}
}

func NewNode(value float64)*Node{
	return &Node{
		Value:value,
	}
}

func NewNodeWithNext(value float64,next * Node)*Node  {
	return &Node{
		Value:value,
		Next:next,
	}
}

// 单链表反转
// input 1->2->3->4
// output 4->3->2->1
/**
 * 思路：pre：指向已经反转好的节点的头结点；current：指向下一个待转换的结点
 *
 * pre = nil, cur = 1->2->3->4->5->nil
 * pre = 1->nil, cur = 2->3->4->5->nil
 * pre = 2->1->nil, cur = 3->4->5->nil
 * pre = 3->2->1->nil, cur = 4->5->nil
 * ......
 * pre = 5->4->3->2->1->nil, cur = nil
 */
func Traverse(head * Node)*Node  {
	var  pre ,current *Node = nil,head
	for{
		if current == nil{
			return pre
		}
		current.Next,pre,current =pre,current,current.Next
	}
}

// 单链表相邻两个节点两两反转
// input 1->2->3->4->5
// output 2->1->4->3->5
/**
 * 思路：如果小于等于一个节点直接返回不需要反转
 *      这里必须要标记当前转换的两个节点的前一个节点，所以在刚开始时必须要构造这么一个节点
 *
 * pre = 2, cur = 0->1->2->3->4->5
 * pre = 2, cur = 0->2->1->3->4->5
 * pre = 2, cur = 0->2->1->4->3->5
 */
func Traverse2(head * Node)*Node{
	if head == nil || head.Next == nil {
		return head
	}

	pre := head.Next
	cur := NewNodeWithNext(0,head)
	for  {
		if cur.Next == nil || cur.Next.Next == nil{
			return pre
		}

		a := cur.Next
		b := cur.Next.Next
		cur.Next,a.Next,b.Next=b, b.Next,a
		cur = a
	}
}

// 检测链表是否有环（快慢指针）
func HasCircle(head * Node)bool{
	if head == nil{
		return false
	}

	slow,fast := head,head
	for   {
		if fast.Next == nil || fast.Next.Next == nil || slow.Next == nil {
			return false
		}

		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast{
			return true
		}
	}
}

// 获取链表的中间节点
func GetMiddleNode(head * Node)* Node{
	if head == nil || head.Next == nil{
		return head
	}

	slow,fast := head,head
	for   {
		if fast.Next != nil && fast.Next.Next != nil && slow.Next != nil {
			slow = slow.Next
			fast=fast.Next.Next
			continue
		}

		return slow
	}
}

// 删除倒数第n个结点，如果倒数n个结点为头结点则不可删除
func RDelNode(head* Node,n int) * Node {
	slow,fast := head,head

	// 1 2 3 4 5
	for ; n > 0 ; n --   {
		if fast.Next == nil {
			break
		}
		fast = fast.Next
	}

	// 如果这里 n == 1 则表示删除头结点，目前可以考虑不删除头结点
	if n > 0 {
		return nil
	}

	for  {
		if fast.Next != nil && slow.Next != nil {
			fast = fast.Next
			slow = slow.Next
			continue
		}
		break
	}

	if slow != nil && slow.Next != nil{
		ret := slow.Next
		slow.Next = slow.Next.Next
		return ret
	}
	return nil
}

// 合并两个有序链表
/**
 * 思路：因为是有序的链表，所以有两种方式可以做
 * 1、重新开启一个链表，然后从两个待合并的链表的头中拿下来一个比较小的，放入新的链表中直到两个链表都为空
 * 2、上一个方式理解起来比较简单，但是会额外浪费m+n的空间；所以第二种方式就是一个开头较大的链表插入到开头较小的链表中即可
 */
func Merge(head1,head2 * Node)*Node{
	if head1 == nil{
		return head2
	}

	if head2 == nil{
		return head1
	}

	head,insert := head1,head2
	if head2.Value < head1.Value {
		head,insert = head2,head1
	}

	cur := head
	for {
		if insert == nil{
			return head
		}

		if cur.Next == nil {
			cur.Next = insert
			return head
		}

		for {
			if  insert.Value > cur.Value && cur.Next != nil && insert.Value < cur.Next.Value  {
				tmp := insert.Next
				insert.Next = cur.Next
				cur.Next = insert

				cur = cur.Next
				insert =tmp
				break
			}

			if cur.Next == nil {
				break
			}
			cur = cur.Next
		}
	}
}

// ------------------------**链表实现LRU淘汰算法（Least Recently Used）**------------------------------------
/**
 *  思路：
 * 1、如果当前数据在链表中存在则将当前数据提到链表头部
 * 2、如果在当前链表中不存在
 *  2.1、当前链表是否已满，当前数据替换链表中的最后一个节点
 *  2.2、当前链表没有满，将当前数据插入到链表的头部
 */
type LeastRecentlyUsed struct {
	Capacity uint64 // 容量
	Number uint64 // 当前链表数量
	Head  * Node // 链表头节点
	mu sync.RWMutex
}

func NewLeastRecentlyUsed(capacity uint64)*LeastRecentlyUsed{
	return &LeastRecentlyUsed{
		Capacity:capacity,
	}
}

// 返回要查找的结点的前一个节点跟自己的节点，如果pre为空则要查找的结点就是头结点
func (l * LeastRecentlyUsed)Find(value interface{})(pre,cur *Node,exist bool){
	l.mu.RLock()
	defer l.mu.RUnlock()

	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		cur = l.Head
		exist = true
		return
	}

	pre = l.Head
	for  {
		if pre.Next == nil {
			return
		}

		if pre.Next.Value == value {
			cur = pre.Next
			exist = true
			return
		}

		pre = pre.Next
	}
}

func (l * LeastRecentlyUsed)Use(value float64)  {
	pre,cur ,ok := l.Find(value)
	l.mu.Lock()

	// 如果当前节点已经存在则将当前节点放在第一个节点
	if ok && pre != nil{
		pre.Next= cur.Next
		cur.Next = l.Head
		l.Head = cur
		l.mu.Unlock()
		return
	}

	// 如果当前的缓存容量已经满了，则将当前节点覆盖最后一个节点
	if l.Capacity <= l.Number {
		if l.Capacity <= 1 {
			l.Head = NewNode(value)
			l.mu.Unlock()
			return
		}

		pre := l.Head
		for  {
			if pre.Next.Next != nil {
				pre = pre.Next
				continue
			}

			pre.Next = NewNode(value)
			l.mu.Unlock()
			return
		}

	}else{
		l.Number ++
		newNode := NewNodeWithNext(value,l.Head)
		l.Head = newNode
		l.mu.Unlock()
		return
	}
}


