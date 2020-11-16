package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/thedevsaddam/renderer"
)

// install third party packages from github with command: go get github.com/whatever/whatever
// in terminal
var washPostXML = []byte(`
<sitemapindex>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/politics.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/business.xml</loc>
	</sitemap>
	<sitemap>
		<loc>https://www.washingtonpost.com/sitemaps/opinions.xml</loc>
	</sitemap>
</sitemapindex>`)

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Locations  []string `xml:"url>loc"`
	LastMod    []string `xml:"url>lastmod"`
	ChangeFreq []string `xml:"url>changefreq"`
}

type NewsMap struct {
	ChangeFreq string
	Location   string
	// The key of the map will be the title

}

// Using this value declaration struct to pass values to the view
// In order to pass multiple values, you have to use a struct like this
type NewsAggPage struct {
	Title string
	News  map[string]NewsMap
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Random about stuff here")
}

func newsAggHandler(w http.ResponseWriter, r *http.Request) {
	//Variables used to parse the XML returned from the HTTP request
	var s SitemapIndex
	var n News
	//Making HTTP request
	resp, _ := http.Get("https://www.washingtonpost.com/sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &s)
	newsMap := make(map[string]NewsMap)
	//fmt.Println(s.Locations) would print all of these sitemap locations in the console

	//range allows us to iterate over our own data structures
	for _, Location := range s.Locations { // basically saying for the entire range
		resp, _ := http.Get(strings.TrimSpace(Location))
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)
		for idx, _ := range n.LastMod { // range is calculated based on the # of lastmod elements there are
			newsMap[n.LastMod[idx]] = NewsMap{n.ChangeFreq[idx], n.Locations[idx]} //newsmap works as expected
		}
		// for idx, data := range newsMap {
		// 	fmt.Println("\n\n\n\n\n", idx)
		// 	fmt.Println("\n", data.Keyword)
		// 	fmt.Println("\n", data.Location)
		// }
	}

	p := NewsAggPage{Title: "Amazing News Aggregator", News: newsMap}
	fmt.Fprint(w, p) //Returns everything we want! Yay!
	tpl := template.Must(template.ParseFiles("site.html"))
	if err != nil {
		//Send back error message
		http.Error(w, "Hey, request was bad!", http.StatusBadRequest)
	}
	tpl.Execute(w, p) // still doesn't seem to be obtaining the template
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
	opts := renderer.Options {
		ParseGlobPattern: "./templates/*.html"
	}

	rnd = renderer.New(opts)
}