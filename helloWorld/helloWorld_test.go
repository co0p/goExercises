package helloWorld

import "testing"

// If I tell the program my name is Alice, it will greet me by saying
// "Hello, Alice!".
func TestGreeting_name(t *testing.T)  {
  greeting := Greeting("Alice")
  if greeting != "Hello, Alice!" {
    t.Error("Expected 'Hello, Alice!', got ", greeting)
  }
}

//  I neglect to give it my name, it will greet me by saying
// "Hello, World!"
func TestGreeting_noname(t *testing.T)  {
  greeting := Greeting("")
  if greeting != "Hello, World!" {
    t.Error("Expected 'Hello, World!', got ", greeting)
  }
}
