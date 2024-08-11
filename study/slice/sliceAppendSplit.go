package main

import "fmt"

func main() {
	//表示设置一个长度为3,容量为5的数组,这里的容量其实就相当于底层存储数据的数组的真实长度,一旦元素个数超过容量,则进行扩容,扩容的方式为扩大到原来
	//容量的二倍
	//如果初始化时未指定容量,则默认的容量等于初始化时的长度,并且后续扩容也以当前容量进行2倍扩大
	slice := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
	// append方法表示向slice的末尾追加元素
	slice = append(slice, 15)
	slice = append(slice, 11)
	slice = append(slice, 12)
	slice = append(slice, 162)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
	fmt.Println("======================================")

	s1 := slice[2:4] //截取slice索引0~4的数据给s1,此时的s1也是一个slice数组
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)

	/**
	需要注意,这种截取s1实际上指向的是slice的内存地址,即当我们append(s1,5)时,slice的索引为4的数据也会被改变,因为他们共用一块内存空间
	这里非常重要,需要留意
	*/
	s1 = append(s1, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(slice), cap(slice), slice)
	fmt.Println("======================================")

	/**
	也是因为截取方法共用内存的原因,所以golang中给我们提供了copy方法,可以实现数组的深拷贝,不在共用一块内存地址,所以对数组的修改也不会影响彼此
	*/
	deepCopySlice := make([]int, 10)
	copy(deepCopySlice, s1)
	s1[1] = 1000
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(deepCopySlice), cap(deepCopySlice), deepCopySlice)
	deepCopySlice = append(deepCopySlice, 8)
	deepCopySlice[0] = 25
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(deepCopySlice), cap(deepCopySlice), deepCopySlice)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)
}
