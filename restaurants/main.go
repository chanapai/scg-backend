package main

import (
	"context"
	"log"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

func main() {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyDXDkKDyBH0OygF_DSN3foQgIKKt7VD8v0"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &maps.NearbySearchRequest{
		// Bang sue latitude and longitude
		Location: &maps.LatLng{Lat: 13.803071, Lng: 100.5369955},
		// 10000 = 10km
		Radius:   10000,
	}

	resp, err := c.NearbySearch(context.Background(), r)
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	pretty.Println(resp)
}
