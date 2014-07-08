//I am not a genius. I seek help. This is a cargo-cult code + hacks. 
//Source http://stackoverflow.com/a/18479916
package main

import (
  "bufio"
  "fmt"
  "log"
  "os"
  "strings"
  "./utils"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) (*map[string]uint, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  dict := make(map[string]uint)
  var lines []string
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
    for _, line := range lines {
      //TODO Start indexing
      words := strings.Split(line," ")
      for _  ,word := range words{
        if dict[word] >0 {
            dict[word] += 1;
          }else{
            dict[word] = 1;
          }
      }
    }
  } 
    
  return &dict, scanner.Err()
}



func main() {
  queryString := "Neutron"
  
  fmt.Println(utils.SearchURLPrefix)
  links := utils.Scrape(utils.SearchURLPrefix+queryString,utils.Selector)   
  for k,v := range links{
    fmt.Println(k,utils.GetInfoCardText(v,utils.InfoCardSelector))

  }
  dict, err := readLines("foo.txt")
  if err != nil {
    log.Fatalf("readLines: %s", err)
  }else{
    fmt.Println(*dict)
  }

}