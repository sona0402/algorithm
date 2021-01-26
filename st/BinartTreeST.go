package main

import (
	"errors"
	"fmt"
	"strings"
)

type BinaryTreeInnerNode struct {
	Key   string
	Value interface{}
	Left  *BinaryTreeInnerNode
	Right *BinaryTreeInnerNode
	N     int
}

type BinaryTreeSt struct {
	Root *BinaryTreeInnerNode
}

func NewBinaryTreeSt() *BinaryTreeSt {
	b := new(BinaryTreeSt)
	return b
}

func (b *BinaryTreeSt) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BinaryTreeSt) Size() int {
	return b.Root.N
}

func (b *BinaryTreeSt) CreateNewNode(k string, value interface{}) *BinaryTreeInnerNode {
	newNode := new(BinaryTreeInnerNode)
	newNode.Key = k
	newNode.Value = value
	newNode.N = 1
	return newNode
}

// 这种方式更加优秀
func (b *BinaryTreeSt) Put(k string, value interface{}) {
	if len(k) == 0 {
		_ = errors.New("k must not null")
	}
	if nil == value {
		_ = errors.New("value must not null")
	}
	// 从跟节点找
	b.Root = b.PutWithNode(b.Root, k, value)
}

func (b *BinaryTreeSt) PutWithNode(iota *BinaryTreeInnerNode, k string, value interface{}) *BinaryTreeInnerNode {
	if nil == iota {
		return b.CreateNewNode(k, value)
	}
	compare := strings.Compare(k, iota.Key)
	if compare < 0 {
		iota.Left = b.PutWithNode(iota.Left, k, value)
	} else if compare > 0 {
		iota.Right = b.PutWithNode(iota.Right, k, value)
	} else {
		iota.Value = value
	}
	iota.N = b.SizeWithNode(iota.Right) + b.SizeWithNode(iota.Left) + 1
	return iota
}

func (b *BinaryTreeSt) SizeWithNode(node *BinaryTreeInnerNode) int {
	if nil == node {
		return 0
	}
	return node.N
}

//func (b *BinaryTreeSt) Put(k string, value interface{}) {
//	if len(k) == 0 {
//		_ = errors.New("k must not null")
//	}
//	if nil == value {
//		_ = errors.New("value must not null")
//	}
//	if b.Root == nil {
//		b.Root = b.CreateNewNode(k, value)
//		return
//	}
//	b.PutWithNode(b.Root, k, value)
//}
//
//func (b *BinaryTreeSt) CreateNewNode(k string, value interface{}) *BinaryTreeInnerNode {
//	newNode := new(BinaryTreeInnerNode)
//	newNode.Key = k
//	newNode.Value = value
//	newNode.N = 1
//	return newNode
//}
//
//func (b *BinaryTreeSt) PutWithNode(iota *BinaryTreeInnerNode, k string, value interface{}) {
//	compare := strings.Compare(k, iota.Key)
//	if compare < 0 {
//		if iota.Left == nil {
//			iota.Left = b.CreateNewNode(k, value)
//			return
//		} else {
//			b.PutWithNode(iota.Left, k, value)
//		}
//	} else if compare > 0 {
//		if iota.Right == nil {
//			iota.Right = b.CreateNewNode(k, value)
//			return
//		} else {
//			b.PutWithNode(iota.Right, k, value)
//		}
//	} else {
//		iota.Value = value
//		return
//	}
//}

func (b *BinaryTreeSt) Get(k string) interface{} {
	if b.Root == nil {
		return nil
	}
	return b.GetByNode(b.Root, k)
}

func (b *BinaryTreeSt) GetByNode(iota *BinaryTreeInnerNode, k string) interface{} {
	// 下面这段代码不如算法4中的代码优秀
	//compare := strings.Compare(k, iota.Key)
	//if compare < 0 {
	//	if iota.Left == nil {
	//		return nil
	//	} else {
	//		return b.GetByNode(iota.Left, k)
	//	}
	//} else if compare > 0 {
	//	if iota.Right == nil {
	//		return nil
	//	} else {
	//		return b.GetByNode(iota.Right, k)
	//	}
	//}
	//return iota.Value
	if nil == iota {
		return nil
	}
	compare := strings.Compare(k, iota.Key)
	if compare < 0 {
		return b.GetByNode(iota.Left, k)
	} else if compare > 0 {
		return b.GetByNode(iota.Right, k)
	}
	return iota.Value
}

func (b *BinaryTreeSt) Contains(k string) bool {
	return b.GetByNode(b.Root, k) != nil
}

// 好难
func (b *BinaryTreeSt) Del(k string) {
	b.Root = b.DelWithNode(b.Root, k)
}

func (b *BinaryTreeSt) Max(x *BinaryTreeInnerNode) *BinaryTreeInnerNode {
	if x.Right == nil {
		return x
	} else {
		return b.Max(x.Right)
	}
}

func (b *BinaryTreeSt) Min(x *BinaryTreeInnerNode) *BinaryTreeInnerNode {
	if x.Left == nil {
		return x
	} else {
		return b.Min(x.Left)
	}
}

func (b *BinaryTreeSt) DelWithNode(x *BinaryTreeInnerNode, k string) *BinaryTreeInnerNode {
	if nil == x {
		return nil
	}
	key := x.Key
	compare := strings.Compare(k, key)
	if compare < 0 {
		b.DelWithNode(x.Left, k)
	} else if compare > 0 {
		b.DelWithNode(x.Right, k)
	} else {
		// 这里是最复杂的
		if x.Left == nil {
			return x.Right
		}
		if x.Right == nil {
			return x.Left
		}
		t := x
		// 获取到右子节点的最小值
		x := b.Min(t.Right)
		// 将自己节点与原右子树断开，并获取到右子节点
		x.Right = b.deleteMin(t.Left)
		// 左子节点直接链接过来
		x.Left = t.Left
	}
	// 修正此节点，将新的孩子节点给父亲节点
	x.N = b.SizeWithNode(x.Right) + b.SizeWithNode(x.Left) + 1
	return x
}

func (b *BinaryTreeSt) deleteMin(x *BinaryTreeInnerNode) *BinaryTreeInnerNode {
	if x.Left == nil {
		return x.Right
	}
	x.Left = b.deleteMin(x.Left)
	x.N = b.SizeWithNode(x.Left) + b.SizeWithNode(x.Right) + 1
	return x
}

// 左中右处理序处理
//func (BinaryTreeSt) Keys() *STIterator {
//	//panic("implement me")
//	return new(STIterator)
//}

func main() {
	st := NewBinaryTreeSt()
	st.Put("6", 9)
	st.Put("7", 7)
	st.Put("4", 4)
	st.Put("3", 3)
	st.Put("8", 8)
	st.Put("10", 10)
	fmt.Println(st.Contains("11"))
	fmt.Println(st.Root.N)
}
