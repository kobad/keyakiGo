package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

var baseURL = "http://www.keyakizaka46.com/"
var t = time.Now()

const layout = "2006-01-02"

func main() {

	if err := os.Mkdir(t.Format(layout), 0777); err != nil {
		fmt.Println(err)
	}
	//targetURL := "http://www.keyakizaka46.com/mob/news/diarShw.php?site=k46o&ima=0000&cd=member"
	targetURL := "http://www.keyakizaka46.com/mob/news/diarKiji.php?site=k46o&ima=0000&cd=member&ct=21"
	doc, err := goquery.NewDocument(targetURL)
	fmt.Println("Go-keyaki")
	if err != nil {
		fmt.Println("url scrapping failed")
	}
	fmt.Println("Start Scrapping")

	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		r := regexp.MustCompile(`diarKiji.php`)
		if r.MatchString(url) {
		}
	})

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		dataURL, _ := s.Attr("data-url")
		imgURL, _ := s.Attr("src")
		url := baseURL + dataURL + imgURL
		saveIMG(url, i)
	})
}

func saveIMG(url string, i int) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf(t.Format(layout)+"/keyaki%d.jpg", i))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, response.Body)
}
