package main

import (
	"flag"
	"fmt"
	"unpaydisco/common"
	"unpaydisco/discovery"
	"unpaydisco/unpaywall"
)

func main() {
	var settingsFile = flag.String("settings", "", "Name of settings file to use")
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
		err := unpaywall.FileToSolr(settings.SolrCoreUrl, *importFile)
		if err != nil {
			fmt.Printf("Error importing JSON file %s to Solr: %s\r\n", *importFile, settings.SolrCoreUrl)
			fmt.Printf("%s\r\n", err)
		} else {
			fmt.Printf("Imported JSON file %s to Solr: %s\r\n", *importFile, settings.SolrCoreUrl)
			fmt.Printf("OK\r\n")
		}
		return
	}

	discovery.StartWebServer(*settingsFile)
}

func displayHelp() {
	text := `
Must indicate a settings.json file with an structure like this:

	{
	  "serverAddress": "localhost:9001",
	  "solrCoreUrl": "http://localhost:8983/solr/bibdata",
	  "solrOptions" : {
	    "defType": "edismax",
	    "qf": "authorsAll title^100"
	  },
	  "solrFacets": {
	    "subjects_str": "Subjects",
	    "publisher_str": "Publisher"
	  },
		"searchFl": ["id", "title", "subjects", "author"],
	  "viewOneFl": ["id", "title", "authorsAll", "_version_"]
	}`
	fmt.Printf("%s\r\n\r\n", text)
}
