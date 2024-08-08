package revgeo

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func ExampleDecoder_geocode() {
	decoder := Decoder{}

	var lat float64
	var lng float64

	lat = 48.75181328781114
	lng = 16.234285804999985
	country, err := decoder.Geocode(lat, lng)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(country)
	// Output:
	// CZE
}
func BenchmarkDecode_Geocode(b *testing.B) {

	decoder := Decoder{}
	// first call invoke GeoJSON load
	decoder.Geocode(rand.Float64()*100, rand.Float64()*100)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decoder.Geocode(rand.Float64()*100, rand.Float64()*100)
	}

}
