package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello World Here I Am"
	split := strings.Fields(str)
	var r []string
	for _, v := range split {
		s := []string{v}
		r = append(s, r...)
	}
	fmt.Println(str, "reversed is:")
	fmt.Println(strings.Join(r, " "))
}
