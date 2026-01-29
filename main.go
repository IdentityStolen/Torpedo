package main

import (
	"fmt"
	"math/rand"
	"time"
)

// struct Point resembles ship location
type Point struct {
	X float32
	Y float32
}

// get locations of enemy's warships
func getFoeShips() map[Point]struct{} {
	return map[Point]struct{}{
		{X: 1, Y: 8}:  {},
		{X: 3, Y: 5}:  {},
		{X: 10, Y: 1}: {},
	}
}

// check if ship is hit
func FireTorpedo(point Point, targets map[Point]struct{}, failures chan<- Point,
	success chan<- Point) {

	delay := time.Duration(rand.Intn(20)) * time.Second
	time.Sleep(delay)

	if _, exists := targets[point]; exists {
		success <- point
	} else {
		failures <- point
	}
}

func main() {
	points := []Point{
		{X: 2, Y: 8},
		{X: 3, Y: 8},
		{X: 4, Y: 8},
		{X: 5, Y: 8},
		{X: 6, Y: 8},
		{X: 7, Y: 8},
		{X: 8, Y: 8},
		{X: 9, Y: 8},
		{X: 10, Y: 8},
		// comment this line to see failure hitting target ships
		{X: 1, Y: 8},
	}

	failures := make(chan Point, 6)
	success := make(chan Point, 1)

	targets := getFoeShips()

	for _, point := range points {
		go FireTorpedo(point, targets, failures, success)
	}

	for range len(points) {
		select {
		case f := <-failures:
			fmt.Printf("Miss at %v\n", f)
		case s := <-success:
			fmt.Printf("Hit at %v\n", s)
			return
		}
	}

	fmt.Println("No enemy ship destroyed, pray for your life now!")
}
