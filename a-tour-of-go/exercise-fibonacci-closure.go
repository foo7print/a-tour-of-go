// 関数を用いた面白い例を見てみましょう。

// fibonacci (フィボナッチ)関数を実装しましょう。この関数は、連続するフィボナッチ数(0, 1, 1, 2, 3, 5, ...)を返す関数(クロージャ)を返します。

package main

import "fmt"

func fibonacci() func() int {
	z0, z1 := 0, 1
	return func() int {
		r := z0
		z0, z1 = z1, z0+z1
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
