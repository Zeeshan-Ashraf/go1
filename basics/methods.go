package basics

import "fmt"

type Structz struct {
	Id   int
	Name string
}

func (z Structz) Fnzee() {
	fmt.Printf("%d", z.Id)
}
