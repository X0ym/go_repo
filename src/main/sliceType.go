package main

import "fmt"

/**
切片（Slice）是相同元素的可变长度的序列，slice是一种引用类型

声明切片
	var name []T
	len()求切片长度 cap()求切片容量
切片表达式 从数组，字符串，切片中构造切片
	方式一：简单切片表达式 name[low : high]
		1. 切片表达式中的low和high表示一个索引范围（左包含，右不包含）
		2. 省略了low则默认为0；省略了high则默认为切片操作数的长度
			a[2:]  // 等同于 a[2:len(a)]
			a[:3]  // 等同于 a[0:3]
			a[:]   // 等同于 a[0:len(a)]
		3. 对于数组或字符串，low和high需要满足：0 <= low <= high <= len(slice)，则索引合法，否则就会索引越界（out of range）
		4. 对于切片，low和high需要满足：0 <= low <= high <= cap(slice),因为切片的容量才是切片底层数组的长度
	方式二：完整切片表达式 name[low : high : max]
		1. 支持数组，指向数组的指针，切片；不支持字符串
		2. 构造与简单切片表达式a[low: high]相同类型、相同长度和元素的切片，且将得到的结果切片的容量设置为 max - low
		3. 在完整切片表达式中只有第一个索引值（low）可以省略；它默认为0
		4. 完整切片表达式需要满足的条件是 0 <= low <= high <= max <= cap(a)
使用make()函数构造切片
	make([]T, size, cap)

切片的本质
	切片的本质是对底层数组的封装，包含了三个信息：底层数组的指针，切片的长度和切片的容量
	切片是引用类型，不能直接使用 == 比较，可与 nil 比较
	一个 nil 值得切片没有底层数组，其长度和容量都为0。但不能说一个长度和容量为0的切片一定是nil
	判断切片非空用 len(s) == 0,而不是 s == nil
	切片可能会共享同一个底层数组，对一个切片的修改可能会影响另一个切片的内容
切片的遍历
	与数组相同
切片添加元素
	使用 append() 方法为切片添加元素
	var s []int
	s = append(s,1) // [1]
	s = append(s,2,3,4) // [1 2 3 4]
	s2 := []int {5,6,7}
	s = append(s, s2...)
切片删除元素
	没有特定的删除函数
	删除 索引为 index 的元素 s = append(s[:2], s[3:])
	删除区间  [l, r] 的元素 s = append[s[:l], s[r+1:])
*/

func sliceTest() {
	name1 := []int{1, 2, 3}
	fmt.Println(name1)
	fmt.Println(len(name1))
	fmt.Println(cap(name1))

	// 切片表达式 方式一
	a := [5]int{1, 2, 3, 4, 5}
	s1 := a[1:3] // s := a[low:high]
	fmt.Printf("slice s1:%v len(s1):%v cap(s1):%v\n", s1, len(s1), cap(s1))

	//s2 := s1[3:4]
	//fmt.Printf("slice s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))

	// slice append
	var s []int
	s = append(s, 1)       // 1. append 一个值
	s = append(s, 2, 3, 4) // 2. append 多个值
	s2 := []int{5, 6, 7, 8}
	s = append(s, s2...) // 3. append slice

	var s3 []int
	for i := 0; i < 10; i++ {
		s3 = append(s3, i)
		fmt.Printf("slice s3:%v ,len(s3):%v, cap(s3):%v, ptr: %p \n", s3, len(s3), cap(s3), s3)
	}
	// slice copy
	// copy函数实现切片复制
	s11 := []int{1, 2, 3, 4, 5}
	s22 := make([]int, 5)
	copy(s22, s11)
	fmt.Println(s22)
	//slice delete

	s4 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s4)
	s4 = append(s4[:2], s4[4:]...)
	fmt.Printf("s4删除索引[2,3]的元素：%v", s4)
	//slice
}
