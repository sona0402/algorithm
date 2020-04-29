package main

import "fmt"

// Marker interface used by {@code List} implementations to indicate that
// * they support fast (generally constant time) random access.
func main() {

	var intAddress [10]int

	// 深入理解计算机系统和汇编中都有定义 地址=基地址+偏移地址
	// 随机访问时间复杂度为O(1) 可以参考Java RandomAccess
	// for typical instances of the class, this loop:

	//for (int i=0, n=list.size(); i &lt; n; i++)
	//        list.get(i);
	// runs faster than this loop:
	//for (Iterator i=list.iterator(); i.hasNext(); )
	//         i.next();

	// 这也就是Java用数组实现的数据会使用RandomAccess进行标记的原因。

	for i, _ := range intAddress {
		fmt.Println(&intAddress[i])
	}

	//0xc0000180a0
	//0xc0000180a8
	//0xc0000180b0
	//0xc0000180b8
	//0xc0000180c0
	//0xc0000180c8
	//0xc0000180d0
	//0xc0000180d8
	//0xc0000180e0
	//0xc0000180e8

}
