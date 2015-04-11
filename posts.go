package posts

import (
	"time"
)

type SiteRec struct {
	Title       string // Blog title
	Description string // Site Description
	Site        string // Main site address e.g. http://blog.example.com/
	Feed        string // The feed address e.g. http://blog.example.com/computers/feed/
	Author      string // The name of the author for the site
	Rtl         string // The main direction of the blog (Left to Right or Right to Left)
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

func (SiteRec) SetSite(Title, Description, Address, Feed, Author, Rtl) SiteRec {
	return SiteRec{
		Title,
		Description,
		Address,
		Feed,
		Author,
		Rtl,
	}
}

func (PostEntry) AddPost(Title, Body, Author, Url, ID, Rtl, Modified) PostEntry {
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

	site := SiteRec.SetSite(
		feed.Title, feed.Description, feed.Link, feed.UpdateURL, feed.Author, false,
	)

	*sites = append(*sites, site)
	entries := map[time.Time][]PostEntry{}
	for _, v := range feed.Item {
		entry := AddPost(
			v.Title,
			v.Content,
			v.Nickname,
			v.Link,
			v.ID,
			false,
			v.Date,
		)

		entries[v.Date] = append(entries[v.date], entry)
	}

	return nil
}
