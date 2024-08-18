package animal

import "fmt"

type Cat struct {
	name string
}

func NewCat(name string) Cat {
	cat := Cat{}
	cat.name = name
	return cat
}

func (cat *Cat) Move() {
	fmt.Printf("%v is moving~\n", cat.name)
	fmt.Println("moved 2 steps")
	fmt.Println("sit down")
	fmt.Println("licking left hand")
	fmt.Println("lie down and sleep...")
}

func (cat *Cat) Name() string {
	return cat.name
}
