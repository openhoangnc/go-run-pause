package main

import (
	"fmt"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Current working directory:", pwd)
	fmt.Println("Args:", os.Args)
	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
