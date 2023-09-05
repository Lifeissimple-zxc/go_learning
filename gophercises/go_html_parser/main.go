package main

import (
	"flag"
	"fmt"
	mycrawler "go_html_parser/my_crawler"
	"os"

	"golang.org/x/net/html"
)

func main() {
	fPtr := flag.String("html", "./html_examples/ex1.html", "Path to html for parsing")
	flag.Parse()

	fmt.Printf("Starting with %s as input file\n", *fPtr)

	file, err := os.Open(*fPtr)
	if err != nil {
		fmt.Printf("Err opening %s. Details: %v\n", *fPtr, err)
		os.Exit(1)
	}

	// Html parsing
	doc, err := html.Parse(file)
	if err != nil {
		fmt.Printf("Parsing Html failed: %v\n", err)
		os.Exit(1)
	}

	var a []mycrawler.Link
	err = HtmlToLinkSlice(doc, &a)
	if err != nil {
		fmt.Printf("Err converting html to link slice of As: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", a)

}

func HtmlToLinkSlice(n *html.Node, container *[]mycrawler.Link) error {
	// Parse Node
	if n.Type == html.ElementNode && n.Data == "a" {
		fmt.Println("Found our guy")
		l := mycrawler.Link{}
		// Get Href
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.Href = a.Val
				break
			}
		}
		// Get Inner HTML
		// This needs to be a recursive function because of the example provided
		l.Text = n.FirstChild.Data

		// Finally append to our slice
		*container = append(*container, l)
	}
	// Parse Children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// fmt.Printf("Working on node: %+v\n", c)
		HtmlToLinkSlice(c, container)
	}
	return nil
}
