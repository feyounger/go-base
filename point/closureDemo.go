package main


import "fmt"

func main() {
	//Test()
	//Test1()// 5 5 5 5 5 除了初始化赋值外还被闭包修改--->捕获变量i的地址
	//Test2()// 4 4 4 4 4 除了初始化赋值外还被外层函数修改--->捕获变量t的地址
	Test3()// 0 1 2 3 4 只被初始化赋值--->值拷贝t
}

func Test()  {
	fmt.Println("test start--------------")
	fs := Closure()
	println(fs())
	println(fs())
}

func Closure() func() int {
	i := 2
	i++
	return func() int {
		return i
	}
}

func Test1() {
	fmt.Println("test1 start--------------")
	fs := Closure1()
	for _, f := range fs {
		f()
	}
	fmt.Println("test1 end--------------")
}

func Closure1() (fs [5]func()) {
	for i := 0; i < 5; i++ {
		fmt.Println("i address=",&i)
		fs[i] = func() {
			fmt.Println(i, &i)
		}
	}
	return
}

func Test2() {
	fmt.Println("test2 start--------------")
	fs := Closure2()
	for _, f := range fs {
		f()
	}
	fmt.Println("test2 end--------------")
}

func Closure2() (fs [5]func()) {
	t := 0
	for i := 0; i < 5; i++ {
		t = i
		fmt.Println("t address=",&t)
		fs[i] = func() {
			fmt.Println(t, &t)
		}
	}
	return
}

func Test3() {
	fmt.Println("test3 start--------------")
	fs := Closure3()
	for _, f := range fs {
		f()
	}
	fmt.Println("test3 end--------------")
}

func Closure3() (fs [5]func()) {
	for i := 0; i < 5; i++ {
		t := i
		fmt.Println("t address=",&t)
		fs[i] = func() {
			fmt.Println(t, &t)
		}
	}
	return
}