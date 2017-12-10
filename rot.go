package main

import (
  "fmt"
  "strings"
)

type rot struct  {
  places    int
  letters   [26]string
}

func (r* rot) init() {
  r.letters[0] = "a"
  r.letters[1] = "b"
  r.letters[2] = "c"
  r.letters[3] = "d"
  r.letters[4] = "e"
  r.letters[5] = "f"
  r.letters[6] = "g"
  r.letters[7] = "h"
  r.letters[8] = "i"
  r.letters[9] = "j"
  r.letters[10] = "k"
  r.letters[11] = "l"
  r.letters[12] = "m"
  r.letters[13] = "n"
  r.letters[14] = "o"
  r.letters[15] = "p"
  r.letters[16] = "q"
  r.letters[17] = "r"
  r.letters[18] = "s"
  r.letters[19] = "t"
  r.letters[20] = "u"
  r.letters[21] = "v"
  r.letters[22] = "w"
  r.letters[23] = "x"
  r.letters[24] = "y"
  r.letters[25] = "z"
}

func (r* rot) decrypt(msg string) string {
  fmt.Println("In decrypt")
  
  decryptedBuffer := strings.Split(msg, "")
  
  if r.places == 0 && msg == "\n" {
    fmt.Println("msg is empty, returning empty string")
    return ""
  }
  
  for i := 1l i <= len(decryptedBuffer)
  
  fmt.Println("returning concatenated buffer")
  return strings.Join(decryptedBuffer, "")
}

func main() {
  msg := "Hey"
  r := rot {
    places: 14,
  }
  fmt.Println(r.decrypt(msg))
}