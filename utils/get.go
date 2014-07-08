
package utils

import (
    "log"
    "github.com/PuerkitoBio/goquery"
    "strings"    
)

const baseUrl string = "http://en.wikipedia.org"
const SearchURLPrefix string = "http://en.wikipedia.org/wiki/"
const Selector string = "#mw-content-text p a"
const imageSelector string = "#mw-content-text table.infobox tbody tr td a img"
const InfoCardSelector string = ".infobox"

func Scrape(url,selector string) (map[string]string) {
    var doc *goquery.Document
    var e error    
    if doc, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
    urlMap := make(map[string]string)

    doc.Find(selector).Each(func(i int, s *goquery.Selection) {
        if s.Text() != "" {
            link,_ := s.Attr("href")                 
            if !strings.HasPrefix(link,"//") && strings.HasPrefix(link,"/wiki"){
            
            // link="http:"+link   
                urlMap[s.Text()] = baseUrl+link
            }
            
        }       
    })

    return urlMap
}

func GetInfoCardText(url,selector string) string {
    var doc *goquery.Document
    var e error
    var infoCardText string
    if doc, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
   
    doc.Find(InfoCardSelector).Each(func(i int, s *goquery.Selection) {
        infoCardText = s.Text() 
        
    });
    return infoCardText
}
