package types

import (
    "net/http"
    "sync"
)





type Ctx struct {
    BaseUrl string
    ImgLinks chan string
    Client *http.Client
    Wg sync.WaitGroup
    VisitedUrl map[string]bool
    DownloadedImgs map[string]bool
}



type Options struct {
	R   bool
	L   uint
	P   string
	Url string
}
