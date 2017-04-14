package main

import (
	"fmt"
)

type Hamster struct {
	Database DatabaseConfiguration
}

func NewHamster(configuration DatabaseConfiguration) *Hamster {
	hamster := new(Hamster)
	hamster.Database = configuration
	return hamster
}

func (*Hamster) run() {
	fmt.Printf("I'm running my lord");

}