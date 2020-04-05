package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func baseTemplate() *template.Template {
	// A template with our helper functions
	t := template.New("layout")
	t.Funcs(template.FuncMap{
		"safeHTML": func(s string) template.HTML { return template.HTML(s) },
		"safeURL":  func(u string) template.URL { return template.URL(u) },
	})
	return t
}

func renderNotFound(s Session) {
	log.Printf(fmt.Sprintf("Not found (%s)", s.Req.URL.Path))
	t, err := baseTemplate().ParseFiles("ui/html/layout.html", "ui/html/notfound.html")
	if err != nil {
		log.Printf("Error rendering not found page :(")
		// perhaps render a hard coded string?
	} else {
		s.Resp.WriteHeader(http.StatusNotFound)
		t.Execute(s.Resp, nil)
	}
}

func renderError(s Session, title string, err error) {
	log.Printf("ERROR: %s - %s (%s)", title, err, s.Req.URL.Path)
	t, err := baseTemplate().ParseFiles("ui/html/layout.html", "ui/html/error.html")
	if err != nil {
		log.Printf("Error rendering error page :(")
		// perhaps render a hard coded string?
	} else {
		s.Resp.WriteHeader(http.StatusInternalServerError)
		t.Execute(s.Resp, nil)
	}
}

func loadTemplate(s Session, viewName string) (*template.Template, error) {
	t, err := baseTemplate().ParseFiles("ui/html/layout.html", viewName)
	if err != nil {
		renderError(s, fmt.Sprintf("Loading view %s", viewName), err)
		return nil, err
	}
	log.Printf("Loaded template %s (%s)", viewName, s.Req.URL.Path)
	return t, nil
}

func renderTemplate(s Session, viewName string, viewModel interface{}) {
	t, err := loadTemplate(s, viewName)
	if err != nil {
		log.Printf("Error loading: %s, %s ", viewName, err)
	} else {
		err = t.Execute(s.Resp, viewModel)
		if err != nil {
			log.Printf("Error rendering: %s, %s ", viewName, err)
		}
	}
}
