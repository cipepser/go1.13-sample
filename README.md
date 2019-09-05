# go1.13-sample

[Go 1.13](https://blog.golang.org/go1.13)がリリースされたので触ってみる。

## Error wrapping

[リリースノート](https://golang.org/doc/go1.13)で一番気になったのが、Error wrappingだった。

[errorsのドキュメント](https://golang.org/pkg/errors/#pkg-index)を確認すると以下4つの関数が存在している。

```go
func As(err error, target interface{}) bool
func Is(err, target error) bool
func New(text string) error
func Unwrap(err error) error
```

`New`はただのコンストラクタなのでいいとして、残り3つは触ってみる。

### Isを使ってみる

```go
package main

import (
	"errors"
	"fmt"
)

var (
	MyError = myError()
)

func myError() error { return errors.New("myErr") }

func simpleError() error {
	return MyError
}

func main() {
	err := simpleError()
	if errors.Is(err, MyError) {
		fmt.Printf(err.Error()) // myErr
	}
}
```

ちなみに`main`を以下のようにしても結果は同じ。

```go
func main() {
	err := simpleError()
	if err != nil {
		fmt.Println(err) // myError
	}
}
```

これまでだったら以下のようにしていたと思う。

```go
func main() {
	err := simpleError()
	if err != nil {
		switch err {
		case MyError:
			fmt.Println("MyError:", err) // MyError: myErr
		default:
			fmt.Println("default:", err)
		}
	}
}
```

じゃあ、`Is`を用いて何が嬉しいのか。
本アップデートの名前にもあるwrappingしたときに生きてくる。
以下のようにwrapしたときの`err`の型を見てみる。

```go
func main() {
	err := wrappedError()
	fmt.Printf("%T", err) // *fmt.wrapError

	fmt.Println()

	err = simpleError()
	fmt.Printf("%T", err) // *errors.errorStringMyError: myErr
}
```

なので、先ほどの`switch`文を以下のように書き換えるとうまくエラーを捕まえることができない。

```go
func main() {
	err := wrappedError()
	if err != nil {
		switch err {
		case MyError:
			fmt.Println("MyError:", err)
		default:
			fmt.Println("default:", err) // default: myErr
		}
	}
}
```

`Is`を使うとちゃんと捕まえることができる。

```go
func main() {
	err := wrappedError()
	if errors.Is(err, MyError) {
		fmt.Printf(err.Error()) // myErr
	}
}
```



## References
- [Go 1\.13 is released \- The Go Blog](https://blog.golang.org/go1.13)
- [Go 1\.13 Release Notes \- The Go Programming Language](https://golang.org/doc/go1.13)
- [Go 1\.13 リリースノート \- Qiita](https://qiita.com/c-yan/items/b2f5e5c168d517594eb2)
- [Go1\.13からは今までのエラーハンドリングが機能しなくなるかもしれない \- Qiita](https://qiita.com/cia_rana/items/72a91175eadc1bffe9b0)
- [errors \- The Go Programming Language](https://golang.org/pkg/errors/#pkg-index)