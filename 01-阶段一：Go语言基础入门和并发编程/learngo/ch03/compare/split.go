package main

import (
	"fmt"
	"strings"
)

func main() {

	/**
	返回值:返回一个字符串切片。
	使用示例
	Split()函数将字符串根据分隔符切割。切割后返回一个字符串切片，切片len和cap值等于原字符串中存在分隔符的数量 + 1（仅在s不是空字符串的情况下成立）。
	 */

	demo := "I&love&Go,&and&I&also&love&Python."
	string_slice := strings.Split(demo, "&")

	fmt.Println("demo:result:",string_slice)
	fmt.Println("demo:len:",len(string_slice))
	fmt.Println("demo:cap:", cap(string_slice))


	/**
	注意事项
	1. 当分隔符不存在于原字符串中时
	当分隔符在原字符串中不存在的情况下，Split()函数仅仅将原字符串转换成一个len和cap值都为1的字符串切片。
	 */
	demo1 := "I love Go, and I also love Python."
	string_slice1 := strings.Split(demo1, "&")

	fmt.Println("demo1:result:",string_slice1)
	fmt.Println("demo1:len:",len(string_slice1))
	fmt.Println("demo1:cap:", cap(string_slice1))

	/**
	2. 当分隔符是空字符串时
	当分隔符是空字符串时，Split()函数将字符串中的每一个字符分割成一个单独的元素。
	 */

	demo2 := "Go"
	string_slice2 := strings.Split(demo2, "")

	fmt.Println("demo2:result:",string_slice2)
	fmt.Println("demo2:len:",len(string_slice2))
	fmt.Println("demo2:cap:", cap(string_slice2))

	/**
	3. 参数都为空字符串
	当Split()函数的两个参数都是空字符串时（即s和sep都是空字符串），Split()函数返回一个len和cap值都为0的空字符串切片。
	 */
	demo3 := ""
	string_slice3 := strings.Split(demo3, "")

	fmt.Println("demo3:result:",string_slice3)
	fmt.Println("demo3:len:",len(string_slice3))
	fmt.Println("demo3:cap:", cap(string_slice3))

	/**
	4. 当s为空字符串，sep不为空字符串时
	不同于上一个场景，在这种情况下虽然得到的结果仍然是字符串切片，但是字符串切片的len和cap值是1（而不是像上一个场景中的值是0）。返回的结果是包含一个空字符串的字符串切片。
	 */

	demo4 := ""
	string_slice4 := strings.Split(demo4, "*")

	fmt.Println("demo4:result:",string_slice4)
	fmt.Println("demo4:len:",len(string_slice4))
	fmt.Println("demo4:cap:", cap(string_slice4))

	element := string_slice4[0]
	fmt.Printf("string_slice[0]:%s, type:%T, len():%d\n",
		element, element, len(element),)

	fmt.Println("element == \"\"?", element == "")


	/**
	5. 返回的字符串切片中不再包含分隔符
	 */
}