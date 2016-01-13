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
  pinyinMap1 := map[string][]string{
    "hello中国": []string{"hello", "zh", "g"},
    "中国":      []string{"zh", "g"},
    "123qwe":  []string{"123qwe"},
  }
  pinyinMap2 := map[string][]string{
    "hello中国": []string{"hello", "zhong", "guo"},
    "中国":      []string{"zhong", "guo"},
    "123qwe":  []string{"123qwe"},
  }
  for k, v := range pinyinMap {
    py := NewPy(STYLE_TONE)
    converted := py.Convert(k)
    for i := 0; i < len(converted); i++ {
      if converted[i] != v[i] {
        t.Errorf("%s is not equal %s", converted, v)
      }
    }
  }

  for k, v := range pinyinMap1 {
    py := NewPy(STYLE_INITIALS)
    converted := py.Convert(k)
    for i := 0; i < len(converted); i++ {
      if converted[i] != v[i] {
        t.Errorf("%s is not equal %s", converted, v)
      }
    }
  }

  for k, v := range pinyinMap2 {
    py := NewPy(STYLE_NORMAL)
    converted := py.Convert(k)
    for i := 0; i < len(converted); i++ {
      if converted[i] != v[i] {
        t.Errorf("%s is not equal %s", converted, v)
      }
    }
  }
}
