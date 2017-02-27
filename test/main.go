package main

import "flag"

var (
	apiurl  string // logfile to monitor
	apikey  string // refresh rate of the log counters
	debug   bool   // number of top sites to display
	refresh int    // time to consider for alarms

)

func init() {
	flag.StringVar(&apiurl, "url", "https://api.bloomsky.com/api/skydata/", "BloomSky API URL")
	flag.StringVar(&apikey, "key", "", "BloomSky API Key as found at https://dashboard.bloomsky.com/user#api")
	flag.IntVar(&refresh, "refresh", 60, "Refresh interval in seconds")
	flag.BoolVar(&debug, "debug", false, "enable debug")
}

func main() {
	flag.Parse()

	connection := bloomsky.NewConfig(apiurl, apikey, nil)

}
