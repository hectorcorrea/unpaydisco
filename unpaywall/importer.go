package unpaywall

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"solr"
)

type Importer struct {
	fileName  string
	batchSize int
	file      *os.File
	s         solr.Solr
}

func NewImporter(solrCoreURL string, fileName string, batchSize int) Importer {
	importer := Importer{
		fileName:  fileName,
		batchSize: batchSize,
		s:         solr.New(solrCoreURL, true),
	}
	return importer
}

func (imp Importer) Import() error {
	file, err := os.Open(imp.fileName)
	defer file.Close()
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for {
		var docs []Document
		docs, err = imp.readBatch(reader)
		if err != nil {
			break
		}

		if len(docs) == 0 {
			break
		}

		err = imp.batchToSolr(docs)
		if err != nil {
			break
		}
	}
	return err
}

func (imp Importer) readBatch(reader *bufio.Reader) ([]Document, error) {
	docs := []Document{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return docs, nil
		}
		if err != nil {
			return docs, err
		}

		var doc Document
		err = json.Unmarshal([]byte(line), &doc)
		if err != nil {
			return []Document{}, err
		}
		docs = append(docs, doc)
		if len(docs) == imp.batchSize {
			return docs, nil
		}
	}
}

func (imp Importer) batchToSolr(batch []Document) error {
	var solrDocs []map[string]interface{}
	for _, doc := range batch {
		solrDoc := map[string]interface{}{
			"id":           doc.Doi,
			"doi_url_s":    doc.DoiURL,
			"year_i":       doc.Year,
			"title_txt_en": doc.Title,
			"journal_s":    doc.JournalName,
			"oa_url_s":     doc.BestOaLocation.URL,
		}
		solrDocs = append(solrDocs, solrDoc)
	}

	return imp.s.Post(solrDocs)
}
