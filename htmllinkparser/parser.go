package htmllinkparser

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/html"
)

func parseNode(n *html.Node) {
	if n.Parent != nil && n.Type == html.ElementNode && n.Parent.Type == html.ElementNode && n.Parent.Data == "a" {
		fmt.Printf("%s\n", n.FirstChild.Data)
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		if n.FirstChild.Type == html.TextNode {
			fmt.Printf("%s\n", n.FirstChild.Data)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseNode(c)
	}
}

func ParseHTMLandExecute(r io.Reader) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal("Fatal error while parsing HTML")
	}
	linksList := getLinks(doc)
	fmt.Println(linksList)
	// texts := grabTextFromLinks(linksList)
	// fmt.Printf("%+v\n", texts)
}

func getLinks(n *html.Node) []string {
	var l []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" && a.Val != "#" {
				l = append(l, a.Val)
			}
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		l = append(l, getLinks(c)...)
	}
	return l
}

func grabTextFromLinks(l []*html.Node) []string {
	text := []string{}
	for _, link := range l {
		text = append(text, grabTextFromNode(link))
	}
	return text
}

func grabTextFromNode(n *html.Node) string {
	str := ""
	if n.Type == html.TextNode {
		str += n.Data
		fmt.Println(str)
	} else {
		return ""
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		str += grabTextFromNode(c)
	}
	return str
}
