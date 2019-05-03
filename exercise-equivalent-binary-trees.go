package main

import (
	"fmt"
	"sync"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
	done := make(chan bool, 1)
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	buf1, buf2 := make([]int, 0), make([]int, 0)

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		Walk(t1, ch1)
		wg.Done()
	}()
	go func() {
		Walk(t2, ch2)
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(done)
	}()

loop:
	for {
		select {
		case num := <-ch1:
			buf1 = append(buf1, num)

		case num := <-ch2:
			buf2 = append(buf2, num)

		case <-done:
			break loop
		}
	}

	if len(buf1) != len(buf2) {
		return false
	}
	for i := range buf1 {
		if buf1[i] != buf2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Walk")
	ch := make(chan int, 1)
	go func() {
		defer close(ch)
		Walk(tree.New(1), ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}

	fmt.Println("Check 1")
	actual1 := Same(tree.New(1), tree.New(1))
	if actual1 {
		fmt.Println("same")
	}

	fmt.Println("Check 2")
	actual2 := Same(tree.New(1), tree.New(2))
	if !actual2 {
		fmt.Println("not same")
	}
}
