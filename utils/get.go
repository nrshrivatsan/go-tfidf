
package main

import (
    "fmt"
    "log"
    "github.com/PuerkitoBio/goquery"
  
)

func main() {
	url := "http://en.wikipedia.org/wiki/Drug_delivery_systems"
	selector := "a"
	scrape(url,selector)
}

func scrape(url,selector string) {
    var doc *goquery.Document
    var e error

    if doc, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
   
    doc.Find(selector).Each(func(i int, s *goquery.Selection) {
    	fmt.Println(s.Text())
    	fmt.Println(s.Attr("href"))    	
        
    })
    
}
