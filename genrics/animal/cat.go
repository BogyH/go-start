package animal

import "fmt"

type Cat struct {
}

func (cat *Cat) Move() {
	fmt.Println("Cat moving~")
	fmt.Println("move 2 steps")
	fmt.Println("seat down")
	fmt.Println("licking left hand")
	fmt.Println("get down and spleep...")
}
