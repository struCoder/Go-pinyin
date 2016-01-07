package pinyingo

import (
  "testing"
)

func TestConvert(t *testing.T) {
  pinyinMap := map[string]interface{}{
    "hello中国": []string{"hello", "zhōng", "guó"},
    "中国hello": []string{"zhōng", "guó", "hello"},
    "中国":      []string{"zhōng", "guó"},
    "123qwe":  []string{"123qwe"},
  }

  for k, v := range pinyinMap {
    converted := Convert(k)
    if converted != v {
      t.Errorf("%s is not equal %s", converted, v)
    }
  }
}
