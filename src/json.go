package main

import (
  "encoding/json"
  "fmt"
  "os"
  "errors"
  "io/ioutil"
)

func main() {
  args := os.Args[1:]

  if len(args) < 1 || len(args) > 1 {
    panic(errors.New("you need to provide a JSON file (key-value, flat) as the only argument"))
  }

  jsonFile := args[0]
  pages := readJsonFile(jsonFile)

  for k, p := range pages {
    fmt.Printf("%s: %s\n" , k, p)
  }
}

func readJsonFile(filePath string) map[string]string {
  raw, err := ioutil.ReadFile(filePath)

  if err != nil {
    panic(err)
  } else {
    fmt.Printf("Using `%s` as input file.\n\n", filePath)
  }

  var c map[string]string
  if err := json.Unmarshal(raw, &c); err != nil {
    panic(err)
  }

  return c
}

func panic(e error) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "ERROR: %v.\n", e)
    os.Exit(1)
  }
}
