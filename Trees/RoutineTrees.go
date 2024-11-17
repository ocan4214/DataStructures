package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

var runtime int = 0

// Walk walks the tree t sending all values
// from the tree to the channel ch.

func Walk(t *tree.Tree, ch chan int) {
	Recursive(t, ch)
	close(ch)
}

func Recursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Recursive(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Recursive(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	var ch1 chan int = make(chan int)
	var ch2 chan int = make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		select {
		case val1 := <-ch1:
			if val1 != <-ch2 {
				return false
			}
		case val2 := <-ch2:
			if val2 != <-ch1 {
				return false
			}
		}
		return true
	}
}

func main() {

	var root *tree.Tree = tree.New(1)
	var ch1 chan int = make(chan int, 10)

	go Walk(root, ch1)

	for v := range ch1 {
		fmt.Printf("Tree Val is %v \n", v)
	}

	fmt.Printf("Result is %v \n", Same(root, tree.New(2)))
}
