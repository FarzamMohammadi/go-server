package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	ID          int    `json:"id"`
	Description string `json: "description`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
	)
	collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
		factID, err := strconv.Atoi(element.Attr("id"))

		if err != nil {
			log.Println("Could not get ID")
		}

		factDesc := element.Text

		fact := Fact{
			ID:          factID,
			Description: factDesc,
		}

		allFacts = append(allFacts, fact)
	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting: ", request.URL.String())
	})

	collector.Visit("https://www.factretriever.com/rhino-facts")

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allFacts)

}
