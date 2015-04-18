package main

import (
	"github.com/SlyMarbo/rss"
	"time"
)

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

func ParseSite(address string) (map[time.Time][]PostEntry, error) {
	feed, err := rss.Fetch(address)
	if err != nil {
		return nil, err
	}

	entries := map[time.Time][]PostEntry{}
	for _, v := range feed.Items {
		entry := AddPost(v.Title, v.Content, feed.Nickname, v.Link, v.ID, false, v.Date)

		entries[v.Date] = append(entries[v.Date], entry)
	}

	return entries, nil
}
