package mycrawler

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func (l *Link) NodeToLink(n *html.Node) {
	// Get Href
	href := getHREF(&n.Attr)
	if href == "" {
		fmt.Println("Did not find href")
	}
	l.Href = href

	// Get  Inner HTML
	innerHTML := getInnerHTML(n)
	if innerHTML == "" {
		fmt.Println("Did not find inner html")
	}
	l.Text = innerHTML

}

// TBD
func getHREF(attrs *[]html.Attribute) string {
	for _, a := range *attrs {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}

func getInnerHTML(n *html.Node) string {
	var sb strings.Builder

	var find func(*html.Node)
	find = func(n *html.Node) {
		if n.Type != html.ElementNode {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				s := strings.TrimSpace(c.Data)
				sb.WriteString(fmt.Sprintf("%s ", s))
			}
			find(c)
		}
	}

	find(n)

	return sb.String()

}
