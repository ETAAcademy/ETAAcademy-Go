package main

import "fmt"

// A report structure composed of structures for temperature and location.
type report struct {
	sol        int
	temperature temperature
	location   location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{
		sol: 15,
		temperature: t,
		location: bradbury,
	}

	fmt.Printf("%+v\n", report) 
	fmt.Printf("a balmy %vº C\n", report.temperature.high)
}

// 3.2 Forwarding
// The temperature type and average method can be used independently of the weather report.
// func (t temperature) average() celsius {
// 	return (t.high + t.low) / 2
// }

// Write a method that forwards to the real implementation instead.
// func (r report) average() celsius {
// 	return r.temperature.average()
// }

// func main() {
// 	bradbury := location{-4.5895, 137.4417}
// 	t := temperature{high: -1.0, low: -78.0}

// 	fmt.Printf("average %vº C\n", t.average())

// 	report := report{
// 		sol: 15, 
// 		temperature: t,
// 		location: bradbury,
// 	}

// 	fmt.Printf("average %vº C\n", report.temperature.average())
// 	fmt.Printf("average %vº C\n", report.average())

// 	fmt.Printf("%+v\n", report) 
// 	fmt.Printf("a balmy %vº C\n", report.temperature.high)
// }

// 3.3 Embedding
// type report struct {
// 	sol int
// 	temperature 
// 	location
// }

// func main() {
// 	bradbury := location{-4.5895, 137.4417}
// 	t := temperature{high: -1.0, low: -78.0}

// 	fmt.Printf("average %vº C\n", t.average())

// 	report := report{
// 		sol: 15, 
// 		temperature: t,
// 		location: bradbury,
// 	}
	   
// 	fmt.Printf("average %vº C\n", report.average())

// 	fmt.Printf("%vº C\n", report.high)
// 	report.high = 32
// 	fmt.Printf("%vº C\n", report.temperature.high)
// }

// 3.4 Name collisions

// func (s sol) days(s2 sol) int {
// 	days := int(s2 - s)
// 	if days < 0 {
// 		days = -days
// 	}
// 	return days
// }

// func (l location) days(l2 location) int {
// 	// To-do: complicated distance calculation 
// 	return 5
// }

// func (r report) days(s2 sol) int {
// 	return r.sol.days(s2)
// }

// func main() {
// 	report := report{sol: 15}
 
// 	fmt.Println(report.sol.days(1446))
// 	fmt.Println(report.days(1446))
// }
