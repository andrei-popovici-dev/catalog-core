package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World from Unikraft!")

	os.Stdout.WriteString("Hello directly to the system\n")

	err := os.Stdout.Sync()
	if err != nil {
		panic(err)

	}

}
