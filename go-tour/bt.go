package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	Walk(t.Right, ch)
	ch <- t.Value
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, _ := <-ch1
		v2, _ := <-ch2
		if v1 != v2 {
			return false
		}
	}

}

func main() {
	a := tree.New(5)
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(Same(a, a))
}
