package main

import (
	"fmt"
)

const flag = "golang>"

func main() {

	for true {
		str := ""
		fmt.Printf("%s", flag)
		fmt.Scanf("%s", &str)

		if str == "exit" {
			break
		} else if str == "test" {
			fmt.Printf("hello, my dear friend!\n")
		} else if str == "" {
			fmt.Printf("%s\r", flag)
		} else {
			fmt.Printf(">>> %s\n", str)
		}
	}
}
