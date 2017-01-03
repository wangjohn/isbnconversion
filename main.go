package isbnconversion

import (
	"bufio"
	"encoding/csv"
	"os"
)

func readcsv() {
	filename := "/Users/wangjohn/Downloads/headlist_new.csv"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(file)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	output := [][]string{}
	for _, record := range records {
		prod, err := ISBN13to10(record[0])
		if err != nil {
			panic(err)
		}
		output = append(output, []string{prod})
	}

	outfile := "/Users/wangjohn/Downloads/chegg_products.csv"
	oFile, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer oFile.Close()

	w := csv.NewWriter(bufio.NewWriter(oFile))
	w.WriteAll(output)
	if err := w.Error(); err != nil {
		panic(err)
	}
}
