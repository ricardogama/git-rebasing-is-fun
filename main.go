package main

import "fmt"
import "github.com/ricardogama/git-rebasing-is-fun/src"

func main() {
  sum := src.Sum(3, 6)
  mul := src.Mul(3, 2)
  div := src.Div(8, 2)

  fmt.Println("mul: ", mul)
  fmt.Println("div: ", div)
  fmt.Println("sum: ", sum)
}
