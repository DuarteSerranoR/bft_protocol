package main

//import "C"
import (
	"bft/Zyzzyva"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	Zyzzyva.Client()
	Zyzzyva.Server()
}
