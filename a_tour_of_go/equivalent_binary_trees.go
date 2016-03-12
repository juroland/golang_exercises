package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var recurseWalk func(t *tree.Tree)
	recurseWalk = func(t *tree.Tree) {
		if t != nil {
			recurseWalk(t.Left)
			ch <- t.Value
			recurseWalk(t.Right)
		}
	}
	recurseWalk(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for val1 := range ch1 {
		if val1 != <-ch2 {
			return false
		}
	}
	_, values := <-ch2
	return !values
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
