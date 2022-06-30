package main

func deferTest(x int)  {
	defer println("a")
	defer println("b")

	defer func() {
		println(100/x)
	}()

	defer println("c")
}

func main() {
	deferTest(0)
}
