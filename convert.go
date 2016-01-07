package pinyingo

import (
  "strings"
  "unicode/utf8"
)

// var reg = regexp.MustCompile("^[\u4e00-\u9fa5]$")

type options struct {
  heteronym bool
  style     int
  segment   bool
}

func (this *options) perStr(pinyinStrs string) string {
  //首字母
  if this.style == 1 {
    for i := 0; i < len(INITIALS); i++ {
      if strings.Index(pinyinStrs, INITIALS[i]) == 0 {
        return INITIALS[i]
      }
    }
    return ""
  } else if this.style == 0 {
    ret := strings.Split(pinyinStrs, ",")
    return ret[0]
  }
  return ""
}

func (this *options) Convert(strs string) []string {
  //获取字符串的长度
  bytes := []byte(strs)
  retArr := make([]string, 0)
  nohans := ""
  var tempStr string
  var single string
  for len(bytes) > 0 {
    r, w := utf8.DecodeRune(bytes)
    bytes = bytes[w:]
    single = get(int(r))
    // 中文字符判断
    tempStr = string(r)
    if len(single) == 0 {
      nohans += tempStr
    } else {
      if len(nohans) > 0 {
        retArr = append(retArr, nohans)
        nohans = ""
      }
      retArr = append(retArr, this.perStr(single))
    }
  }
  //处理末尾非中文的字符串
  if len(nohans) > 0 {
    retArr = append(retArr, nohans)
  }
  return retArr
}

func NewPy() *options {
  return &options{false, 0, false}
}
