package main

import (
  "encoding/json"
  "fmt"
  "os"
  "errors"

//"io/ioutil"
  "io/ioutil"
)

func main() {
  args := os.Args[1:]

  if len(args) < 1 || len(args) > 1 {
    panic(errors.New("You need to provide a JSON file (key-value, flat) as the only argument."))
  }

  jsonFile := args[0]
  fmt.Println(jsonFile)

  pages := readJsonFile(jsonFile)
  //fmt.Println(pages)

  for k, p := range pages {
    fmt.Printf("%s: %s\n" , k, p)
  }


  //byt := []byte(`{"key1":"test","key2":"test2"}`)
  //
  //var dat map[string]interface{}
  //
  //if err := json.Unmarshal(byt, &dat); err != nil {
  //  panic(err)
  //}
  //fmt.Println(dat)
}

func readJsonFile(filePath string)map[string]interface{} {
  raw, err := ioutil.ReadFile(filePath)

  if err != nil {
    panic(err)
  }

  var c map[string]interface{}
  json.Unmarshal(raw, &c)
  return c
}

func panic(e error) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "ERROR: %v\n", e)
    os.Exit(1)
  }
}
