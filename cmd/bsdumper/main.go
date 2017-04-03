package main

import (
	"flag"
	"log"

	"github.com/prune998/bloomskyapi"
)

var (
	apiurl  string // logfile to monitor
	apikey  string // refresh rate of the log counters
	jpgpath string // path to a folder to store the jpegs
	debug   bool   // number of top sites to display
	refresh int    // time to consider for alarms

)

func init() {
	flag.StringVar(&apiurl, "url", "https://api.bloomsky.com/api/skydata/", "BloomSky API URL")
	flag.StringVar(&apikey, "key", "", "BloomSky API Key as found at https://dashboard.bloomsky.com/user#api")
	flag.StringVar(&jpgpath, "path", "./images", "path to a folder to store the jpg pictures")
	flag.IntVar(&refresh, "refresh", 60, "Refresh interval in seconds")
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func main() {
	flag.Parse()

	c, err := bloomskyapi.NewClient(apiurl, apikey, "")
	if err != nil {
		log.Fatal("Configuration issue")
	}

	data, err := c.Fetch()
	if err != nil {
		log.Fatal("Error processing data, ", err)
	}

	data.PrintJsonData()
}
