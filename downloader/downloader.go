package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/iamajraj/go-concurrent-file-downloader/utils"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func DownloadFile(urlStr, outpurDir string, progress *mpb.Progress) error {
	parsedUrl, err := url.Parse(urlStr)

	if err != nil {
		return fmt.Errorf("error parsing the url %s: %v", urlStr, err)
	}

	filename := path.Base(parsedUrl.Path)
	if filename == "" || filename == "/" {
		filename = "default_download"
	}

	outputPath := path.Join(outpurDir, filename)

	resp, err := http.Get(urlStr)
	if err != nil {
		return fmt.Errorf("error downloading %s: %v", urlStr, err)
	}
	defer resp.Body.Close()

	contentLength := resp.ContentLength
	if contentLength <= 0 {
		contentLength = 1
	}

	bar := progress.AddBar(
		contentLength,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("%s: ", utils.Truncate(filename, 30)), decor.WC{W: 32, C: decor.DSyncWidth}),
			decor.Percentage(decor.WC{W: 5}),
		),
		mpb.AppendDecorators(decor.CountersKiloByte("% .2f / % .2f", decor.WC{W: 15})),
		mpb.BarWidth(50),
	)

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file for %s: %v", urlStr, err)
	}
	defer outFile.Close()

	reader := bar.ProxyReader(resp.Body)
	defer reader.Close()

	_, err = io.Copy(outFile, reader)

	if err != nil {
		return fmt.Errorf("error saving file for %s: %v", urlStr, err)
	}
	return nil
}
