package revgeo

import ("testing"
"math/rand"
"fmt"
"log"
)

func ExampleDecoder_geocode() {
	decoder := Decoder{}
	decoder.loadGeometry()

	var lat float64
	var lng float64

	lat = 48.75181328781114
	lng = 16.234285804999985
	country, err := decoder.geocode(lng, lat)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(country)
	// Output:
	// CZE
}
func BenchmarkDecode_geocode(b *testing.B) {
	
	decoder := Decoder{}
	decoder.loadGeometry()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		decoder.geocode(rand.Float64()*100, rand.Float64()*100)		
	}
	
}