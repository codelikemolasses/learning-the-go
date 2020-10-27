package main

import (
  b64 "encoding/base64"
  "fmt"
)

func main() {
  data := "abc123!?$*&()'-=@~"
  
  sEnc := b64.StdEncoding.EncodeToString([]byte(data))
  fmt.Println(sEnc)
  
  sDec, _ := b64.StdEncoding.DecodeString("ZmxhZ3toM2xsMF93MHJsZF80c190MGxkX2J5XzRfdGgwdXM0bmRfYzBtbTF0c30=")
  fmt.Println(string(sDec))
  fmt.Println()
  
  uEnc := b64.URLEncoding.EncodeToString([]byte(data))
  fmt.Println(uEnc)
  uDec, _ := b64.URLEncoding.DecodeString(uEnc)
  fmt.Println(string(uDec))
}
