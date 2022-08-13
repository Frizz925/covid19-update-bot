package mhlw

type RDF struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"date"`
}
