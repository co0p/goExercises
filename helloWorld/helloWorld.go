package helloWorld

func Greeting(name string) string {

  if (len(name) > 0) {
    return "Hello, " + name + "!"
  }
  return "Hello, World!"
}
