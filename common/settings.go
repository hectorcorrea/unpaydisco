package common

import (
	"encoding/json"
	"io/ioutil"
)

type Import struct {
	BatchSize int `json:"batchSize"`
}
type Settings struct {
	ServerAddress string            `json:"serverAddress"`
	SolrCoreUrl   string            `json:"solrCoreUrl"`
	SolrOptions   map[string]string `json:"solrOptions"`
	SolrFacets    map[string]string `json:"solrFacets"`
	SearchFl      []string          `json:"searchFl"`
	ViewOneFl     []string          `json:"viewOneFl"`
	Import        Import            `json:"import"`
}

func LoadSettings(filename string) (Settings, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Settings{}, err
	}

	var settings Settings
	err = json.Unmarshal(bytes, &settings)
	return settings, err
}
