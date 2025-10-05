package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
    "log"
	"os"
    "io"
)

type Options struct {
	R   bool
	L   uint
	P   string
	Url string
}

var Opt Options

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage:\n%s [-rlp] URL\nOptions:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&Opt.R, "r", false, "A `boolean` for recursively downloads the images in a URL")
	flag.UintVar(&Opt.L, "l", 5, "the maximum depth level of the recursive download default value is 5")
	flag.StringVar(&Opt.P, "p", "data", "the path where the downloaded files will be saved default value is ./data")

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
	}
	Opt.Url = flag.Arg(0)
}

// client http
func NewClient() *http.Client {
	return &http.Client{}
}

func NewRequest(url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        log.Fatal(err)
    }
	userAgents := []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.3",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.3",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.3",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36 Trailer/93.3.8652.5"}

	req.Header.Set("User-Agent", userAgents[int(rand.Int31())%len(userAgents)])
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/*,*/*;q=0.8")

	return req
}

func main() {

    client := NewClient()

    req := NewRequest(Opt.Url)
    resp , err := client.Do(req)

    if err != nil {
        log.Fatal(err)
    }
    bodyBytes, err := io.ReadAll(resp.Body)

    fmt.Println("resp")
    fmt.Println(string(bodyBytes))

}
