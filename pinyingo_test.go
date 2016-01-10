package pinyingo

import (
  "testing"
)

func TestConvert(t *testing.T) {
  pinyinMap := map[string][]string{
    "hello中国": []string{"hello", "zhōng", "guó"},
    "中国hello": []string{"zhōng", "guó", "hello"},
    "中国":      []string{"zhōng", "guó"},
    "123qwe":  []string{"123qwe"},
  }

  for k, v := range pinyinMap {
    py := NewPy()
    converted := py.Convert(k)
    for i := 0; i < len(converted); i++ {
      if converted[i] != v[i] {
        t.Errorf("%s is not equal %s", converted, v)
      }
    }
  }
}
