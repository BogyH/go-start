package animal

import "fmt"

type Animal interface {
	Move()
	Name() string
}

func TouchAnimal(a Animal) {
	fmt.Printf("Touching an animal: %v\n", a.Name())
	a.Move()
	fmt.Printf("Touching done.")
}
