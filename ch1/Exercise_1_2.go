package main

import (
	"fmt"
	"os"
)

func EX12() {
	var str string
	for arg := range os.Args {
		str += " " + fmt.Sprint(arg) + ":" + os.Args[arg]
	}
	fmt.Println(str)
}
