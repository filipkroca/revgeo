# Package revgeo provides a reverse geocoding capability. Latitude and longitude to ISO country code

Certain purpose:

- determine in which country is input LatLng coordinate
- very fast processing thru GeoJSON polygons

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
