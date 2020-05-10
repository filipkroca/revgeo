// Copyright 2020 Filip Kroča. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package revgeo provides a reverse geocoding capability. Latitude and longitude to ISO country code.
// This package uses GeoJSON polygons.
// DATASET source: https://datahub.io/core/geo-countries#data-cli
package revgeo

import (
	"compress/gzip"
	"github.com/paulmach/orb/geojson"
	"log"
	"io/ioutil"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/planar"
	"github.com/gobuffalo/packr/v2"
)

// dataPath path to compressed geojson
const dataPath = "./data"
const fileName = "countries.geojson.gz"

// Decoder holds gemetry in memory and provides method Geocode()
type Decoder struct {
	fc *geojson.FeatureCollection
}

func (d *Decoder) loadGeometry() {

	// embed assets with packr
	box := packr.New("data", dataPath)
		
	file, err := box.Open(fileName)
	defer file.Close()
	if err != nil {
    log.Panic(err)
	}

	gzipReader, err := gzip.NewReader(file)
  defer gzipReader.Close()
  if err != nil {
    log.Panic(err)
	}

	s, err := ioutil.ReadAll(gzipReader)

	featureCollection, err := geojson.UnmarshalFeatureCollection(s)
	if err != nil {
    log.Panic(err)
	}

	d.fc = featureCollection
	
}

// Geocode gets lat and lng, returns country ISO code
func (d *Decoder) Geocode(lat float64, lng float64) (string, error) {

	point := orb.Point{lat, lng}
	if d.fc == nil {
		// load gemetry from GeoJSON
		d.loadGeometry()
	}
	for _, feature := range d.fc.Features {
			// Try on a MultiPolygon to begin
			multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
			if isMulti {
					if planar.MultiPolygonContains(multiPoly, point) {
							return fmt.Sprintf("%v", feature.Properties["ISO_A3"]), nil
					}
			} else {
					// Fallback to Polygon
					polygon, isPoly := feature.Geometry.(orb.Polygon)
					if isPoly {
							if planar.PolygonContains(polygon, point) {
								return fmt.Sprintf("%v", feature.Properties["ISO_A3"]), nil
							}
					}
			}
	}
	return "", fmt.Errorf("Unable to find country for lat: %v lng: %v", lat, lng)
}
