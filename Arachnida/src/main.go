package main

import (
	"Arachnida/src/types"
	"Arachnida/src/utils"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var Opt types.Options
var ctx types.Ctx

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [-rlp] URL\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.BoolVar(&Opt.R, "r", false, "A `boolean` for recursively downloads the images in a URL")
	flag.UintVar(&Opt.L, "l", 2, "the maximum depth level of the recursive download default value is 5")
	flag.StringVar(&Opt.P, "p", "data", "the path where the downloaded files will be saved default value is ./data")

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
	}
	Opt.Url = flag.Arg(0)

	urlInfo, err := url.ParseRequestURI(Opt.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Invalid url \n", os.Args[0])
		os.Exit(1)
	}
	// getting base url
	urlInfo.Path = ""
	urlInfo.RawQuery = ""
	urlInfo.Fragment = ""

	ctx.BaseUrl = urlInfo.String()
	fmt.Println(urlInfo.String())
}

// client http
func NewClient() *http.Client {
	return &http.Client{}
}

func main() {
	var (
		client *http.Client
		nodes   []*types.UrlNode
		childNodes []*types.UrlNode
	)

	client = NewClient()
	nodes = append (nodes, &types.UrlNode{Url: Opt.Url})

	for Opt.L > 0 {
        for i := 0; i  < len(nodes); i++{
		    utils.HandleRequest(client, nodes[i], &ctx)
            childNodes = append(nodes[i].C, childNodes)
        }
        nodes =childNodes
        Opt.L -= 1 
	}

}
