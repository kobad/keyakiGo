package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var memberList = map[string]string{
	"ishimori":  "01",
	"zu-min":    "02",
	"uemu-":     "03",
	"ozeki":     "04",
	"odanana":   "05",
	"koike":     "06",
	"kobayashi": "07",
	"fu-chan":   "08",
	"si-chan":   "09",
	"mona":      "10",
	"yukka":     "11",
	"kuritarou": "12",
	"na-ko":     "13",
	"habu":      "14",
	"aoi":       "15",
	"techi":     "17",
	"moriya":    "18",
	"yonesan":   "19",
	"berika":    "20",
	"berisa":    "21",
	"neru":      "22"}
var baseURL = "http://www.keyakizaka46.com"
var t = time.Now()
var targetName = ""

const layout = "2006-01-01"

func main() {
	fmt.Println("Go-keyaki")
	flag.Parse()
	targetName = flag.Arg(0)
	number := memberList[flag.Arg(0)]
	targetURL := "http://www.keyakizaka46.com/mob/news/diarKiji.php?site=k46o&ima=0000&cd=member&ct=" + number

	doc, err := goquery.NewDocument(targetURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, "url scrapping failed")
		os.Exit(1)
	}

	if err := os.Mkdir(targetName+t.Format(layout), 0777); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	doc.Find("article").Each(func(i int, s *goquery.Selection) {
		s.Children().Find("img").Each(func(_ int, arg1 *goquery.Selection) {
			url, _ := arg1.Attr("src")
			imgURL := baseURL + url
			fmt.Println("Get: " + imgURL)
			if err := saveIMG(imgURL, i); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		})
	})
}

func saveIMG(url string, i int) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf(targetName+t.Format(layout)+"/%s%d.jpg", targetName, i))
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, response.Body)

	return nil
}
