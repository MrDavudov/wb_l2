package main

import (
	"fmt"
	"math/big"
)

func Zeros(n int) (zero int) {
  if n == 0 {
    return 0
  }
  var res int64 = big.NewInt(1)
  for i:=1;i<=n;i++ {
    big.NewInt(res) = big.NewInt(res) * big.NewInt(i)
  }
  fmt.Println(res)

  str := fmt.Sprintf("%d", res)
  for i:=len(str)-1;i>0;i-- {
    if "0" == string(str[i]) {
      zero++
    } else {
      return zero
    }
  }
  return zero
}

func main() {
	fmt.Println(Zeros(26))
}