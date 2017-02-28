package bloomsky

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIVersion - API version, to support API changes one day
const APIVersion = "1.3"

// DefaultAPIURL Default API URL
const DefaultAPIURL = "https://api.bloomsky.com/api/skydata/"

// Data - data from the BLoomSky API
type Data struct {
	UTC      int
	CityName string
	Storm    struct {
		RainRate           float64
		SustainedWindSpeed float64
		RainDaily          float64
		WindDirection      string
		WindGust           float64
	}
	Searchable   bool
	DeviceName   string
	RegisterTime int
	DST          int
	BoundedPoint string
	LON          float64
	Point        struct {
	}
	VideoList      []string
	VideoListC     []interface{}
	DeviceID       string
	NumOfFollowers int
	LAT            float64
	ALT            float64
	Data           struct {
		Luminance   int
		Temperature float64
		ImageURL    string
		TS          int
		Rain        bool
		Humidity    int
		Pressure    float64
		DeviceType  string
		Voltage     int
		Night       bool
		UVIndex     string
		ImageTS     int
	}
	FullAddress      string
	StreetName       string
	PreviewImageList []interface{}
}

// config - config to reach the API
type connection struct {
	APIVersion string
	APIURL     string
	APIKey     string
	Client     *http.Client
}

// NewConfig - create the object to connect to the API
func NewConfig(apiurl, apikey, apiversion string) (*connection, error) {
	if apiversion == "" {
		apiversion = APIVersion
	}
	if apiurl == "" {
		apiurl = DefaultAPIURL
	}
	if apikey == "" {
		return nil, fmt.Errorf("API key can't be empty")
	}

	return &connection{
			APIVersion: apiversion,
			APIURL:     apiurl,
			APIKey:     apikey,
			Client:     &http.Client{},
		},
		nil
}

// Get - get data from the API
func Get(c *connection) (*Data, error) {
	req, err := http.NewRequest("GET", c.APIURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.APIKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	r := []Data{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	//d = &r[0]
	//	fmt.Printf("StreetName : %v\n", r[0].StreetName)
	//d = record
	// fmt.Printf("\n%s\n", r)
	return &r[0], nil

}

// PrintJsonData
func (d *Data) PrintJsonData() error {
	fmt.Printf("%s\n", d)
	return nil
}

func (d *Data) PrintTemp() error {
	fmt.Printf("%s\n", d)
	return nil
}

func (d *Data) PrintLoc() error {
	fmt.Printf("%s\n", d.StreetName)
	return nil
}

// GetJpeg

// BuildMovie
