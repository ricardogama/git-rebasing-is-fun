package main

import "fmt"
import "github.com/ricardogama/git-rebasing-is-fun/src"

func main() {
  sum := src.Sum(3, 6)
  div := src.Div(8, 2)
  mul := src.Mul(3, 2)

  fmt.Println("div: ", div)
  fmt.Println("sum: ", sum)
  fmt.Println("mul: ", mul)
}
