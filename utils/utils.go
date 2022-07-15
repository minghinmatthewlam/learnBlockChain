package utils

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) PrintInfo() {
	p.Name = "Matthew"
	p.Age = 23
	fmt.Println(p.Name, p.Age)
}
