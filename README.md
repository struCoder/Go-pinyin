Han yu pin yin package for Golang
==================================

Get Pinyin Of Simplified Chinese


Change log
============

#### 2017-04-24
1. fix STYLE\_FIRST\_LETTER

#### 2016-06-21
1. optimized code

#### 2016-03-18
1. optimize performance(resolve out of slice len when building)


Install
========

In you project root path just run `go get github.com/struCoder/Go-pinyin`


Test
====
```bash
go test

```


How to Use
===========
```go

import (
  "fmt"
  "github.com/struCoder/Go-pinyin"
)

func main() {
  str := "中国"
  str1 := "重阳"
  py := pinyingo.NewPy(pinyingo.STYLE_TONE, pinyingo.NO_SEGMENT)       //string with tone        -> 中国: ["zhōng", "guó"]
  //py := pinyingo.NewPy(pinyingo.STYLE_NORMAL, pinyingo.NO_SEGMENT)   //string without tone     -> 中国: ["zhong", "guo"]
  //py := pinyingo.NewPy(pinyingo.STYLE_INITIALS, pinyingo.NO_SEGMENT) // get initials of string -> 中国: ["zh", "g"]

  //segment
  py := pinyingo.NewPy(pinyingo.STYLE_TONE, pinyingo.USE_SEGMENT)       //string with tone        -> 重阳: ["chóng", "yáng"]

  fmt.Println(py.Convert(str))
}

```

Features
====
-  [x] convert han zi to pinyin
-  [x] get han zi initial
-  [x] deal with segment


support this project :)
=======================
![pay to me](http://7xjbiz.com1.z0.glb.clouddn.com/me/cEKqrjQYK8xnzzsY?imageView2/0/w/480)

License
========
#### MIT
