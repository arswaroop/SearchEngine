package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	url string
	doc *goquery.Document
}

func NewScraper(u string) *Scraper {

	//if string is not http, return nil
	if !strings.HasPrefix(u, "http") {
		return nil
	}
	//capture response
	response, err := http.Get(u)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer response.Body.Close()

	//building document variable using goqeury build document.
	d, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//return url and the document
	return &Scraper{
		url: u,
		doc: d,
	}

}

// to return body of the document
func (s *Scraper) Body() string {
	body := s.doc.Find("body").Text()
	// Remove leading/ending white spaces
	body = strings.TrimSpace(body)

	return body
}

//if the scraper has links
func (s *Scraper) buildLink(href string) string {
	var link string

	if strings.HasPrefix(href, "/") {
		link = strings.Join([]string{s.url, href}, "")
	} else {
		link = href
	}

	link = strings.TrimRight(link, "/")
	link = strings.TrimRight(link, ":")

	return link
}

//function to return all the links that are available in the website.
func (s *Scraper) Links() []string {
	links := make([]string, 0)
	var link string

	s.doc.Find("body a").Each(func(index int, item *goquery.Selection) {
		link = ""

		linkTag := item
		href, _ := linkTag.Attr("href")

		if !strings.HasPrefix(href, "#") && !strings.HasPrefix(href, "javascript") {
			link = s.buildLink(href)
			if link != "" {
				links = append(links, link)
			}
		}
	})

	return links
}
