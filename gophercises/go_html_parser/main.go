package main

import (
	"flag"
	"fmt"
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

	// Simple loop from the package's doc
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

}
