package location

import "math"

const (
	r = 6371.0 // km
	p = math.Pi / 180
)

// GetDistance returns Kilometers
func GetDistance(p1, p2 Location) float64 {
	lat1 := *p1.Latitude
	lon1 := *p1.Longitude
	lat2 := *p2.Latitude
	lon2 := *p2.Longitude

	a := 0.5 - math.Cos(
		(lat2-lat1)*p,
	)/2 + math.Cos(
		lat1*p,
	)*math.Cos(
		lat2*p,
	)*(1-math.Cos((lon2-lon1)*p))/2

	return 2 * r * math.Asin(math.Sqrt(a))
}
