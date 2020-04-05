package main

import (
	"hectorcorrea.com/unpaydisco/pkg/unpaywall"
	"flag"
	"fmt"

	"hectorcorrea.com/unpaydisco/pkg/common"
)

func main() {
	var settingsFile = flag.String("settings", "", "Name of settings file to use (required)")
	var importFile = flag.String("import", "", "File name to import")
	flag.Parse()

	if *settingsFile == "" {
		displayHelp()
		return
	}

	settings, err := common.LoadSettings(*settingsFile)
	if err != nil {
		fmt.Printf("Error loading settings file: %s\r\n", *settingsFile)
		fmt.Printf("%s\r\n", err)
		displayHelp()
		return
	}

	if *importFile != "" {
		doImport(settings.SolrCoreUrl, *importFile, settings.Import.BatchSize)
		return
	}

	StartWebServer(*settingsFile)
}

func doImport(solrCoreURL string, fileName string, batchSize int) {
	indexer := unpaywall.NewIndexer(solrCoreURL, fileName, batchSize)
	err := indexer.Import()
	if err != nil {
		fmt.Printf("Error importing JSON file %s to Solr: %s\r\n", fileName, solrCoreURL)
		fmt.Printf("%s\r\n", err)
	} else {
		fmt.Printf("Imported JSON file %s to Solr: %s\r\n", fileName, solrCoreURL)
		fmt.Printf("OK\r\n")
	}
	return
}

func displayHelp() {
	help := `
A sample discovery layer on top of the data from Unpaywall.org

unpaydisco -settings settings.json [-import unpaydata.json]

	settings.json is a file with the settings to run the web server and connect to Solr

	unpaydata.json is a file with data from Unpaywall to import into Solr`

	fmt.Printf("%s\r\n\r\n", help)

	sample := `
The format of the settings.json is as follows:

	{
		"serverAddress": "localhost:9001",
		"solrCoreUrl": "http://localhost:8983/solr/bibdata",
		"solrOptions" : {
			"defType": "edismax",
			"qf": "authorsAll title^100",
			"wt": "json",
			"facet.limit": "20",
			"facet.mincount": "1",
			"hl": "on"
		},
		"solrFacets": {
			"journal_s": "Journal",
			"year_i": "Year"
		},
		"searchFl": ["id", "title_txt_en", "oa_url_s"],
		"viewOneFl": []
	}
`
	fmt.Printf("%s\r\n\r\n", sample)
}
