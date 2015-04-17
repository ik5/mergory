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

func set_sites(key, value string, settings *Settings) {

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
				set_settings(key, value, &settings)

			case "default":
				set_settings(key, value, &settings)

			}
		}
	}

	//site := AddSite(feed.Title, feed.Description, feed.Link, feed.UpdateURL, feed.Nickname, false)
	return Settings{}, nil
}
