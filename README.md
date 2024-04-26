# pdfpage

pdfpage iterates through the PDF files in a directory, prints, and calculates the sum
of all the files page-count.

## Installation

```shell
$ go install github.com/mehdieidi/pdfpage@latest
...
```

## Usage

The path of the root directory is given using the -d flag.
Default is the current directory.
```shell
$ pdfpage -d "./books"
...
```

## Todo

* Recursive directory search.
* Fix pdf library issues.
* Filtering support.

