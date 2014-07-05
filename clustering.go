// clustering.go
package main

import (
  "fmt"
  "./utils"
  "strings"
  "math"
)

// Returns a map where the key is a word and the int is the number
// of times that word appears in the set of documents
// 
// Providing a threshold (1.0 >= x > 0.0) will return only the words
// that appear in all the documents (x*100)% of the time
func termFrequency(recordArray [][]string, threshold float64) (m map[string] int, err error) {
  saveMap := make(map[string] map[string] int)
  for _, record := range recordArray {
    url := record[0]
    
    if _, ok := saveMap[url]; ok {
      continue
    }

    words := utils.LowercaseWords(strings.Fields(record[2]))
  
    for i := range words {
      w, err := utils.RemoveNonAlphaNumeric(words[i])
      if err != nil {
        continue
      } else {
        words[i] = w
      }
    }

    words, err = utils.RemoveStopwords(words)
    if err != nil {
      return nil, err
    }

    saveMap[url] = utils.WordFrequency(words)
  }

  documentFrequencyMap := make(map[string] int)

  for _, wordCountMap := range saveMap {
    for word := range wordCountMap {
      if _, ok := documentFrequencyMap[word]; ok {
        documentFrequencyMap[word]++
      } else {
        documentFrequencyMap[word] = 1
      }
    }
  }

  if threshold != 0.0 {
    for word, value := range documentFrequencyMap {
      if float64(value)/float64(len(saveMap)) < threshold {
        delete(documentFrequencyMap, word)
      }
    }
  }
  return documentFrequencyMap, nil
}

// Inverse Document Frequency
func inverseDocumentFrequency(recordArray [][]string) (m map[string] float64, err error) {
  d := float64(len(recordArray))
  
  wordCountMap := make(map[string] int)
  for _, record := range recordArray {
    words := utils.LowercaseWords(strings.Fields(record[2]))
  
    for i := range words {
      w, err := utils.RemoveNonAlphaNumeric(words[i])
      if err != nil {
        continue
      } else {
        words[i] = w
      }
    }

    words, err = utils.RemoveStopwords(words)
    if err != nil {
      return nil, err
    }

    words = utils.RemoveDuplicates(words)

    for _, word := range words {
      if _, ok := wordCountMap[word]; ok {
        wordCountMap[word]++
      } else {
        wordCountMap[word] = 1
      }
    }
  }

  idfMap := make(map[string] float64)
  for word, value := range wordCountMap {
    idfMap[word] = math.Log(d/float64(value))
  }
  return idfMap, nil
}

// Term Frequency-Inverse Document Frequency (TF-IDF)
func termFrequencyInverseDocumentFrequency(fileName string) (m map[string] float64, err error) {
  recordArray, err := utils.ReadRecords(fileName)
  if err != nil {
    return nil, err
  }

  tfidfMap := make(map[string] float64)

  tf, err := termFrequency(recordArray, 0.0)
  if err != nil {
    return nil, err
  }
  
  idf, err := inverseDocumentFrequency(recordArray)
  if err != nil {
    return nil, err
  }

  for word, docFreq := range idf {
    tfidfMap[word] = float64(tf[word]) * docFreq
  }
  return tfidfMap, nil
}

func main() {
  tfidfMap, err := termFrequencyInverseDocumentFrequency("pocket.csv")
  if err != nil {
    fmt.Println(err)
  } else {
    for k,v := range tfidfMap {
      fmt.Println(k , "->", v )
    }
  }
}