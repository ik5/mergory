/*
 */
package main

import (
	//"os"
	//"time"
	"fmt"
)

func main() {
	//var sites = []SiteRec{}
	//var entries = map[time.Time][]PostEntry{}
	settings, err := LoadConf("config/test.ini")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", settings)

}
