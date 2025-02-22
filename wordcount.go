package main

import (
  "fmt"
  "os"
  //"bufio"
  "strings"
)

func main() {
  s,err := readImput()
  if err != nil {
    printErrors(err)
  }
  strSlice := splitStr(s)
  
  fmt.Println(len(strSlice))
}

func splitStr(s string) []string {
    s = strings.TrimSpace(s)
    if s == "" {
      return []string{}
    }
    return strings.SplitAfter(s, " ")
}

func readImput()(string, error){
  //reader := bufio.NewReader(os.Stdin)
  //str,err := reader.ReadString('\n')
  str := os.Args[1]
  return str, nil
}

func printErrors(err error) {
  fmt.Println(err)
}