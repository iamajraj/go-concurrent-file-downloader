package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/iamajraj/go-concurrent-file-downloader/downloader"
	"github.com/iamajraj/go-concurrent-file-downloader/utils"
	"github.com/vbauerster/mpb/v8"
)

func main() {
	outputDir := flag.String("output", ".", "Directory to save downloaded files")
	flag.Parse()

	urls := flag.Args()
	if len(urls) == 0 {
		fmt.Println("Please provide a list of URLs to download")
		return
	}

	if err := os.MkdirAll(*outputDir, os.ModePerm); err != nil {
		fmt.Printf("Failed to create output directory %s: %v\n", *outputDir, err)
		return
	}

	progress := mpb.New(mpb.WithWaitGroup(&sync.WaitGroup{}))
	workerPool := utils.NewWorkerPool(3)

	for _, url := range urls {
		workerPool.AddTask(func() {
			err := downloader.DownloadFile(url, *outputDir, progress)
			if err != nil {
				fmt.Println("Error:", err)
			}
		})
	}

	workerPool.Wait()
	progress.Wait()
	fmt.Println("All downloads completed!")
}
