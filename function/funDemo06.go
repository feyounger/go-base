package main

func deferT() {
	x, y := 10, 20

	defer func(i, j int) {
		println("defer:", i, j)
	}(x, y)

	x += 100
	y += 10
	println("x = ", x , "y = ",y)
}

func main() {
	deferT()
}
