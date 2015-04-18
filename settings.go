package main

import (
	ini "github.com/vaughan0/go-ini"
	"reflect"
	"strings"
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

func set_settings(key string, value interface{}, settings *Settings) {
	switch key {
	case "site_name":
		settings.SiteName = ToStr(reflect.ValueOf(value))
	case "description":
		settings.Description = ToStr(reflect.ValueOf(value))
	case "display_url":
		settings.DisplayURL = ToStr(reflect.ValueOf(value))
	case "owner_name":
		settings.OwnerName = ToStr(reflect.ValueOf(value))
	case "owner_email":
		settings.OwnerEmail = ToStr(reflect.ValueOf(value))
	case "items_per_page":
		settings.ItemsPerPage = ToInt(reflect.ValueOf(value))
	case "public_dir":
		settings.PublicDir = ToStr(reflect.ValueOf(value))
	case "template_dir":
		settings.TemplateDir = ToStr(reflect.ValueOf(value))
	case "template_name":
		copy(settings.TemplateName[:], strings.Split("|", ToStr(reflect.ValueOf(value))))
	}

}

func (Settings) LoadConf(filename string) (Settings, error) {
	file, err := ini.LoadFile(filename)
	if err != nil {
		return Settings{}, err
	}

	var settings Settings
	var site_section bool

	for name, _ := range file {
		var site SiteRec
		site_section = false
		for key, value := range file[name] {
			switch name {
			case "settings":
				set_settings(key, value, &settings)

			case "default":
				set_settings(key, value, &settings)

			default:
				site_section = true
				site.Site = name
				switch key {
				case "title":
					site.Title = value
				case "description":
					site.Description = value
				case "site":
					site.Site = value // override section name that is a site
				case "feed":
					site.Feed = value
				case "author":
					site.Author = value
				case "Rtl":
					site.Rtl = value == "true"
				}

			}
		}

		if site_section {
			settings.Sites = append(settings.Sites, site)
			site_section = false
		}

	}

	return Settings{}, nil
}
