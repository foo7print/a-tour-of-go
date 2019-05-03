// Sqrt 関数を 以前の演習 からコピーし、 error の値を返すように修正してみてください。

// Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

// 新しい型:

// type ErrNegativeSqrt float64
// を作成してください。

// そして、 ErrNegativeSqrt(-2).Error() で、 "cannot Sqrt negative number: -2" を返すような:

// func (e ErrNegativeSqrt) Error() string
// メソッドを実装し、 error インタフェースを満たすようにします。

// 注意: Error メソッドの中で、 fmt.Sprint(e) を呼び出すことは、無限ループのプログラムになることでしょう。 最初に fmt.Sprint(float64(e)) として e を変換しておくことで、これを避けることができます。 なぜでしょうか？

// 負の値が与えられたとき、 ErrNegativeSqrt の値を返すように Sqrt 関数を修正してみてください。

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z0, z := 1.0, x
	for math.Abs(z-z0)/z > 1e-15 {
		z0, z = z, z-(z*z-x)/(2*z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
