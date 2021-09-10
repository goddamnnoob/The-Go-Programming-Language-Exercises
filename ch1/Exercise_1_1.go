package main

//Mo dif y the echoprog ram to als o pr int os.Args[0], t he name of t he −ommandthat invo ked it.

import (
	"fmt"
	"os"
	"strings"
)

func EX11() {
	fmt.Println(strings.Join(os.Args[:], ""))
}
