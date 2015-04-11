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
	Site     SiteRec   // The "pointer" to the SiteRec
	Title    string    // The post title
	Body     string    // The post body
	Author   string    // The post Author
	Url      string    // The direct post Url
	ID       string    // The actual post id url
	Rtl      bool      // The direction of the post (Left to Right or Right to Left)
	Modified time.Time // The date and time of the Post
}

func ParseSite(address string, sites *[]SiteRec, posts *[]PostEntry) error {
	feed, err := rss.Fetch(address)
	if err != nil {
		return err
	}

	site := SiteRec{
		feed.Title,
		feed.Description,
		feed.Link,
		feed.UpdateURL,
		feed.Author,
		false,
	}

	return nil
}
