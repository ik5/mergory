package main

/*
import (
	ini "github.com/vaughan0/go-ini"
)
*/

type Settings struct {
	Planet struct {
		SiteName     string    // The ttle of the site
		DisplayURL   string    // The url to be displayd
		OwnerName    string    // The owner of the site
		OwnerEmail   string    // The owner email
		ItemsPerPage int       // The number of items at a page
		PublicDir    string    // The path to generate the content into
		TemplateDir  string    // The path of the templates
		TemplateName [5]string /*
			    The name of the templates to use:
					   0. index
					   1. atom
					   2. rss20
					   3. rss10
					   4. opml
		*/
	}
}

func LoadConf(filename string) {

}
