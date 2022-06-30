package main

func main() {
	println(add(1, 2))
}

func add(a, b int) (c int) {
	defer func() {
		c += 10
	}()
	c = a + b
	return
}
