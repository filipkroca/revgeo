# Package revgeo provides a reverse geocoding capability. Latitude and longitude to ISO country code

Certain purpose:

- determine in which country is input LatLng coordinate
- very fast processing thru GeoJSON polygons

## func (\*Decoder) Geocode

Geocode gets lat and lng, returns country ISO code

```text
  func (d *Decoder) Geocode(lat float64, lng float64) (string, error)
```

## Example

```go
package main

import ("github.com/filipkroca/revgeo"
"fmt"
"log"
)

func main() {
  decoder := revgeo.Decoder{}

  var lat float64
  var lng float64

  lat = 48.75181328781114
  lng = 16.234285804999985
  country, err := decoder.Geocode(lng, lat)

  if err != nil {
    log.Panicln(err)
  }

  fmt.Println(country)
  // Output:
  // CZE
}
```

## Performance

```text
BenchmarkDecode_Geocode-8   300   5265501 ns/op      49 B/op      2 allocs/op
```
