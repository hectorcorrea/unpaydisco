package main

import (
	"net/http"
)

func about(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	s := NewSession(values, resp, req)
	renderTemplate(s, "ui/html/about.html", settings)
}

func home(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	s := NewSession(values, resp, req)
	renderTemplate(s, "ui/html/index.html", settings)
}

func search(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	search := NewSearch(settings)
	results, err := search.Search(req.URL.Query())

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error during search", err)
	} else {
		renderTemplate(s, "ui/html/results.html", results)
	}
}

func viewOne(values RouteValues, resp http.ResponseWriter, req *http.Request) {
	search := NewSearch(settings)
	record, err := search.Get(values["id"])

	s := NewSession(values, resp, req)
	if err != nil {
		renderError(s, "Error retrieving document from Solr", err)
	} else {
		renderTemplate(s, "ui/html/one.html", record)
	}
}
