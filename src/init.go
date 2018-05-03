package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "io/ioutil"
  "encoding/json"
  "github.com/aws/aws-sdk-go/aws"
  "flag"
)

func main() {
  showHelp := flag.Bool("help", false, "Displays help")
  flag.Parse()

  if *showHelp {
    helpPrompt()
    os.Exit(0)
  }

  args := os.Args[1:]
  var fileName string

  if len(args) > 0 {
    fileName = args[0]
  } else {
    fileName = ""
  }

  jsonMap := getJson(fileName)

  p := []*string{}
  for _, v := range jsonMap {
    p = append(p, aws.String(v))
  }

  results := SSMGet(p)

  var finalMap map[string]string
  finalMap = make(map[string]string)

  for k, v := range jsonMap {
    // check if results[v] exists before assign to fail if there are errors
    finalMap[k] = results[v]
  }

  outputString, _ := json.Marshal(finalMap)

  fmt.Print(string(outputString))
}

func helpPrompt() {
  fmt.Printf("USAGE: go-ssm-params [file.json].\n" +
    "It will try to read a valid flat JSON from stdin if no file passed.\n\n" +

    "Examples reading from stdin:\n" +
    "----------------------------\n" +
    "\t$ cat file.json | go-ssm-params\n" +
    "\t$ go-ssm-params < file.json")
}

func getJson(fileName string) map[string]string {
  // if empty filename, assumes it reads from stdin
  var raw []byte
  var jsonMap map[string]string

  if fileName == "" {
    // Reading from stdin
    raw = stdinRead()
  } else {
    // Reading from a file
    raw = fileRead(fileName)
  }

  if err := json.Unmarshal(raw, &jsonMap); err != nil {
    exitPanic(err)
  }

  return jsonMap
}

func stdinRead() []byte {
  var lines []string
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return []byte(strings.Join(lines, ""))
}

func fileRead(filePath string) []byte {
  raw, err := ioutil.ReadFile(filePath)

  if err != nil {
    exitPanic(err)
  }

  return raw
}

func exitPanic(e error) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "ERROR: %v.\n", e)
    os.Exit(1)
  }
}
