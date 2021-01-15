package main

import (
	"fmt"
)

type St interface {
	Put(k STComparableKey, value interface{})
	Get(k STComparableKey) interface{}
	Del(k STComparableKey) bool
	Contains(k STComparableKey) bool
	IsEmpty() bool
	Keys() STIterator
	Size() int
}

// key
//type Comparable interface {
//	Compare(key Comparable) int
//}

type STIterator struct {
	SKeys []STComparableKey
}

type STComparableKey struct {
	K int
}

func (k *STComparableKey) Compare(key STComparableKey) int {
	//c := Comparable(k)
	//return c.Compare(key)
	return k.K - STComparableKey(key).K
}

// inner data
type StNode struct {
	Next  *StNode
	Key   STComparableKey
	Value interface{}
}

// outer data
type SequentialSearchST struct {
	First *StNode
	N     int
}

// 第一个版本key只保证对比，不保证数据有序
func NewSequentialSearchST() *SequentialSearchST {
	st := new(SequentialSearchST)
	st.First = nil
	st.N = 0
	return st
}

func (st SequentialSearchST) Size() int {
	return st.N
}

func (st SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

func (st SequentialSearchST) Contains(k STComparableKey) bool {
	return nil != st.Get(k)
}

func (st SequentialSearchST) Put(key STComparableKey, value interface{}) {

	// 暂时先不做校验，后续补充上
	//if(nil == key){
	//}

	for node := st.First; nil != node; node = node.Next {
		if node.Key.Compare(key) == 0 {
			// 如果key命中了则更新value
			node.Value = value
			return
		}
	}

	// 每次更新数据都把新数据放到头部,
	st.First = &StNode{Key: key, Value: value, Next: st.First}
	st.N++

}

func (st SequentialSearchST) Get(k STComparableKey) interface{} {

	if st.First == nil {
		return nil
	}

	for node := st.First; nil != node; node = node.Next {
		if node.Key.Compare(k) == 0 {
			return node.Value
		}
	}

	return nil
}

func (st SequentialSearchST) Keys() STIterator {
	iterator := STIterator{}
	iterator.SKeys = make([]STComparableKey, 100, 100)

	if st.First == nil {
		return iterator
	}

	for s := st.First; s != nil; s = s.Next {
		keys := append(iterator.SKeys, s.Key)
		iterator.SKeys = keys
	}
	return iterator
}

//  -------------------上面比较简单，为啥我老想用双向链表

//  写过一次了还想着用pre?
func (st SequentialSearchST) Del(k STComparableKey) {
	st.First = st.DelWithNode(st.First, k)
}

//  重点
//  写过一次了还想着用pre? 一会debug
func (st SequentialSearchST) DelWithNode(s *StNode, k STComparableKey) *StNode {
	if s == nil {
		return nil
	}

	if s.Key.Compare(k) == 0 {
		st.N--
		return s.Next
	}
	s.Next = st.DelWithNode(s.Next, k)
	return s
}

// test
func main() {

	st := NewSequentialSearchST()
	fmt.Println(st.Size())
	fmt.Println(st.IsEmpty())

	st.Put(STComparableKey{1}, 1)
	st.Put(STComparableKey{2}, 2)
	fmt.Println(st.Keys().SKeys)

	//strings.Compare("a","a")

}
