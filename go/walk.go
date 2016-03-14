package main

import "code.google.com/p/go-tour/tree"
import (
    "fmt"
)
// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	WalkRecursive(t,ch)
    close(ch)
}

func WalkRecursive(t *tree.Tree, ch chan int) {
    if t == nil { return }
    WalkRecursive(t.Left, ch)
    ch <- t.Value
    WalkRecursive(t.Right, ch)
}    

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    go Walk(t1,ch1)
    go Walk(t2,ch2)
    
    for val1 := range ch1 {
        if val1 != <- ch2 {
            return false
        }
    }

    return true
}

func main() {
    tree1 := tree.New(1)
    tree2 := tree.New(1)
    
    fmt.Println(Same(tree1,tree2))
}