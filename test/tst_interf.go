package main

import(
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String()string{
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    var a IPAddr
    a[0] = 4
    a[1] = 5
    a[2] = 6
    a[3] = 7
	fmt.Println(a)
}
