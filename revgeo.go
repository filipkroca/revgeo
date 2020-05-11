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
	"bytes"
)

// dataPath path to compressed geojson
const dataPath = "./data"
const fileName = "countries.geojson.gz"

// Decoder holds gemetry in memory and provides method Geocode()
type Decoder struct {
	polygons []polygonWithIso
	multiPolygons []multiPolygonWithIso
}

type polygonWithIso struct {
	polygon orb.Polygon
	iso string
}

type multiPolygonWithIso struct {
	multiPolygon orb.MultiPolygon
	iso string
}

func (d *Decoder) loadGeometry() {

	var polygons []polygonWithIso
	var multiPolygons []multiPolygonWithIso

	data, err := Asset("data/countries.geojson.gz")
	if err != nil {
		log.Panicln(err)
	}
	

	gzipReader, err := gzip.NewReader(bytes.NewReader(data))
  defer gzipReader.Close()
  if err != nil {
    log.Panic(err)
	}

	s, err := ioutil.ReadAll(gzipReader)

	featureCollection, err := geojson.UnmarshalFeatureCollection(s)
	if err != nil {
    log.Panic(err)
	}

	for _, feature := range featureCollection.Features {
		// Try on a MultiPolygon to begin
		multiPoly, isMulti := feature.Geometry.(orb.MultiPolygon)
		if isMulti {
			multiPolygons = append(multiPolygons, multiPolygonWithIso{multiPoly, feature.Properties["ISO_A3"].(string)})
		} else {
				polygon, isPoly := feature.Geometry.(orb.Polygon)
				if isPoly {
					polygons = append(polygons, polygonWithIso{polygon, feature.Properties["ISO_A3"].(string)})
				}
		}
	}

	d.polygons = polygons
	d.multiPolygons = multiPolygons
	
}

// Geocode gets lat and lng, returns country ISO code
func (d *Decoder) Geocode(lat float64, lng float64) (string, error) {

	point := orb.Point{lat, lng}
	if d.polygons == nil {
		// load gemetry from GeoJSON
		d.loadGeometry()
	}

	for _, polygon := range d.polygons {
		if planar.PolygonContains(polygon.polygon, point) {
			return fmt.Sprintf("%v", polygon.iso), nil
		}
	}

	for _, multiPolygon := range d.multiPolygons {
		if planar.MultiPolygonContains(multiPolygon.multiPolygon, point) {
			return fmt.Sprintf("%v", multiPolygon.iso), nil
		}
	}
	
	return "", fmt.Errorf("Unable to find country for lat: %v lng: %v", lat, lng)
}