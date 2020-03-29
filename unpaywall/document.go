package unpaywall

import "fmt"

// OaLocation represents the URL of the Open Access Location
type OaLocation struct {
	URL string `json:"url"`
}

// Author represents the author information
type Author struct {
	Given  string `json:"given"`
	Family string `json:"family"`
}

// Document represents an Unpaywall document
type Document struct {
	Doi            string     `json:"doi"`
	DoiURL         string     `json:"doi_url"`
	Year           int        `json:"year"`
	Title          string     `json:"title"`
	JournalName    string     `json:"journal_name"`
	BestOaLocation OaLocation `json:"best_oa_location"`
	Authors        []Author   `json:"z_authors"`
	Genre          string     `json:"genre"`
	Publisher      string     `json:"publisher"`
}

func (d Document) String() string {
	return fmt.Sprintf("%s - %s", d.Doi, d.Title)
}
