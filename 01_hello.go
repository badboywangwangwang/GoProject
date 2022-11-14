//1)go语言已包作为管理单位
//2)每个文件必须先声明包
//3)程序必须有一个main包
package main

import (
	"fmt"
	"strings"
)

func test()(a,b,c int)  {
	return 1,2,3
}
func main() {

	//布尔值：
	//%t    单词true或false

	//整数：
	//%b    表示为二进制
	//%c    该值对应的unicode码值
	//%d    表示为十进制
	//%o    表示为八进制
	//%q    该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
	//%x    表示为十六进制，使用a-f
	//%X    表示为十六进制，使用A-F
	//%U    表示为Unicode格式：U+1234，等价于"U+%04X"

	//浮点数、复数的两个组分：
	//%b    无小数部分、二进制指数的科学计数法，如-123456p-78；参见strconv.FormatFloat %e    科学计数法，如-1234.456e+78 %E    科学计数法，如-1234.456E+78 %f    有小数部分但无指数部分，如123.456 %F    等价于%f %g    根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	//%G    根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）

	//字符串和[]byte：
	//%s    直接输出字符串或者[]byte %q    该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	//%x    每个字节用两字符十六进制数表示（使用a-f）
	//%X    每个字节用两字符十六进制数表示（使用A-F）

	//指针：
	//%p    表示为十六进制，并加上前导的0x

	//%f:    默认宽度，默认精度
	//%9f    宽度9，默认精度
	//%.2f   默认宽度，精度2 %9.2f  宽度9，精度2 %9.f   宽度9，精度0


	//+    总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）；
	//-    在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）；
	//#    切换格式：
	//八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）；
	//对%q（%#q），如果strconv.CanBackquote返回真会输出反引号括起来的未转义字符串；
	//对%U（%#U），如果字符是可打印的，会在输出Unicode格式、空格、单引号括起来的go字面值；
	//' '    对数值，正数前加空格而负数前加负号；
	//对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格；
	//0    使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面；

	/**
	1.初始化（变量可复用）
	 */
	var a int
	a = 10
	a = 20
	a = 30
	fmt.Println("1.初始化（变量可复用）:",a)

	/**
	2.初始化并且赋值
	*/
	var b int = 10
	fmt.Println("2.初始化并且赋值（变量可复用）:",b)

	/**
	3.初始化并且赋值（2）变量不可复用
	*/
	c := 30
	fmt.Println("3.初始化并且赋值（变量不可复用）:",c)

	/**
	4.字符串拼接
	*/
	d := 40
	fmt.Printf("4.字符串拼接：%d\n",d)

	/**
	5.多重赋值（赋值转换）
	*/
	e,f := 10,20
	e,f = f,e
	fmt.Printf("5.多重赋值（赋值转换）:%d:%d\n",e,f)

	/**
	6.匿名变量
	*/
	_,b,_ = test()
	fmt.Printf("6.匿名变量:%d\n",b)

	/**
	7.常量
	*/
	const i int = 10

	fmt.Printf("7.常量:%d\n",i)

	/**
	8.数据类型
	*/
	const j  = 10
	fmt.Printf("8.数据类型:%T\n",j)

	/**
	9.多重变量赋值
	*/
	var(
		k = 10
		l = 2.0
	)
	fmt.Printf("9.多重变量赋值：%d:%f\n",k,l)

	/**
	10.多重常量赋值
	*/
	const(
		n = 10
		m = 2.0
	)
	fmt.Printf("10.多重常量赋值：%d:%f\n",n,m)

	/**
	11.枚举
	*/
	const(
		o = iota
		p = iota
		q = iota
	)
	fmt.Printf("11.枚举：%d:%d:%d\n",o,p,q)

	/**
	12.枚举(一个iota)
	*/
	const(
		r = iota
		s
		t
	)
	fmt.Printf("12.枚举(一个iota)：%d:%d:%d\n",r,s,t)

	/**
	13.枚举(同一行iota)
	*/
	const(
		u = iota
		v,w,y =iota,iota,iota
		z = iota
	)
	fmt.Printf("13.枚举(同一行iota)：%d:%d:%d:%d:%d\n",u,v,w,y,z)

	/**
	14.布尔类型
	*/
	//var aa bool
	aa := false
	var aaa string
	if(aa){
		aaa = "123"
	}else{
		aaa = "456"
	}
	fmt.Printf("14.布尔类型:%t:%s\n",aa,aaa)

	/**
	15.字符串类型
	*/
	var str string = "我是字符串"
	str1 := "我是字符串1"
	fmt.Printf("15.字符串类型:%s:%c:%s:%d\n",str,str[0],str1,len(str1))

	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, "神")
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])
}