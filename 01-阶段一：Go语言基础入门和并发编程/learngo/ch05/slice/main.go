package main

import (
	"fmt"
	"testing"
)

func main() {
	//go折中， slice 切片 - 数组
	var courses []string
	fmt.Printf("%T\r\n", courses)

	//这个方法很特别， 第一次用很头疼， 原理
	//courses = append(courses, "go")
	//courses = append(courses, "grpc")
	//courses = append(courses, "gin")
	//
	//fmt.Println(courses[1])
	//
	////切片的初始化 3种：1：从数组直接创建 2：使用string{} 3: make
	//allCourses := [5]string{"go", "grpc", "gin", "mysql", "elasticsearch"}
	//courseSlice := allCourses[0:len(allCourses)] //左闭右开 [) python
	//fmt.Println(courseSlice)

	//courseSlice := []string{"go", "grpc", "gin", "mysql", "elasticsearch"}

	////make 函数
	//allCourses2 := make([]string, 3)
	//allCourses2[0] = "c"
	//allCourses2[1] = "c"
	//allCourses2[2] = "c"
	//fmt.Println(allCourses2)
	//
	//var allCourses3 []string
	//allCourses3 = append(allCourses3, "c")
	//fmt.Println(allCourses3)

	//访问切片的元素， 访问单个， 访问多个
	//fmt.Println(courseSlice[1])
	//
	////[start:end]
	///*
	//1. 如果只有start 没有end 就表示从start开始到结尾的所有数据
	//2. 如果没有start 有end 就表示从0到end之前的所有数据
	//3. 如果有start 没有有end
	//4. 有start 有end
	// */
	//fmt.Println(courseSlice[1:4])
	//fmt.Println(courseSlice[1:])
	//fmt.Println(courseSlice[:])

	//courseSlice := []string{"go", "grpc"}
	//courseSlice2 := []string{"mysql", "es", "gin"}
	//courseSlice = append(courseSlice, courseSlice2[1:]...)
	//
	////如何删除slice中的元素： 比较麻烦
	//fmt.Println(courseSlice)

	courseSlice := []string{"go", "grpc", "mysql", "es", "gin"}
	myslice := append(courseSlice[:2], courseSlice[3:]...)
	fmt.Println(myslice)

	courseSlice = courseSlice[:3]
	fmt.Println(courseSlice)

	//复制 slice
	//courseSliceCopy := courseSlice
	courseSliceCopy2 := courseSlice[:]
	fmt.Println(courseSliceCopy2)

	var courseSliceCopy = make([]string, len(courseSlice))
	copy(courseSliceCopy, courseSlice)
	fmt.Println(courseSliceCopy)

	fmt.Println("------------------------")
	courseSlice[0] = "java"
	fmt.Println(courseSliceCopy2)
	fmt.Println(courseSliceCopy)

	////删除指定元素
	//arraydata := []int{2,3,4,5,6}
	//bbss := delectData1(arraydata,5)
	//fmt.Println(bbss)

}


/**
1.截取法（修改原切片）删除指定元素
这里利用对 slice 的截取删除指定元素。注意删除时，后面的元素会前移，所以下标 i 应该左移一位。
 */
func delectData1( a []int,stringInt int) []int {
	for i := 0 ;i<len(a);i++ {
		if a[i] == stringInt{
			a = append(a[:i],a[i+1:]...)
			i--
		}
	}
	return a
}

/**
2.拷贝法（不改原切片）删除指定元素
这种方法最容易理解，重新使用一个 slice，将要删除的元素过滤掉。缺点是需要开辟另一个 slice 的空间，优点是容易理解，而且不会修改原 slice。
 */
func delectData2( a []int,stringInt int) []int {
	tmp := make([]int,0,len(a))
	for _,v := range a{
		if v  != stringInt{
			tmp = append(tmp,v)
		}
	}
	return tmp
}

/**
3.移位法（修改原切片）删除指定元素
3.1 方式一
利用一个下标 index，记录下一个有效元素应该在的位置。遍历所有元素，当遇到有效元素，将其移动到 index 且 index 加一。最终 index 的位置就是所有有效元素的下一个位置，最后做一个截取就行了。这种方法会修改原来的 slice。
该方法可以看成对第一种方法截取法的改进，因为每次指需移动一个元素，性能更加。
 */
func delectData3( a []int,stringInt int) []int {
	i := 0
	for _,vuelis := range a{
		if vuelis != stringInt{
			a[i] = vuelis
			i++
		}
	}
	return a[:i]
}

/**
3.2 方式二
创建了一个 slice，但是共用原始 slice 的底层数组。这样也不需要额外分配内存空间，直接在原 slice 上进行修改。
 */
func delectData4( a []int,stringInt int) []int {
	tgt := a[:0]
	for _,v := range a{
		if v != stringInt{
			tgt = append(tgt,v)
		}
	}
	return tgt
}


/**
4.性能对比
假设我们的切片有 0 和 1，我们要删除所有的 0。
这里分别对长度为 10、100、1000 的切片进行测试，来上下上面四种实现的性能差异。
生成切片函数如下：
 */
func getSlice(n int)[]int  {
	a := make([]int , 0 , n)
	for i := 0;i < n; i++{
		if i%2 == 0{
			a = append(a,0)
			continue
		}
		a = append(a,1)
	}
	return a
}
//基准测试代码如下：
func BenchmarkDeleteSloce1(b *testing.B)  {
	for i := 0; i<b.N; i++ {
		_ = delectData1(getSlice(10),0)
	}
}
func BenchmarkDeleteSloce2(b *testing.B)  {
	for i := 0; i<b.N; i++ {
		_ = delectData2(getSlice(10),0)
	}
}
func BenchmarkDeleteSloce3(b *testing.B)  {
	for i := 0; i<b.N; i++ {
		_ = delectData3(getSlice(10),0)
	}
}
func BenchmarkDeleteSloce4(b *testing.B)  {
	for i := 0; i<b.N; i++ {
		_ = delectData4(getSlice(10),0)
	}
}

/**
5.小结
从基准测试结果来看，性能最佳的方法是移位法，其中又属第一种实现方式较佳。性能最差的也是最常用的方法是截取法。随着切片长度的增加，上面四种删除方式的性能差异会愈加明显。
实际使用时，我们可以根据不用场景来选择。如不能修改原切片使用拷贝法，可以修改原切片使用移位法中的第一种实现方式。
 */

