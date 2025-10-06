package utils

import (
	"log"
	"math/rand"
    "fmt"
    "os"
	"net/http"
	"net/url"
	"golang.org/x/net/html"
	"Arachnida/src/types"
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







func HandleRequest(client *http.Client, node *types.UrlNode, ctx *types.Ctx) *types.UrlNode {

    var (
	    process func(*html.Node)
        newUrl  string
    )

	req := NewRequest(node.Url)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: request failled error '%s'", os.Args[0], err.Error())
        return  nil
	}
	// c := resp.Cookies()
    defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: can't parse html in this url %s ", os.Args[0], node.Url)
		return nil
	}

	process = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
            for _, ele := range n.Attr {

                 if ele.Key == "href"  {
                     fmt.Println(ele.Val)
                }

                
                if ele.Key == "href" && (strings.HasPrefix(ele.Val, "/") || strings.HasPrefix(ele.Val, ctx.BaseUrl)){
                    newUrl, err = url.JoinPath(ctx.BaseUrl, ele.Val)
                    if err == nil {
                        node.C = append(node.C, &types.UrlNode{Url : newUrl })
                    }
                }
            }
		}
        if n.Type == html.ElementNode && n.Data == "img" {
            for _, ele := range n.Attr {
                if ele.Key == "href" {
                    fmt.Println(ele.Val) 
                }
            }
		}

		// traverse the child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			process(c)
		}
	}

    process(doc)
    fmt.Println("****************************************************")
    fmt.Println(node)
    return node
}
