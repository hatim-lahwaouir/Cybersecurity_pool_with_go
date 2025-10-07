package main

import (
	"Arachnida/src/types"
	"Arachnida/src/utils"
	"flag"
	"fmt"
	"net/http"
	"net/url"
    "time"
	"os"
)

var Opt types.Options
var ctx types.Ctx

func init() {
	// parsing args
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

func init() {
	// creating directory
	err := os.MkdirAll(Opt.P, 0750)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: Error creating data folder %s\n", os.Args[0], err.Error())
		os.Exit(1)
	}
}

func init() {
	// initializing variables and creating threads for downloading images
	ctx.ImgLinks = make(chan string, 300)
	ctx.Client = &http.Client{}

	n_goroutine := 10
	for i := 0; i < n_goroutine; i++ {
		go func() {
			defer ctx.Wg.Done()
			utils.DownloadImg(&ctx, Opt.P)
		}()
		ctx.Wg.Add(1)
	}
}

func main() {
	var (
		nodes      []*types.UrlNode
		childNodes []*types.UrlNode
        ticker  *time.Ticker
	)

    ticker = time.NewTicker(300 * time.Millisecond)
    defer ticker.Stop()
	nodes = append(nodes, &types.UrlNode{Url: Opt.Url})



	for Opt.L > 0 {
        <- ticker.C
		for i := 0; i < len(nodes); i++ {
			utils.HandleRequest(&ctx, nodes[i])
			childNodes = append(childNodes, nodes[i].C...)
		}
		nodes = childNodes
		Opt.L -= 1
	}
	close(ctx.ImgLinks)
	ctx.Wg.Wait()
}
