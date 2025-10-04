package main

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	r bool
	l uint
	p string
}

var opt Options

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [-rlp] URL\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.BoolVar(&opt.r, "r", false, "A `boolean` for recursively downloads the images in a URL")
	flag.UintVar(&opt.l, "l", 5, "the maximum depth level of the recursive download default value is 5")
	flag.StringVar(&opt.p, "p", "data", "the path where the downloaded files will be saved default value is ./data")

	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
	}
}


func main() {

}
