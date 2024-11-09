package main

import (
	"fmt"
)

func Prefx(sp []string) string {
	if len(sp) == 0 {
		return ""
	}
	pref := sp[0]
	for _, str := range sp[1:] {
		for len(pref) > 0 && len(str) < len(pref) || str[:len(pref)] != pref {
			pref = pref[:len(pref)-1]
		}
	}
	return pref
}

func main() {
	var a int
	sp := []string{}
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		inp := ""
		fmt.Scan(&inp)
		sp = append(sp, inp)
	}
	fmt.Println(Prefx(sp))
}
