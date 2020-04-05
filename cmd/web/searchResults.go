package main

import (
	"fmt"
	"strings"

	"github.com/hectorcorrea/solr"
	// "solr"
)

type SearchItem struct {
	Doi           string
	DoiURL        string
	Title         string
	Year          int
	JournalName   string
	PublisherName string
	Genre         string
	OaURL         string
	Authors       []string
	SolrDoc       solr.Document
	HasAuthors    bool
	AuthorsString string
	UnpayURL      string
}

type SearchResults struct {
	Q           string
	Documents   []solr.Document
	Facets      solr.Facets
	NumFound    int
	Start       int
	Rows        int
	First       int
	Last        int
	BaseUrl     string
	Url         string
	UrlNoQ      string
	NextPageUrl string
	PrevPageUrl string
	Response    solr.SearchResponse
	Items       []SearchItem
}

func NewSearchResults(resp solr.SearchResponse, baseUrl string) SearchResults {
	results := SearchResults{
		NumFound:    resp.NumFound,
		Facets:      resp.Facets,
		Start:       resp.Start,
		Rows:        resp.Rows,
		BaseUrl:     baseUrl,
		Url:         baseUrl + "?" + resp.Url,
		PrevPageUrl: baseUrl + "?" + resp.PrevPageUrl,
		NextPageUrl: baseUrl + "?" + resp.NextPageUrl,
		Response:    resp,
		Documents:   resp.Documents,
		Items:       solrDocumentsToSearchItems(resp.Documents),
	}

	if results.NumFound > 0 {
		results.First = results.Start + 1
		results.Last = results.First + results.Rows
		if results.Last > results.NumFound {
			results.Last = results.NumFound
		}
	}

	// TODO: Update the solr module to URL encode the values used in the facets
	// otherwise the links are broken when they have some special characters
	// (e.g. college & libraries)
	results.Facets.SetAddRemoveUrls(results.Url)

	if resp.Q != "*" {
		results.Q = resp.Q
		results.UrlNoQ = baseUrl + "?" + resp.UrlNoQ
	}

	return results
}

func solrDocToSearchItem(doc solr.Document) SearchItem {
	item := SearchItem{Doi: doc.Data["id"].(string)}
	item.UnpayURL = fmt.Sprintf("https://api.unpaywall.org/v2/%s?email=unpaydisco", item.Doi)

	if doc.Data["title_txt_en"] != nil {
		item.Title = doc.Data["title_txt_en"].(string)
		if doc.IsHighlighted("title_txt_en") {
			item.Title = doc.HighlightFor("title_txt_en")
		}
	}

	if doc.Data["doi_url_s"] != nil {
		item.DoiURL = doc.Data["doi_url_s"].(string)
	}

	if doc.Data["journal_s"] != nil {
		item.JournalName = doc.Data["journal_s"].(string)
	}

	if doc.Data["oa_url_s"] != nil {
		item.OaURL = doc.Data["oa_url_s"].(string)
	}

	if doc.Data["year_i"] != nil {
		item.Year = int(doc.Data["year_i"].(float64))
	}

	if doc.Data["authors_ss"] != nil {
		for _, author := range doc.Data["authors_ss"].([]interface{}) {
			item.Authors = append(item.Authors, author.(string))
		}
		item.HasAuthors = len(item.Authors) > 0
		item.AuthorsString = strings.Join(item.Authors, ", ")
	}

	if doc.Data["genre_s"] != nil {
		item.Genre = doc.Data["genre_s"].(string)
	}

	if doc.Data["publisher_s"] != nil {
		item.PublisherName = doc.Data["publisher_s"].(string)
	}
	return item
}

func solrDocumentsToSearchItems(docs []solr.Document) []SearchItem {
	items := []SearchItem{}
	for _, doc := range docs {
		item := solrDocToSearchItem(doc)
		items = append(items, item)
	}
	return items
}
