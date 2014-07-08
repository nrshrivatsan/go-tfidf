// utils.go
// Implements utilities for the primary clustering methods
package utils

import (
  "encoding/csv"
  "os"
  "strings"
  "regexp"
)

// func ReadRecords(fileName string) (s [][]string, err error) {
//   file, err := os.Open(fileName)
//   if err != nil {
//     return nil, err
//   }
//   defer file.Close()

//   reader := csv.NewReader(file)
//   reader.TrailingComma = true

//   recordArray, err := reader.ReadAll()
//   if err != nil {
//     return nil, err
//   }
//   return recordArray, nil
// }

func RemoveDuplicates(words []string) []string {
  newWords := make([]string, len(words))
  for _, word := range words {
    contains := false
    for _, w := range newWords {
      if strings.EqualFold(w, word) {
        contains = true
        break
      }
    }
    if !contains {
      newWords = append(newWords, word)
    }
  }
  return newWords
}

func WordFrequency(words []string) (s map[string] int) {
  wordCountMap := make(map[string] int)
  for _, word := range words {
    if strings.EqualFold(word, "") {
      continue
    }
    if _, ok := wordCountMap[word]; ok {
      wordCountMap[word]++
    } else {
      wordCountMap[word] = 1
    }
  }
  return wordCountMap
}

func RemoveStopwords(words []string) (s []string, err error) {
  stopwords, err := stopwords()
  if err != nil {
    return nil, err
  }

  for _, word := range stopwords { 
    in, i := wordInList(word, words);
    for in {
      words = words[:i+copy(words[i:], words[i+1:])]
      in, i = wordInList(word, words)
    }
  }
  return words, nil
}

func LowercaseWords(words []string) []string {
  for i := range words {
    words[i] = strings.ToLower(words[i])
  }
  return words
}

func RemoveNonAlphaNumeric(s string) (str string, err error) {
  r, err := regexp.Compile("[^\\w]|[-+]?\\d+")
  if err != nil {
    return "", err
  }
  return r.ReplaceAllString(s, ""), nil
}

func stopwords() (s []string, err error) {
  file, err := os.Open("stopwords.csv")
  if err != nil {
    return nil, err
  }
  defer file.Close()
  record, err := csv.NewReader(file).Read()
  if err != nil {
    return nil, err
  }
  return record, nil
}

func wordInList(a string, list []string) (b bool, index int) {
  for i, b := range list {
    if b == a {
      return true, i
    }
  }
  return false, -1
}
