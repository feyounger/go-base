package extend

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet int8
	*Animal
}

func (d Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main()  {
	dog := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "XI",
		},
	}
	dog.move()
	dog.wang()
}