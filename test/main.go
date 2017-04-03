package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/fogleman/gg"

	"../bloomsky"
)

// "github.com/prune998/bloomskyapi/bloomsky"
// for testing :
// curl -ks "https://api.bloomsky.com/api/skydata/?unit=us" -H "Authorization: XXX"| jq '.'
var (
	apiurl  string // logfile to monitor
	apikey  string // refresh rate of the log counters
	jpgpath string // path to a folder to store the jpegs
	debug   bool   // number of top sites to display
	refresh int    // time to consider for alarms
	font    string // font name to write on the images

)

func init() {
	flag.StringVar(&apiurl, "url", "https://api.bloomsky.com/api/skydata/", "BloomSky API URL")
	flag.StringVar(&apikey, "key", "", "BloomSky API Key as found at https://dashboard.bloomsky.com/user#api")
	flag.StringVar(&jpgpath, "path", "./images", "path to a folder to store the jpg pictures")
	flag.IntVar(&refresh, "refresh", 60, "Refresh interval in seconds")
	flag.BoolVar(&debug, "debug", false, "enable debug")
	flag.StringVar(&font, "font", "/Library/Fonts/Arial Black.ttf", "path to a TrueType font file")
}

func main() {
	flag.Parse()

	c, err := bloomsky.NewClient(apiurl, apikey, "")
	if err != nil {
		log.Fatal("Configuration issue")
	}

	data, err := c.Fetch()
	if err != nil {
		log.Fatal("Error processing data, ", err)
	}

	fmt.Printf("StreetName : %s\n", data.StreetName)
	fmt.Printf("Temp : %f\n\n", data.Data.Temperature)
	// fmt.Printf("%s\n", data)
	data.PrintLoc()
	data.PrintJsonData()

	// process the Image
	im, err := gg.LoadImage("jpxnrKmnqZKlq6enqJ1kr5apmpepnJg=.jpg")
	if err != nil {
		log.Fatal(err)
	}
	dc := gg.NewContext(640, 640)
	dc.DrawImage(im, 0, 0)
	//dc.SetRGB(1, 1, 1)
	//dc.Clear()

	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(font, 24); err != nil {
		panic(err)
	}
	windData := fmt.Sprintf("wind %2fkph %s", data.Storm.SustainedWindSpeed, data.Storm.WindDirection)
	dc.DrawStringAnchored(windData, 640/2, 30, 0.5, 0.5)
	//dc := gg.NewContext(640, 640)
	//dc.DrawRoundedRectangle(0, 0, 640, 640, 64)
	//dc.Clip()
	//dc.DrawImage(im, 0, 0)
	dc.SavePNG("out.png")
}
