package main

import (
	"fmt"
)

func main() {
	var item [16]int
	for i := 0; i < 16; i++ {
		item[i] = i
	}
	swapInt := swap(item)
	fmt.Print(swapInt)
}

// 面试题，将奇数放到前面 偶数放到后面
// 时间复杂度为O(n) 空间复杂度为 i , length
func swap(p [16]int) [16]int {
	i := 0
	length := len(p) - 1

	for i < length {
		if p[i]%2 == 0 && p[length]%2 != 0 {
			p[i], p[length] = p[length], p[i]
		}

		if p[i]%2 != 0 {
			i++
		}

		if p[length]%2 == 0 {
			length--
		}
	}

	return p
}
