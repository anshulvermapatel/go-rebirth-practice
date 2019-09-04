package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

/*
func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page")
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}

*/

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Location []string `xml:"url>loc"`
	Titles   []string `xml:"url>news>title"`
	Keywords []string `xml:"url>news>keywords"`
}

type NewsAgg struct {
	Title string
	News  map[string]NewsMap
}

type NewsMap struct {
	Location string
	Keyword  string
}

var start time.Time
var wg sync.WaitGroup

func newsagg_handler(w http.ResponseWriter, r *http.Request) {
	start = time.Now()

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.washingtonpost.com/news-sitemaps/index.xml", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Accept-Language", "en-US")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}

	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	newsMap := make(map[string]NewsMap)
	mapChan := make(chan News, 30)
	for _, location := range s.Locations {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			//fmt.Printf("%d", k)
			l = strings.TrimSpace(l)
			//fmt.Printf("For - %s\n", location)
			var n News

			client := &http.Client{}
			req, err := http.NewRequest("GET", l, nil)
			if err != nil {
				log.Fatal(err)
				return
			}
			req.Header.Set("Connection", "keep-alive")
			req.Header.Set("Accept-Language", "en-US")
			req.Header.Set("User-Agent", "Mozilla/5.0")
			resp, err := client.Do(req)

			//resp, err := http.Get(l)
			if err != nil {
				log.Fatal(err)
			}
			bytes, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()
			xml.Unmarshal(bytes, &n)
			/*
				for i, title := range n.Titles {
					newsMap[title] = NewsMap{n.Location[i], n.Keywords[i]}
				}
			*/
			newsToChan(mapChan, n)
		}(location)
	}
	wg.Wait()
	close(mapChan)

	for channel := range mapChan {
		news := channel
		for i, title := range news.Titles {
			newsMap[title] = NewsMap{news.Location[i], news.Keywords[i]}
		}
	}

	newsAgg := NewsAgg{Title: "This is my News Aggregator", News: newsMap}

	t, err := template.ParseFiles("webapp_template.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.Execute(w, newsAgg))
	fmt.Println(time.Since(start))
}

func newsToChan(c chan News, n News) {
	c <- n
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Hello this is my new aggregator </h1> <br> <h4><a href='http://localhost:8000/newsagg' _target='blank' > Click Here </a></h4>")
}

func main() {

	http.HandleFunc("/", index_handler)
	http.HandleFunc("/newsagg", newsagg_handler)
	//http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)

	/*
		for i, nm := range newsMap {
			fmt.Println(i, "------------")
			fmt.Println("Location -", nm.Location)
			fmt.Println("Keyword -", nm.Keyword)
		}
	*/

	// For iterating over the stuff, we will also use maps
	/*
		fmt.Println("Location are -")
		for _, location := range n.Location {
			fmt.Println(location)
		}
		fmt.Println("Titles are -")
		for _, title := range n.Titles {
			fmt.Println(title)
		}
		fmt.Println("Keywords are -")
		for _, keyword := range n.Keywords {
			fmt.Println(keyword)
		}
	*/
}
