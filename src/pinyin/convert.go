package pinyin

import (
  "fmt"
)

func Convert(strs string) {
  for _, v := range strs {
    fmt.Println(get(int(v)))
  }
}
