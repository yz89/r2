package main

import (
	"fmt"
	"r2/context"
	"r2/interp"
)

func main() {
	// exp, err := interp.MatchSexp(interp.R29)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(exp)
	// for _, e := range exp {
	// 	fmt.Println(e)
	// }

	env0 := context.Env{}
	res, err := interp.Execute(interp.R27, env0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
