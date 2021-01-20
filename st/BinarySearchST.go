package main

import (
	"fmt"
	"strings"
)

//
// select min max  floor ceiling
//
// binarySearchST
type BinarySearchST struct {
	Keys  []string
	Value []interface{}
	N     int
}

func NewBinarySearchST(n int) *BinarySearchST {
	r := new(BinarySearchST)
	r.N = 0
	r.Keys = make([]string, n, n)
	r.Value = make([]interface{}, n, n)
	return r
}

func (st *BinarySearchST) Size() int {
	return st.N
}

func (st *BinarySearchST) IsEmpty() bool {
	return st.Size() == 0
}

func (st *BinarySearchST) Get(k string) interface{} {
	if st.N == 0 {
		return nil
	}
	// 这里返回的值会是0-n中的一个
	rankIndex := st.rank(k)
	if rankIndex < st.N &&
		strings.Compare(st.Keys[rankIndex], k) == 0 {
		return st.Value[rankIndex]
	}
	return nil
}

func (st *BinarySearchST) resize(len int) {
	newKeys := make([]string, len, len)
	newValues := make([]interface{}, len, len)
	keys := st.Keys
	value := st.Value
	for i := 0; i < st.N; i++ {
		_ = append(newKeys, keys[i])
		_ = append(newValues, value[i])
	}
	st.Keys = newKeys
	st.Value = newValues
}

func (st *BinarySearchST) rank(k string) int {
	keys := st.Keys
	lo := 0
	high := st.N - 1
	for lo <= high {
		mid := lo + (high-lo)/2
		compare := strings.Compare(keys[mid], k)
		if compare < 0 {
			// 这里又一次命中了性能不优
			//lo = mid
			lo = mid + 1
		} else if compare > 0 {
			// 这里又一次命中了性能不优
			//high = mid
			high = mid - 1
		} else {
			// 这里又一次命中了性能不优
			return mid
		}
	}
	return lo
}

func (st *BinarySearchST) Put(k string, v interface{}) {
	// 下移动，因为有可能只是更新数据
	//if st.N == len(st.Keys) {
	//	st.resize(len(st.Keys) << 1)
	//}
	keys := st.Keys
	value := st.Value
	rank := st.rank(k)
	if st.N == len(st.Keys) {
		st.resize(len(st.Keys) << 1)
	}
	if strings.Compare(keys[rank], k) == 0 {
		value[rank] = v
		return
	}
	for i := st.N; i > rank; i-- {
		keys[i] = keys[i-1]
		value[i] = value[i-1]
	}
	keys[rank] = k
	value[rank] = v
	st.N++
}

func (st *BinarySearchST) Del(k string) bool {
	if st.IsEmpty() {
		return false
	}
	keys := st.Keys
	value := st.Value
	rank := st.rank(k)
	if rank < st.N && strings.Compare(keys[rank], k) == 0 {
		for i := rank; i < st.N; i++ {
			keys[i] = keys[i+1]
			value[i] = value[i+1]
		}

		keys[st.N] = ""
		value[st.N] = nil
		st.N--
		return true
	}
	return false
}

func (st *BinarySearchST) Contains(k string) bool {
	rank := st.rank(k)
	return rank < st.N && strings.Compare(st.Keys[rank], k) == 0
}

func main() {
	st := NewBinarySearchST(10)
	fmt.Println(st.IsEmpty())
	fmt.Println(st.Size())

	st.Put("1", 1)
	fmt.Println(st.Contains("1"))
	fmt.Println(st.IsEmpty())
	fmt.Println(st.Size())

	// put
	st.Put("3", 3)
	st.Put("4", 4)
	st.Put("2", 2)
	st.Put("6", 6)
	st.Put("5", 5)

	keys := st.Keys
	fmt.Println("--------------")
	for _, key := range keys {
		fmt.Println(key)
	}
	fmt.Println("--------------")

	fmt.Println(st.Get("1"))
	fmt.Println(st.Get("5"))

	fmt.Println("--------------")
	fmt.Println(st.Del("1"))
	fmt.Println(st.Del("5"))
	for _, key := range keys {
		fmt.Println(key)
	}
	fmt.Println("--------------")

}
