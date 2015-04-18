package main

import "time"

type SiteRec struct {
	Title       string // Blog title
	Description string // Site Description
	Site        string // Main site address e.g. http://blog.example.com/
	Feed        string // The feed address e.g. http://blog.example.com/computers/feed/
	Author      string // The name of the author for the site
	Rtl         bool   // The main direction of the blog (Left to Right or Right to Left)
}

type PostEntry struct {
	//Site     SiteRec   // The "pointer" to the SiteRec
	Title    string    // The post title
	Body     string    // The post body
	Author   string    // The post Author
	Url      string    // The direct post Url
	ID       string    // The actual post id url
	Rtl      bool      // The direction of the post (Left to Right or Right to Left)
	Modified time.Time // The date and time of the Post
}

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
