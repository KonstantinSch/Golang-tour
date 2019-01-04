package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wcArr := strings.Fields(s)
	wcMap := make(map[string]int)

	for _, v := range wcArr {
		fmt.Println(v)
		wcMap[v]++
	}

	return wcMap
}

//return map[string]int{"x":1}

func main() {
	//fmt.Println(WordCount("I`m a freelancer freelancer"))
	wc.Test(WordCount)
}
