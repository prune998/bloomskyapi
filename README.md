# BloomSkyAPI
BloomSky API access in GO

## Description
This is a package to consume the BloomSky API.
It is used to retrieve data from your SKY2 and STORM modules.

## Usage
This is a go package, so go get it :

```
go get github.com/prune998/bloomskyapi
```

### Example application
#### bsdumper
This application connect to the API and dump the JSON data

```JSON
{
  "UTC": -5,
  "CityName": "Charny",
  "Storm": {
    "UVIndex":  1
    "RainRate": 0,
    "SustainedWindSpeed": 0,
    "RainDaily": 0,
    "WindDirection": "N",
    "WindGust": 0
  },
  "Searchable": true,
  "DeviceName": "3372",
  "RegisterTime": 1479772486,
  "DST": 0,
  "BoundedPoint": "",
  "LON": -72.23369140625,
  "Point": {},
  "VideoList": [
    "http://s3.amazonaws.com/bskytimelapses/jpxnGERVGERGERVV_-5_2017-02-24.mp4",
    "http://s3.amazonaws.com/bskytimelapses/jpxnGERVGERGERVV_-5_2017-02-25.mp4",
    "http://s3.amazonaws.com/bskytimelapses/jpxnGERVGERGERVV_-5_2017-02-26.mp4",
    "http://s3.amazonaws.com/bskytimelapses/jpxnGERVGERGERVV_-5_2017-02-27.mp4",
    "http://s3.amazonaws.com/bskytimelapses/jpxnGERVGERGERVV_-5_2017-02-28.mp4"
  ],
  "VideoListC": null,
  "DeviceID": "E07EF5C1CB5",
  "NumOfFollowers": 7,
  "LAT": 48.87451172,
  "ALT": 69,
  "Data": {
    "Luminance": 9999,
    "Temperature": 32.990000021457675,
    "ImageURL": "http://s3-us-west-1.amazonaws.com/bskyimgs/jpxnGERVGERGERVVqJ1krp2ql5OpmJw=.jpg",
    "TS": 1488328969,
    "Rain": false,
    "Humidity": 57,
    "Pressure": 29.88,
    "DeviceType": "SKY2",
    "Voltage": 2599,
    "Night": true,
    "UVIndex": 9999,
    "ImageTS": 1488322507
  },
  "FullAddress": "Rue Ernest-Talbot, QC, CA",
  "StreetName": "Rue Ernest-Talbot",
  "PreviewImageList": []
}
```
