package main

import (
	ini "github.com/vaughan0/go-ini"
)

type Settings struct {
	SiteName     string    // The ttle of the site
	Description  string    // The site description
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

	Sites []SiteRec
}

func (Settings) LoadConf(filename string) (Settings, error) {
	file, err := ini.LoadFile(filename)
	if err != nil {
		return Settings{}, err
	}

	var settings Settings

	for name, _ := range file {
		for key, value := range file[name] {
			switch name {
			case "settings":
				switch key {
				case "site_name":
					settings.SiteName = value
				case "description":
					settings.Description = value
				case "display_url":
					settings.DisplayURL = value
				case "owner_name":
					settings.OwnerName = value
				case "owner_email":
					settings.OwnerEmail = value
				}
			//case "default":

			default:
			}
		}
	}

	//site := AddSite(feed.Title, feed.Description, feed.Link, feed.UpdateURL, feed.Nickname, false)
	return Settings{}, nil
}
