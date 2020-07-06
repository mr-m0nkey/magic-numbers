package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	//TODO properly handle arguments
	magicNuber := os.Args[1]
	url := "https://nhentai.net/g/" + magicNuber

	title, details, err := getDetails(url)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Title")
	fmt.Printf(title)
	fmt.Println(details)
	fmt.Println(url)	
}

func getDetails(url string) (string, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}

	// Convert HTML into goquery document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}

	// Save each .post-title as a list
	titles := ""
	doc.Find(".title").Each(func(i int, s *goquery.Selection) {
		titles += s.Text() + "\n"
	})

	details := ""
	doc.Find(".tag-container").Each(func(i int, s *goquery.Selection) {
		details += s.Text()
	})


	return titles, details, nil
}

