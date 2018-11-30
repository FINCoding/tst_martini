package main

import (
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	var str []string

	str = splitString(s)
	return calcCountMap(str)
}

func splitString(s string) []string{
	return strings.Fields(s)
}

func calcCountMap(s []string) map[string]int {
    var m map[string]int

    m = make(map[string]int)

	for i := range s{
        m[s[i]] = m[s[i]] + 1
	}
	return m
}

func main() {
	fmt.Print(WordCount("I ate a donut. Then I ate another donut."))
}
