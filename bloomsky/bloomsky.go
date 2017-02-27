package bloomsky

import "fmt"

// BsAPIVersion - API version, to support API changes one day
const BsAPIVersion = "1.3"

// DefaultBsAPIURL Default API URL
const DefaultBsAPIURL = "https://api.bloomsky.com/api/skydata/"

// BloomSkyData - data from the BLoomSky API
type BloomSkyData []struct {
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

// BsConfig - config to reach the API
type BsConfig struct {
	APIVersion string
	APIURL     string
	APIKey     string
}

// New - create the object to connect to the API
func New(apiurl, apikey string) (*BsConfig, error) {
	if apiurl == "" {
		apiurl = DefaultBsAPIURL
	}
	if apikey == "" {
		return nil, fmt.Errorf("API key can't be empty")
	}
	return &BsConfig{
		APIVersion: BsAPIVersion,
		APIURL:     apiurl,
		APIKey:     apikey,
	}
}

// Get - get data from the API
func (*BloomSkyData) Get() {

}

// GetData -

// PrintJsonData

// GetJpeg

// BuildMovie
