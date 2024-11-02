package main

import (
	"fmt"
)

// Coordinate represented in degrees, minutes, seconds for a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h rune
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	
	fmt.Println(lat.decimal(), long.decimal())
	
	curiosity := location{lat.decimal(), long.decimal()}
	
	fmt.Println(curiosity)
	
}

// 2.2 Constructor
// type location struct {
// 	lat, long float64
// }

// curiosity := location{lat.decimal(), long.decimal()}

// // newLocation creates a location from latitude and longitude d/m/s coordinates.
// func newLocation(lat, long coordinate) location {
//  return location{lat.decimal(), long.decimal()}
// }

// func main() {
// 	// Bradbury Landing: 4º35'22.2" S, 137º26'30.1" E
// 	lat := coordinate{4, 35, 22.2, 'S'}
// 	long := coordinate{137, 26, 30.12, 'E'}
	
// 	fmt.Println(lat.decimal(), long.decimal())
	
// 	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
// 	fmt.Println(curiosity)
// }

// 2.3 Class alternative
// type world struct {
// 	radius float64
// }
   
// // Calculates distance using the Spherical Law of Cosines.
// func (w world) distance(p1, p2 location) float64 {
// 	s1, c1 := math.Sincos(rad(p1.lat))
// 	s2, c2 := math.Sincos(rad(p2.lat))
// 	clong := math.Cos(rad(p1.long - p2.long))
// 	return w.radius * math.Acos(s1*s2+c1*c2*clong)
// }
   
// // rad converts degrees to radians.
// func rad(deg float64) float64 {
// 	return deg * math.Pi / 180
// }

// func main() {
// 	var mars = world{radius: 3389.5}
// 	spirit := location{-14.5684, 175.472636}
// 	opportunity := location{-1.9462, 354.4734}
	
// 	dist := mars.distance(spirit, opportunity)
// 	fmt.Printf("%.2f km\n", dist)
	
// }
