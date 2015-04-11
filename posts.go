package main

import (
	"github.com/SlyMarbo/rss"
	"time"
)

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

func AddSite(Title, Description, Address, Feed, Author string, Rtl bool) SiteRec {
	return SiteRec{
		Title,
		Description,
		Address,
		Feed,
		Author,
		Rtl,
	}
}

func AddPost(Title, Body, Author, Url, ID string, Rtl bool, Modified time.Time) PostEntry {
	return PostEntry{
		Title,
		Body,
		Author,
		Url,
		ID,
		Rtl,
		Modified,
	}
}

func ParseSite(address string, sites *[]SiteRec) (map[time.Time][]PostEntry, error) {
	feed, err := rss.Fetch(address)
	if err != nil {
		return nil, err
	}

	site := AddSite(feed.Title, feed.Description, feed.Link, feed.UpdateURL, feed.Nickname, false)

	*sites = append(*sites, site)
	entries := map[time.Time][]PostEntry{}
	for _, v := range feed.Items {
		entry := AddPost(v.Title, v.Content, feed.Nickname, v.Link, v.ID, false, v.Date)

		entries[v.Date] = append(entries[v.Date], entry)
	}

	return entries, nil
}
