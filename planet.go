/*
 */
package main

import (
	"github.com/SlyMarbo/rss"
	//"os"
	"time"
)

var sites = []SiteRec{}
var entries = map[time.Time][]PostEntry{}

func ParseSite(address string) error {
	_, err := rss.Fetch("http://example.com/rss")
	if err != nil {
		return err
	}

	//feed.Title

	return nil
}
