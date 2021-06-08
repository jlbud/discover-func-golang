package main

import (
	"strings"
	"testing"
	"unicode"
)

// 用双链表
// 去掉汉字中间和两边的空格
func TestHan(t *testing.T) {
	str := "Hello    world 世 界 fdsa 我们 fdfdsa 我们"
	rstr := []rune(str)

	list := DoubleLinkedList{}
	index := make([]int, 0)
	for k, v := range rstr {
		if unicode.Is(unicode.Han, v) {
			index = append(index, k)
		}
		node := &Node{
			status: 0,
			data:   string(v),
		}
		list.Append(node)
	}

	for _, v := range index {
		for i := v - 1; i > 0; i-- {
			if list.Get(i).data == " " {
				list.Get(i).status = 1
			} else {
				break
			}
		}
		for i := v + 1; i < list.Size(); i++ {
			if list.Get(i).data == " " {
				list.Get(i).status = 1
			} else {
				break
			}
		}
	}
	t.Log(list.ToString())
}

/**
  双向链表
*/
type DoubleLinkedList struct {
	//链表头节点
	head *Node
	//链表尾部节点
	tail *Node
	//长度
	size int
}

/**
  遍历器接口
*/
type Interater interface {
	/**
	  是否还有下一个
	*/
	hasNext() bool
	/**
	  下一个节点
	*/
	next() *Node
}

/**
  遍历器
*/
type ListInterater struct {
	list        *DoubleLinkedList
	currentNode *Node
}

/**
  遍历器是否还有下一个节点
*/
func (inter *ListInterater) hasNext() bool {
	if inter.currentNode == nil && inter.list.head != nil {
		return true
	} else if inter.currentNode == nil && inter.list.head == nil {
		return false
	} else {
		return inter.currentNode.next != nil
	}

}

/**
  遍历器获取下一个节点
*/
func (inter *ListInterater) next() *Node {
	if inter.currentNode == nil && inter.list.head != nil {
		inter.currentNode = inter.list.head
	} else if inter.currentNode == nil && inter.list.head == nil {
		return nil
	} else {
		inter.currentNode = inter.currentNode.next
	}
	return inter.currentNode
}

/**
  节点
*/
type Node struct {
	data   interface{}
	status int
	prev   *Node
	next   *Node
}

func (list *DoubleLinkedList) GetHead() *Node {
	return list.head
}

func (list *DoubleLinkedList) GetTail() *Node {
	return list.tail
}

func (list *DoubleLinkedList) Size() int {
	return list.size
}

func (list *DoubleLinkedList) initFromNode(node *Node) {
	list.head = node
	list.tail = node
	list.size = 1
}

func (list *DoubleLinkedList) addHead(node *Node) {
	oldHead := list.head
	oldHead.prev = node
	node.next = oldHead
	node.prev = nil
	list.head = node
	list.size++
}

/**
  获取链表的遍历器
*/
func (list *DoubleLinkedList) Iterater() *ListInterater {
	inter := ListInterater{
		list:        list,
		currentNode: nil,
	}
	return &inter
}

/**
  从头部添加节点
*/
func (list *DoubleLinkedList) AddFromHead(node *Node) {
	if list.head == nil && list.tail == nil {
		list.initFromNode(node)
	} else {
		list.addHead(node)
	}
}

/**
  从尾部添加节点
*/
func (list *DoubleLinkedList) addTail(node *Node) {
	oldTail := list.tail
	oldTail.next = node
	node.prev = oldTail
	node.next = nil
	list.tail = node
	list.size++
}

/**
  删除指定位置节点
*/
func (list *DoubleLinkedList) Remove(index int) {
	node := list.Get(index)
	if node == nil {
		return
	}
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	list.size--
}

/**
  从后面追加节点
*/
func (list *DoubleLinkedList) Append(node *Node) {
	if list.head == nil && list.tail == nil {
		list.initFromNode(node)
	} else {
		list.addTail(node)
	}
}

/**
  获取指定位置节点
*/
func (list *DoubleLinkedList) Get(index int) *Node {
	next := list.head
	for i := 0; i < index; i++ {
		if next.next != nil {
			next = next.next
		} else {
			return nil
		}
	}
	return next
}

/**
  打印链表
*/
func (list *DoubleLinkedList) ToString() string {
	str := strings.Builder{}
	//str.WriteString("[")
	inter := list.Iterater()
	for inter.hasNext() {
		node := inter.next()
		if node.status == 0 {
			str.WriteString(node.data.(string))
		}
		//if inter.hasNext() {
		//	str.WriteString(",")
		//}
	}
	//str.WriteString("]")
	return str.String()
}
