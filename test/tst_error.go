package main

import (
	"fmt"
)

type Errfunc float64

func (a Errfunc) Error() string {
	return "Ошибка елки-палки"
}

func Div(a,b Errfunc ) string {
    if b == 0 {
        return b.Error()
    } else {
        return fmt.Sprintf("%d", a/b)
    }
}

func main() {
    fmt.Print(Div(1, 5))
}
