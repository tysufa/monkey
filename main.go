package main

import (
  "fmt"
  "os"
  "os/user"
  "qfa/repl"
)

func main() {
  user, err := user.Current()

  if err != nil {
    panic(err)
  }

  fmt.Println("Hello %s this is THE qfa, the best programming language ever made !\n", user.Username)
  fmt.Println("You now have the privilege to type commands in the best programing blablabla\n")
  repl.Start(os.Stdin, os.Stdout)
}
