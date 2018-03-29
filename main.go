package main

import "fmt"
import "github.com/ricardogama/git-rebasing-is-fun/src"

func main() {
  mul := src.Mul(3, 2)
  sum := src.Sum(3, 6)
  div := src.Div(8, 2)

  fmt.Println("mul: ", mul)
  fmt.Println("div: ", div)
  fmt.Println("sum: ", sum)
}
