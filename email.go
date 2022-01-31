package main

import (
  "os"
  "fmt"
  "regexp"
  "encoding/csv"
  "github.com/gocolly/colly"
)

func main()  {
  fName := "data.csv"
  file,_ :=  os.Create(fName)
  defer file.Close()

  writer := csv.NewWriter(file)
  defer writer.Flush()

  c := colly.NewCollector()

  r := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)

  c.OnHTML("body", func(e *colly.HTMLElement) {
    //fmt.Println(e.Text)
    all := r.FindAllString(e.Text, -1)
    for _, g :=range all {
      writer.Write([]string {
        g,
      })
      fmt.Println(r.MatchString(e.Text))
    }
  })

  url := "https://i2crm.ru/"
  fmt.Printf("Parsing start \n")
  c.Visit(url)
  fmt.Printf("Parsing complete \n")
}
//gross := e.ChildText(“.text-muted:contains(‘Gross’) ~ span[name=nv]”)
//gross = e.ChildText(“[class=’text-muted text-small’] span:contains(‘$’)”)
//vote := e.ChildAttr(“span[name=nv]”, “data-value”)
//rating := e.ChildText(“[class=’ipl-rating-star small’] .ipl-rating-star__rating”)
//name := e.ChildText(“.lister-item-index ~ a”)
//number := e.ChildText(“.lister-item-index”)
