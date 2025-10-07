package utils

import (
	"Arachnida/src/types"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
    "time"
)

var allowedMimeType = map[string]bool{"image/gif": true, "image/png": true, "image/jpeg": true, "image/bmp": true}

func DownloadImg(ctx *types.Ctx, dir string) {
	var (
		imgPath string
        req     *http.Request
        ticker  *time.Ticker
	)


    ticker = time.NewTicker(300 * time.Millisecond)
    defer ticker.Stop()

	for imgUrl := range ctx.ImgLinks {
        <- ticker.C
        if strings.HasPrefix(imgUrl, "//") {
            imgUrl = "https:" + imgUrl
        } else if  strings.HasPrefix(imgUrl, "/") { imgUrl = ctx.BaseUrl +  imgUrl }
        req = NewRequest(imgUrl)

		r, err := ctx.Client.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : request failled  %s\n", os.Args[0], err.Error())
			continue
		}
		defer r.Body.Close()
		id := uuid.New()
		imgPath = path.Join(dir, strings.ReplaceAll(id.String(), "-", ""))

		b, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : error readding file  %s\n", os.Args[0], err.Error())
			continue
		}
		s := http.DetectContentType(b)
		if !allowedMimeType[s] {
			fmt.Fprintf(os.Stderr, "%s : Not allowed mimetype  %s\n", os.Args[0], s)
			continue
		}

		f, err := os.Create(imgPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : can't create image error  %s\n", os.Args[0], err.Error())
			continue
		}
		defer f.Close()

		_, err = f.Write(b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s : can't write to file error  %s\n", os.Args[0], err.Error())
		}
        fmt.Println("file downloaded at ", imgPath)
	}
}
