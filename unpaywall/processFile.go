package unpaywall

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"solr"
)

// ReadFile reads an unpaywall JSON file and converts it
// to Document objects
func ReadFile(fileName string) ([]Document, error) {
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		return []Document{}, err
	}

	docs := []Document{}
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		var doc Document
		err = json.Unmarshal([]byte(line), &doc)
		if err != nil {
			break
		}
		docs = append(docs, doc)
	}

	if err != io.EOF {
		return docs, err
	}
	return docs, nil
}

func FileToSolr(solrCoreURL string, fileName string) error {
	docs, err := ReadFile(fileName)
	if err != nil {
		return err
	}

	var solrDocs []map[string]interface{}
	for _, doc := range docs {
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

	s := solr.New(solrCoreURL, true)
	err = s.Post(solrDocs)
	return err
}
