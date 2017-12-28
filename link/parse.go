package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a Link (<a href="...">) in an HTML
// document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and return
// a slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	linkNodes := dfs(doc)
	links := processLinkNodes(linkNodes)
	return links, nil
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func processLinkNodes(linkNodes []*html.Node) []Link {
	var links []Link
	for _, node := range linkNodes {
		link := Link{}
		href := ""
		for _, attribute := range node.Attr {
			if strings.ToLower(attribute.Key) == "href" {
				href = attribute.Val
				break
			}
		}
		link.Href = href
		link.Text = text(node)
		links = append(links, link)
	}
	return links
}

func dfs(n *html.Node) []*html.Node {
	if n == nil {
		return nil
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var linkNodes []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodes = append(linkNodes, dfs(c)...)
	}
	return linkNodes
}
