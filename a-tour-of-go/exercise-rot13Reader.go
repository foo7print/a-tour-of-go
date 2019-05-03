// よくあるパターンは、別の io.Reader をラップし、ストリームの内容を何らかの方法で変換するio.Readerです。

// 例えば、 gzip.NewReader は、 io.Reader (gzipされたデータストリーム)を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開したデータストリーム)を実装しています。

// io.Reader を実装し、 io.Reader でROT13 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用して読み出すように rot13Reader を実装してみてください。

// rot13Reader 型は提供済みです。 この Read メソッドを実装することで io.Reader インタフェースを満たしてください。

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r3r rot13Reader) Read(buf []byte) (int, error) {
	size, err := r3r.r.Read(buf)
	if err == io.EOF {
		return size, err
	}

	for i, char := range buf {
		switch {
		case 'A' <= char && char <= 'M':
			buf[i] = 'N' + (char - 'A')
		case 'N' <= char && char <= 'Z':
			buf[i] = 'A' + (char - 'N')
		case 'a' <= char && char <= 'm':
			buf[i] = 'n' + (char - 'a')
		case 'n' <= char && char <= 'z':
			buf[i] = 'a' + (char - 'n')
		default:
			continue
		}
	}

	return size, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
