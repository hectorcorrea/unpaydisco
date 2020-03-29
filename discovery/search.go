package discovery

import (
	"net/url"
	"unpaydisco/common"

	"github.com/hectorcorrea/solr"
	// "solr"
)

type Search struct {
	settings common.Settings
}

func NewSearch(settings common.Settings) Search {
	return Search{settings: settings}
}

func (search Search) Get(id string) (solr.Document, error) {
	params := solr.NewGetParams("id:"+id, search.settings.ViewOneFl, search.settings.SolrOptions)
	s := solr.New(search.settings.SolrCoreUrl, true)
	doc, err := s.Get(params)
	if err != nil {
		return solr.Document{}, err
	}
	return doc, nil
}

func (search Search) Search(qs url.Values) (SearchResults, error) {
	params := solr.NewSearchParamsFromQs(qs, search.settings.SolrOptions, search.settings.SolrFacets)
	params.Fl = search.settings.SearchFl
	solr := solr.New(search.settings.SolrCoreUrl, true)
	resp, err := solr.Search(params)
	if err != nil {
		return SearchResults{}, err
	}
	rootURL := search.settings.Discovery.RootURL
	searchURL := rootURL + "/search"
	results := NewSearchResults(resp, searchURL)
	return results, nil
}
