package main

import "fmt"

func main () {
  var sum, prev_fib, current_fib, new_fib int64
  prev_fib = 1
  current_fib = 1
  new_fib = prev_fib + current_fib
  for new_fib < 4000000 {
    if new_fib % 2 == 0 {
      sum += new_fib
      fmt.Print(new_fib, ", ")
    }
    prev_fib = current_fib
    current_fib = new_fib
    new_fib = prev_fib + current_fib
  }
  fmt.Println("\nSum of even fibonacci terms below 4M: ", sum)
}
