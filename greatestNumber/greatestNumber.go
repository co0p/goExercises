package main

import "fmt"

// greatestNumber returns the greatest number from the given list of numbers
func greatestNumber(numbers ...int) (max int) {
  max = numbers[0]
  for _, number := range numbers {
    if number > max {
      max = number
    }
  }
  return
}

func main() {
  maxNumber := greatestNumber(-2,-5,-23)
  fmt.Println(maxNumber)
}
