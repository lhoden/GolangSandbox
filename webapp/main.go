package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/thedevsaddam/renderer"
)

/* This Go web application uses Go routines, channels, and synchronization
in order to maximize performance */

var wg sync.WaitGroup

// Renderer variable used to render all of the html views
var rnd *renderer.Render

// Template variable used to pass data to all the html views
//var tpl *template.Template

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Locations  []string `xml:"url>loc"`
	LastMod    []string `xml:"url>lastmod"`
	ChangeFreq []string `xml:"url>changefreq"`
	Title      []string
}

type NewsMap struct {
	LastMod  string
	Location string
	Title    string
	// The key of the map will be the title

}

// Using this value declaration struct to pass values to the view
// In order to pass multiple values, you have to use a struct like this
type NewsAggPage struct {
	Header string
	News   map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random about stuff here")
}

// Go routine function
func newsRoutine(c chan News, Location string) {
	defer wg.Done()
	var n News
	resp, _ := http.Get(strings.TrimSpace(Location))
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()
	var size int = len(n.Locations)
	if size >= 1 {
		size = size - 1
	}
	// consider adding wait group here!
	title := ExampleScrape(n.Locations[size])
	fmt.Println(title)
	n.Title = append(n.Title, "nothing")

	c <- n // sending the value of n over to the channel

}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	//Variables used to parse the XML returned from the HTTP request
	var s SitemapIndex
	//var title string
	//Making HTTP request
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsMap := make(map[string]NewsMap)
	resp.Body.Close()
	queue := make(chan News, 500) // buffer is 30, but you could put anyting. It just needs to be bigger than
	// the number we're using
	//fmt.Println(s.Locations) would print all of these sitemap locations in the console

	//range allows us to iterate over our own data structures
	for _, Location := range s.Locations { // basically saying for the entire range
		wg.Add(1)                       // want to add a wait group for each go routine to incorporate synchronization
		go newsRoutine(queue, Location) //passing the channel and the Location object to the go routine

	}

	wg.Wait()
	close(queue)

	for elem := range queue { // iterating over the channel (for each of the elements in the channel)
		for idx, _ := range elem.LastMod { // range is calculated based on the # of elements there are
			//title = "nothing"
			//title = ExampleScrape(elem.Locations[idx])
			newsMap[elem.Locations[idx]] = NewsMap{elem.LastMod[idx][0:10], elem.Locations[idx], "nothing"} //newsmap works as expected
			/* the [0:10] appended to the elem.LastMod[idx] here is just syntax for how you take substrings in Go
			Instead of substrings, we do "slices" of strings. Just like substring though, the first index is 0 and
			the last index included in the slice does not count. It's just a syntax thing and works the same way. */

		}
	}

	p := NewsAggPage{Header: "Amazing News Aggregator", News: newsMap}
	rnd.HTML(w, http.StatusOK, "agg", p)
}

// [5]type == array
// []type == slice

func main() {

	//Website page deployment
	http.HandleFunc("/", indexHandler)       /* creates a registration for the index page */
	http.HandleFunc("/about/", aboutHandler) /* creates a registration for another page */
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)

}

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "src/github.com/lhoden/webapp/tpl/*.html",
	}

	rnd = renderer.New(opts)
	//tpl = template.Must(template.ParseFiles("src/github.com/lhoden/webapp/tpl/agg.html"))
}

// Scrapes the webpage to locate the title
func ExampleScrape(address string) string {
	title := ""
	// Request the HTML page.
	// res, err := http.Get(address)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// if res.StatusCode != 200 {
	// 	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// }

	// // Load the HTML document
	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	doc, err := goquery.NewDocument(address)
	if err != nil {
		panic("meh")
	}

	// Find the review items
	doc.Find("head").Each(func(i int, s *goquery.Selection) {
		title = s.Find("title").Text()
	})
	return title
}
