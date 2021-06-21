package main

import (
	"fmt"

	"github.com/lone1y-51/utils/tree"
)

func main() {
	tr := &tree.BRTree{}
	_ = tr.Insert(1)
	fmt.Println(tr.ToString())
	_ = tr.Insert(3)
	fmt.Println(tr.ToString())
	_ = tr.Insert(5)
	fmt.Println(tr.ToString())
	_ = tr.Insert(10)
	fmt.Println(tr.ToString())
	_ = tr.Insert(8)
	fmt.Println(tr.ToString())
	_ = tr.Insert(9)
	fmt.Println(tr.ToString())
	_ = tr.DeleteValue(10)
	fmt.Println(tr.ToString())
	fmt.Println("lone1y utils")
}
