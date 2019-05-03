// 関数とループを使った簡単な練習として、平方根の計算を実装してみましょう: 数値 x が与えられたときに z² が最も x に近い数値 z を求めたいと思います。

// コンピュータは通常ループを使って x の平方根を計算します。 いくつかの z を推測することから始めて、z² がどれほど x に近づいているかに応じて z を調整できます。

// z -= (z*z - x) / (2*z)

// 実際の平方根に近い答えになるまでこの調整を繰り返すことによって、推測はより良いものなります。

// これを func Sqrt に実装してください。 何が入力されても z の適切な開始推測値は 1 です。 まず計算を 10 回繰り返してそれぞれの z を表示します。 x (1, 2, 3, ...) のさまざまな値に対する答えがどれほど近似し、 推測が速くなるかを確認してください。

// Hint: 浮動小数点の変数を初期化して宣言するには、型でキャストするか、浮動小数点を使ってみてください:

// z := 1.0
// z := float64(1)

// 次に値が変化しなくなった (もしくはごくわずかな変化しかしなくなった) 場合にループを停止させます。 それが 10 回よりも多いか少ないかを確認してください。 x や x/2 のように他の初期推測の値を z に与えてみてください。 あなたの関数の結果は標準ライブラリの math.Sqrt にどれくらい近づきましたか？

// (Note: If you are interested in the details of the algorithm, the z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing. この一般的なアプローチはニュートン法と呼ばれています。 多くの関数で有効に働きますがとくに平方根では殊更有効です。)

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	// 1e-9は0.000000
	// for diff := 1.0; math.Abs(diff) > 1e-9; {
	// 	diff = (z * z - x) / (2 * z)
	// 	z = z -diff
	// }

	// i++の後ろのセミコロンいらない
	// for i := 0; i < 1000; i++; {
	// 	z -= (z * z - x) / (2 * z)
	// }

	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
