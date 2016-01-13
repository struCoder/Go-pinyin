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
  str := "中国"
  py := pinyingo.NewPy(pinyingo.STYLE_TONE)       //string with tone        -> 中国: ["zhōng", "guó"]
  //py := pinyingo.NewPy(pinyingo.STYLE_NORMAL)   //string without tone     -> 中国: ["zhong", "guo"]
  //py := pinyingo.NewPy(pinyingo.STYLE_INITIALS) // get initials of string -> 中国: ["zh", "g"]
  fmt.Println(py.Convert(str))
}

```

Features
====
-  [x] convert han zi to pinyin
-  [x] get han zi initial
-  [ ] deal with heteronym


License
========
#### MIT
