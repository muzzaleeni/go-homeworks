package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree, ch chan int)
	walker = func(t *tree.Tree, ch chan int) {
		if t != nil {
			walker(t.Left, ch)
			ch <- t.Value
			walker(t.Right, ch)
		}
	}
	walker(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for x := range ch1 {
		if x != <-ch2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for v := range ch {
		println(v)
	}
	println(Same(tree.New(1), tree.New(1))) // true
	println(Same(tree.New(1), tree.New(2))) // false
}
