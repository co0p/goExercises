package main

import "fmt"
import "goExercises/helloWorld"

func main() {
  var name string
  fmt.Println("What's your name?")
  fmt.Scanln(&name)
  fmt.Println(helloWorld.Greeting(name))
}
