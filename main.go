package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func countPages(pdfPath string) (int, error) {
	count, err := pdfcpu.PageCountFile(pdfPath)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func main() {
	dir := flag.String("d", ".", "Directory to check PDF files.")
	flag.Parse()

	fmt.Println("[+] pdfpage running...")

	files, err := os.ReadDir(*dir)
	if err != nil {
		panic(err)
	}

	var totalPageCount int
	var failures []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if filepath.Ext(file.Name()) == ".pdf" {
			pdfPath := filepath.Join(*dir, file.Name())

			pageCount, err := countPages(pdfPath)
			if err != nil {
				fmt.Printf("[x] error getting page count of file %s: %v\n", file.Name(), err)
				failures = append(failures, file.Name())
				continue
			}

			fmt.Printf("%s: %d pages\n", file.Name(), pageCount)
			totalPageCount += pageCount
		}
	}

	fmt.Println("---------------------------------------------")

	fmt.Println("[+] List of failed PDFs:")
	for _, failure := range failures {
		fmt.Printf("- %s\n", failure)
	}

	fmt.Println("[+] Total page count:", totalPageCount)
}
