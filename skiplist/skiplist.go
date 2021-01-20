package skiplist

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

const (
	MAX_LEVEL = 16
	SKIPLISTP = 25
)

type Node struct {
	// 当前节点的值
	Value interface{}
	// 用于排序的值,如果可以通过value排序则可以不需要该值
	Score int
	// 当前节点所在的层数
	Level int
	// 当前节点每一层对应的后面已个节点的值
	Forwards []*Node
}

func NewSkipListNode(v interface{}, score, level int) *Node {
	return &Node{
		Value:    v,
		Score:    score,
		Level:    level,
		Forwards: make([]*Node, level),
	}
}

type SkipList struct {
	// 调表的带头节点
	head *Node
	// 当前跳表的层高
	level int
	// 当前跳表中的结点数
	length int
}

func NewSkipList() *SkipList {
	return &SkipList{
		head:   NewSkipListNode(math.MinInt64, math.MinInt64, MAX_LEVEL),
		level:  1,
		length: 0,
	}
}

func (s *SkipList) Level() int {
	return s.level
}

func (s *SkipList) Length() int {
	return s.length
}

func (s *SkipList) RandomLevel() int {
	level := 1
	rand.Seed(time.Now().UnixNano())
	for rand.Intn(100) < SKIPLISTP {
		level++
	}

	if level > MAX_LEVEL {
		return MAX_LEVEL
	}
	return level
}

func (s *SkipList) Insert(value interface{}, score int) int {
	// 需要找到每一层的链表中最后一个小于score的值
	cur := s.head
	updates := [MAX_LEVEL]*Node{}

	// 此处必须从最高层开始找
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.Forwards[i] != nil {
			// 如果找到完全相同的结点则不插入
			if cur.Forwards[i].Value == value && cur.Forwards[i].Score == score {
				return -1
			}
			// 找到第一个大于当前要插入的结点的分值结束
			if cur.Forwards[i].Score > score {
				updates[i] = cur
				break
			}
			// 将当前层最优一个节点加入
			cur = cur.Forwards[i]
		}
		updates[i] = cur
	}

	level := s.RandomLevel()
	currentNode := NewSkipListNode(value, score, level)
	for i := level - 1; i >= 0; i-- {
		nextNode := updates[i].Forwards[i]
		updates[i].Forwards[i] = currentNode
		currentNode.Forwards[i] = nextNode
	}

	if level > s.level {
		s.level = level
	}
	s.length += 1
	return 0
}

func (s *SkipList) Find(value interface{}, score int) *Node {
	cur := s.head
	for i := cur.Level - 1; i >= 0; i-- {
		for cur.Forwards[i] != nil {
			if cur.Forwards[i].Value == value && cur.Forwards[i].Score == score {
				return cur.Forwards[i]
			}

			if cur.Forwards[i].Score > score {
				break
			}

			cur = cur.Forwards[i]
		}
	}
	return nil
}

func (s *SkipList) Delete(value interface{}, score int) int {
	if value == nil {
		return -1
	}

	// 这里需要找到每一层中要删除的结点的前一个节点
	cur := s.head
	updates := [MAX_LEVEL]*Node{}
	exists := false
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		updates[i] = s.head
		for cur.Forwards[i] != nil {
			if cur.Forwards[i].Value == value && cur.Forwards[i].Score == score {
				exists = true
				updates[i] = cur
				break
			}
			if cur.Forwards[i].Score > score {
				break
			}
			cur = cur.Forwards[i]
		}
	}

	if !exists {
		return -1
	}

	// 获取当前节点
	currentNode := updates[0].Forwards[0]
	for i := currentNode.Level - 1; i >= 0; i-- {
		// 如果当前层需要更新的是头结点，且当前要删除的结点在该层的下一个节点为空
		if updates[i] == s.head && currentNode.Forwards[i] == nil {
			s.level -= 1
		}

		if updates[i].Forwards[i] != nil {
			updates[i].Forwards[i] = updates[i].Forwards[i].Forwards[i]
		}
	}

	s.length--
	return 0
}

func (s *SkipList) Print() {
	for i := s.level - 1; i >= 0; i-- {
		cur := s.head
		var values []string
		for cur.Forwards[i] != nil {
			values = append(values, fmt.Sprintf("%v:%v", cur.Forwards[i].Value, cur.Forwards[i].Score))
			cur = cur.Forwards[i]
		}
		fmt.Println(strings.Join(values, "->"))
	}
}
