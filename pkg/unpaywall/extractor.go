package unpaywall

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Extractor is used to extract the open access documents from a Unpaywall
// data snapshot.
type Extractor struct {
	fileName  string
	batchSize int
	file      *os.File
	docsRead  int
	docsOA    int
}

// NewExtractor creates a new extractor
func NewExtractor(fileName string) Extractor {
	extractor := Extractor{
		fileName: fileName,
	}
	return extractor
}

// Extract is the core method that extracts the open access documents
// from the data snapshot file. Data is outputted to StdOut.
func (ext *Extractor) Extract() error {
	var file *os.File
	var err error

	if ext.fileName == "-" {
		file = os.Stdin
	} else {
		file, err = os.Open(ext.fileName)
		defer file.Close()
		if err != nil {
			return err
		}
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		var doc Document
		err = json.Unmarshal([]byte(line), &doc)
		if err != nil {
			return err
		}

		ext.docsRead++
		if doc.BestOaLocation.URL != "" {
			ext.docsOA++
			fmt.Printf("%s", line)
		}
	}
}
