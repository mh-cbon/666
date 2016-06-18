package main

import (
  "fmt"
  "os"
  "os/exec"

  "github.com/fatih/color"
)

func main () {
  args := os.Args

  if len(args)<2 {
    fmt.Fprintln(os.Stderr, "Missing arguments")
    os.Exit(1)
  }

  bin := args[1]
  cmdArgs := make([]string, 0)
  if len(args)>1 {
    cmdArgs = args[2:]
  }

  oCmd := exec.Command(bin, cmdArgs...)
	oCmd.Stdout = os.Stdout
	oCmd.Stderr = os.Stderr
  err := oCmd.Run()

  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    fmt.Fprintln(os.Stderr, color.RedString(" ✘ Failed"))
    os.Exit(1)
  }
  fmt.Println(color.GreenString(" ✔ Success"))

  }
