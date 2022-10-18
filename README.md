### About
[lisp-rs](https://github.com/vishpat/lisp-rs) の Golang によるリファレンス実装です。

### Usage

REPL の起動：

```
$ go run main.go
```

実行例：

```
$ go run main.go
> (define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))
Void
> (fib 20)
10946
```

### Test

テストの実行：

```
$ go test ./... 
```