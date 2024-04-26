package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func getPDFPageCount(pdfPath string) (int, error) {
	count, err := api.PageCountFile(pdfPath)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	dirFlag := flag.String("d", ".", "Directory to check PDF files.")
	flag.Parse()

	fmt.Println("[+] pdf page running...")

	files, err := os.ReadDir(*dirFlag)
	if err != nil {
		panic(err)
	}

	var totalCount int
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".pdf" {
			pdfPath := filepath.Join(*dirFlag, file.Name())

			pageCount, err := getPDFPageCount(pdfPath)
			if err != nil {
				log.Println("error getting page count:", err)
				continue
			}

			fmt.Printf("%s: %d pages\n", file.Name(), pageCount)
			totalCount += pageCount
		}
	}
}
