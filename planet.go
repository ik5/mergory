/*
 */
package main

import (
	"github.com/SlyMarbo/rss"
	//ini "github.com/vaughan0/go-ini"
	//"os"
	"time"
)

type SiteRec struct {
	Title  string // Blog title
	Site   string // Main site address e.g. http://blog.example.com/
	Feed   string // The feed address e.g. http://blog.example.com/computers/feed/
	Author string // The name of the author for the site
	Rtl    string // The main direction of the blog (Left to Right or Right to Left)
}

type PostEntry struct {
	Site     SiteRec   // The "pointer" to the SiteRec
	Title    string    // The post title
	Body     string    // The post body
	Author   string    // The post Author
	Url      string    // The direct post Url
	Rtl      bool      // The direction of the post (Left to Right or Right to Left)
	Modified time.Time // The date and time of the Post
}

var sites = []SiteRec{}
var entries = []map[time.Time][]PostEntry{}

func LoadConf(filename string) {

}

func ParseSite(address string) error {
	_, err := rss.Fetch("http://example.com/rss")
	if err != nil {
		return err
	}

	//feed.Title

	return nil
}
