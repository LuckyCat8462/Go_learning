package main

import "fmt"

func main() {
	// 一、概念
	// 切片：slice，它的底层是数组，可以动态改变长度，是一种对数组的抽象。
	// 1.1、定义一个切片slice
	// names := [10]string{"北京", "上海", "广州", "深圳"}
	names := []string{"北京", "上海", "广州", "深圳"} // 切片不需要说明长度
	for i, v := range names {
		fmt.Println("i:", i, "v:", v)
	}
	// 二、基础操作
	// 2.1.追加数据-append()函数
	// 允许追加空元素names1 = append(names1, 0)
	names1 := append(names, "杭州")
	fmt.Println("names:", names)
	fmt.Println("names1:", names1)

	// 2.2.长度-len()函数
	fmt.Println("追加元素前,name的长度:", len(names), "，容量:", cap(names))
	names = append(names, "杭州")

	// 追加元素之后，如果容量不足，会动态分配2倍空间
	fmt.Println("追加元素后,name的长度:", len(names), "，容量:", cap(names))
	names = append(names, "云南")
	fmt.Println("追加元素后,name的长度:", len(names), "，容量:", cap(names))

	// 2.3.容量-cap()函数-测量切片最长可以达到多少
	// 对于一个切片，不仅有'长度'的概念len()，还有一个'容量'的概念cap()

	// 验证容量不足，会动态分配2倍空间
	nums := []int{}
	for i := 0; i < 35; i++ {
		nums = append(nums, i)
		fmt.Println("len:", len(nums), ", cap:", cap(nums))

	}

	// 2.4、copy()函数
	// 在截取的过程中，可能会污染原函数，如果想让切片完全独立于原始数组，可以使用copy()函数来完成
	namesCopy := make([]string, len(names))
	// func copy(dst, src []Type) int
	// names是一个数组，copy函数介收的参数类型是切片，所以需要使用[:]将数组变成切片
	copy(namesCopy, names[:])
	namesCopy[0] = "哈尔滨"
	fmt.Println("namesCopy:", namesCopy)
	fmt.Println("naemes本身:", names)

	//	三、常规操作
	//	3.1.基于旧数组创建新数组
	names3 := [7]string{"北京", "上海", "广州", "深圳", "杭州", "重庆", "南京"}

	// 想基于names3创建一个新的数组
	names4 := [3]string{}
	names4[0] = names3[0]
	names4[1] = names3[1]
	names4[2] = names3[2]

	//	切片可以基于一个数组，灵活的创建新的数组
	names5 := names3[0:3] // 左闭右开，左边第0个要，右边第3个不要
	fmt.Println("names5", names5)

	// 3.2.切片截取
	//	3.2.1.如果从0元素开始截取，那么冒号左边的数字可以省略
	names_a := names3[:5]
	fmt.Println("name_a:", names_a)

	// 3.2.2. 如果截取到数组最后一个元素，那么冒号右边的数字可以省略
	names_b := names3[5:]
	fmt.Println("name_b:", names_b)

	// 3.2.3. 如果想从左至右全部使用，那么冒号左右两边的数字都可以省略
	names_c := names3[:]
	fmt.Println("name_c:", names_c)

	// 3.2.4. 也可以基于一个字符串进行切片截取 : 取字符串的字串 hello world
	sub1 := "hello world"[5:7]
	fmt.Println("sub1:", sub1) // 'wo'

	//	3.3.nil空切片
	//	可以在创建空切片的时候，明确指定切片的容量，这样可以提高运行效率

	// 创建一个容量是20，当前长度是0的string类型切片
	// make的时候，初始的值都是对应类型的零值 : bool ==> false, 数字==> 0, 字符串 ==> 空
	str2 := make([]string, 10, 20) // 第三个参数不是必须的，如果没填写，则默认与长度相同
	fmt.Println("str2:", &str2[0])

	fmt.Println("str2: len:", len(str2), ", cap:", cap(str2))
	str2[0] = "hello"
	str2[1] = "world"

}
