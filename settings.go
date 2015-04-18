package main

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	ini "github.com/vaughan0/go-ini"
	"reflect"
	"strings"
)

type Settings struct {
	SiteName     string    // The ttle of the site
	Description  string    // The site description
	DisplayURL   string    `valid:"url"` // The url to be displayd
	OwnerName    string    // The owner of the site
	OwnerEmail   string    `valid:email` // The owner email
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
		site := ToStr(reflect.ValueOf(value))
		if valid.IsURL(site) {
			settings.DisplayURL = site
		}
	case "owner_name":
		settings.OwnerName = ToStr(reflect.ValueOf(value))
	case "owner_email":
		email := ToStr(reflect.ValueOf(value))
		if valid.IsEmail(email) {
			settings.OwnerEmail = email
		}
	case "items_per_page":
		item := ToStr(reflect.ValueOf(value))
		if valid.IsInt(item) {
			i, _ := valid.ToInt(item)
			settings.ItemsPerPage = int(i)
		} else {
			settings.ItemsPerPage = 25
		}
	case "public_dir":
		path := ToStr(reflect.ValueOf(value))
		var is_path, _ = valid.IsFilePath(path)
		if is_path && ValidDir(path) {
			settings.PublicDir = path
		} else {
			settings.PublicDir = ""
		}
	case "template_dir":
		path := ToStr(reflect.ValueOf(value))
		var is_path, _ = valid.IsFilePath(path)
		if is_path && ValidDir(path) {
			settings.TemplateDir = path
		} else {
			settings.TemplateDir = ""
		}

	case "template_name":
		var templates []string
		arr := len(settings.TemplateName)
		copy(templates, strings.Split("|", ToStr(reflect.ValueOf(value))))
		//copy(settings.TemplateName[:], templates)
		for i, v := range templates {
			if i > arr { // we can have only 5 items (or less)
				break
			}

			settings.TemplateName[i] = v
		}

	}

}

func LoadConf(filename string) (Settings, error) {
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
					if valid.IsURL(value) {
						site.Site = value // override section name that is a site
					} else {
						return Settings{}, errors.New("Invalid site url")
					}
				case "feed":
					if valid.IsRequestURL(value) {
						site.Feed = value
					} else {
						return Settings{}, errors.New("Invalid feed url")
					}
				case "author":
					site.Author = value
				case "rtl":
					content := strings.ToLower(value)
					if content == "true" || content == "false" {
						site.Rtl, _ = valid.ToBoolean(content)
					} else {
						site.Rtl = false
					}
				}

			}
		}

		if site_section {
			settings.Sites = append(settings.Sites, site)
			site_section = false
		}

	}

	return settings, nil
}
