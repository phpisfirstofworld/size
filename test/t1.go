package main

import (
	"fmt"
	"github.com/phpisfirstofworld/size"
)

func main() {

	s := size.NewSize(56609)

	fmt.Println(s.Kb())
	fmt.Println(s.Mb())
	fmt.Println(s.Gb())

}
