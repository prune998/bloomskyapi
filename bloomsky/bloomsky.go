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
	StreetName string `json:"street_name"`
	UtcOffset  int    `json:"utc_offset"`
	Outdoor    struct {
		UvIndex        string  `json:"uv_index"`
		Luminance      int     `json:"luminance"`
		Temperature    float64 `json:"temperature"`
		IsRaining      bool    `json:"is_raining"`
		Humidity       int     `json:"humidity"`
		Pressure       float64 `json:"pressure"`
		IsNight        bool    `json:"is_night"`
		ImageURL       string  `json:"image_url"`
		ImageTimestamp int     `json:"image_timestamp"`
		DeviceType     string  `json:"device_type"`
		DataTimestamp  int     `json:"data_timestamp"`
		Voltage        int     `json:"voltage"`
	} `json:"outdoor"`
	FavoritesCount interface{} `json:"favorites_count"`
	Altitude       float64     `json:"altitude"`
	Indoor         struct {
		Temperature interface{} `json:"temperature"`
		Humidity    interface{} `json:"humidity"`
	} `json:"indoor"`
	FollowersCount      int      `json:"followers_count"`
	FullAddress         string   `json:"full_address"`
	DeviceName          string   `json:"device_name"`
	Latitude            float64  `json:"latitude"`
	IsDst               int      `json:"is_dst"`
	VideoUrls           []string `json:"video_urls"`
	Longitude           float64  `json:"longitude"`
	IsSearchable        bool     `json:"is_searchable"`
	CityName            string   `json:"city_name"`
	RegisteredTimestamp int      `json:"registered_timestamp"`
	DeviceID            string   `json:"device_id"`
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
func (d *Data) Get(c *connection) error {
	req, err := http.NewRequest("GET", c.APIURL, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", c.APIKey)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//var record *Data
	if err := json.NewDecoder(resp.Body).Decode(d); err != nil {
		return err
	}
	//d = record
	return nil

}

// PrintJsonData
func (d *Data) PrintJsonData() error {
	fmt.Printf("%s", d)
	return nil
}

func (d *Data) PrintTemp() error {
	fmt.Printf("%s", d)
	return nil
}

func (d *Data) PrintLoc() error {
	fmt.Printf("%s", d.StreetName)
	return nil
}

// GetJpeg

// BuildMovie
