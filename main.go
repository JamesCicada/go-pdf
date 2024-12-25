package main

import (
	"bytes"
	"fmt"

	"github.com/dslipak/pdf"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := readPdf("/book.pdf")
	check(err)
	fmt.Print(content)
}

func readPdf(path string) (string, error) {
	f, err := pdf.Open(path)
	// remember close file
	r := f.Trailer().Reader()
	defer r.Close()
	check(err)
	var buf bytes.Buffer
	b, err := f.GetPlainText()
	check(err)
	buf.ReadFrom(b)
	return buf.String(), nil
}
