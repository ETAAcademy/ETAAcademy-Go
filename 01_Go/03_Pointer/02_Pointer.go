package main

import (
	"fmt"
	"strings"
)

type person struct {
	name        string
	superpower  string
	age         int
}

// Function to increase the age of a person (pass by value).
// func birthday(p person) {
// 	p.age++
// }

// Function to increase the age of a person (pass by reference).
func birthday(p *person) {
	p.age++
}

// Method to increase the age of the person.
func (p *person) birthday() {
	p.age++
}

type stats struct {
	level     int
	endurance int
	health    int
}

// Function to level up the character's stats.
func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

type talker interface {
	talk() string
}

// Function to shout the talker's message in uppercase.
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martain struct{}

// Method for martain to talk.
func (m martain) talk() string {
	return "neck neck"
}

type laser int

// Method for laser to talk (pointer receiver).
func (l *laser) talk() string {
	return "pew pew"
}

func main() {

	// venus := person{
	// 	name:       "Venus",
	// 	superpower: "imagination",
	// 	age:       16,
	// }
	// Functions are passed by value.
	// birthday(venus)
	// fmt.Println("%+v\n", venus) // Output: 16

	venus := person{
		name:       "Venus",
		superpower: "imagination",
		age:       16,
	}
	
	birthday(&venus) // Call birthday with a pointer to venus
	fmt.Println("%+v\n", venus) // Output: 17

	athena := &person{
		name:       "Athena",
		superpower: "mind",
		age:       18,
	}

	athena.birthday() // Call birthday method on athena
	fmt.Println("%+v\n", athena) // Output: 19

	boss := character{name: "Mary"}

	levelUp(&boss.stats) // Level up the stats of boss
	fmt.Printf("%+v\n", boss.stats) // Output: {name:Boss stats:{level:1 endurance:56 health:280}}

	shout(martain{})      // Output: NECK NECK
	shout(&martain{})     // Output: NECK NECK

	pew := laser(2)
 
	// shout(pew) // Error: pew does not implement talker interface
	shout(&pew) // Output: PEW PEW
}
