package main

import "fmt"

func main() {
	//	一、概念
	// 1.1.MAP 在 Go 中称为字典或集合，是一种无序的键值对的集合。
	// 1.2.检索：Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。
	// 1.3.Map 是一种集合，可以像迭代数组和切片那样迭代Map。但是 Map 是无序的，所以遍历 Map 时返回的键值对的顺序是不确定的。
	// 1.4.在获取 Map 的值时，如果键不存在，返回该类型的零值。
	// 1.5.Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

	// 二、操作
	//	2.1.定义Map
	//	可以使用 内建函数make() 或使用 map关键字 来定义map
	var idNames map[int]string // 定义一个map，此时这个map是不能直接赋值的，它是空的

	// 2.2.分配空间
	// 使用make，可以不指定长度，但是指定长度性能更好。
	idNames = make(map[int]string, 10) // 这是最常用的方法
	idScore := make(map[int]float64)

	idNames[0] = "neko"
	idNames[1] = "zhi"

	// 2.4.遍历map
	for key, value := range idNames {
		fmt.Println("key:", key, "value:", value)
	}

	// 1.5. 删除map中的元素
	// 使用自由函数delete()来删除指定的key，参数为 map 和其对应的 key
	fmt.Println("idNames删除前:", idNames)
	delete(idNames, 1)
	delete(idNames, 100) // 不会报错
	fmt.Println("idNames删除后:", idNames)

	// 并发任务处理的时候，需要对map进行上锁

	//	三、注意点
	// 3.1、校验key是否存在	value,ok = addr[] if ok { } else { }
	//	//在map中不存在访问越界的问题，它认为所有的key都是有效的，所以访问一个不存在的key不会崩溃，而是返回这个类型的零值
	//	//零值：  bool=》false， 数字=》0，字符串=》空

	name9 := idNames[9]
	fmt.Println("name9:", name9)               // 空
	fmt.Println("idScore[100]:", idScore[100]) // 0

	// 无法通过获取value来判断一个key是否存在,因此需要一个能够校验key是否存在的方式
	value, ok := idNames[0] // 如果id=0是存在的，那么value就是key=0对应值，ok返回true, 反之返回零值，false
	if ok {
		fmt.Println("id=0这个key是存在的，value为:", value)
	}

	value, ok = idNames[10]
	if ok {
		fmt.Println("id=10这个key是存在的，value为:", value)
	} else {
		fmt.Println("id=10这个key不存在，value为:", value)
	}

}
