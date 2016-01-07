Han yu pin yin package for Golang
==================================

Get Pinyin Of Simplified Chinese

Install
========

In you project root path just run `go get github.com/struCoder/Go-pinyin`


Test
====
just run `go test`


How to Use
===========
```go

import (
  "fmt"
  "github.com/struCoder/Go-pinyin"
)

func main() {
  str := "中文字符-汉语拼音"
  py := pinyin.NewPy()
  fmt.Println(py.Convert(str))
}

```
License
========
#### MIT
