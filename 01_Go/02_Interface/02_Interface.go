package main

import (
	"fmt"
	"strings"
	"time"
)

var t interface {
	talk() string
}

// Type martian satisfies the interface t (1).
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

// Type laser satisfies the interface t (2).
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// Interfaces are often declared as types and end with "-er".
type talker interface {
	talk() string
}

// Function accepts any value that satisfies the talker interface.
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

// Converts "Earth time" to "eta star time" (bad version).
// func stardate(t time.Time) float64 {
// 	doy := float64(t.YearDay())
// 	h := float64(t.Hour()) / 24.0
// 	return 1000 + doy + h
// }

type etastadater interface {
	YearDay() int
	Hour() int
}

// Converts "Earth time" to "eta star time" (good version).
func etastardate(t etastadater) float64 {
	day := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + day + h
}

type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}

func main() {
	t = martian{} // Since martian implements the interface t, it can be assigned this way.
	fmt.Println(t.talk()) // Output: nack nack

	t = laser(3)
	fmt.Println(t.talk()) // Output: pew pew pew

	shout(martian{}) // Output: NACK NACK
	shout(laser(2))  // Output: PEW PEW

	s := starship{laser(2)}
	fmt.Println(s.talk()) // Output: pew pew
	shout(s)              // Output: PEW PEW

	// Convert "Earth time".
	day := time.Date(2046, 8, 9, 6, 6, 6, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", etastardate(day)) 

	// Convert star time from s.
	s2 := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", etastardate(s2)) 
}
