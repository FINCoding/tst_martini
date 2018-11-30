package main

import (
    "strings"
    "fmt"
)

func WordCount(s string) {
    var str []string
    var m map[string]int
    m = make(map[string]int)
	str = strings.Fields(s)
	fmt.Print(str)
	for i := range str{
        m[str[i]] = m[str[i]] + 1
	}
	fmt.Print(m)
}

func main() {
	WordCount("Are you used to used ?")
}
