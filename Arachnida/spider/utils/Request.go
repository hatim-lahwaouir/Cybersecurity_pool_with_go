package utils

import (
	"Arachnida/spider/types"
	"golang.org/x/net/html"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

var userAgents = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.3",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.3",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.3",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Trailer/93.3.8652.5"}

func NewRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", userAgents[int(rand.Int31())%len(userAgents)])
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")

	return req
}

func HandleRequest(ctx *types.Ctx, node *types.UrlNode) *types.UrlNode {

	var (
		process func(*html.Node)
		newUrl  string
	)

	req := NewRequest(node.Url)
	resp, err := ctx.Client.Do(req)

	if err != nil {
		return nil
	}
	// c := resp.Cookies()
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	process = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, ele := range n.Attr {

				if ele.Key == "href" && (strings.HasPrefix(ele.Val, "/") || strings.HasPrefix(ele.Val, ctx.BaseUrl)) {
                    if _, ok := ctx.VisitedUrl[ele.Val]; !ok {
                        newUrl, err = url.JoinPath(ctx.BaseUrl, ele.Val)
                        if err == nil {
                            node.C = append(node.C, &types.UrlNode{Url: newUrl})
                        }
                        ctx.VisitedUrl[ele.Val] = true
                    }
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, ele := range n.Attr {
				if ele.Key == "src" {
                    if _, ok := ctx.DownloadedImgs[ele.Val]; !ok{
					     ctx.ImgLinks <- ele.Val
                         ctx.DownloadedImgs[ele.Val] = true
                    }
				}
			}
		}

		// traverse the child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			process(c)
		}
	}

	process(doc)
	return node
}
