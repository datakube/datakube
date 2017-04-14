package main

type Hamster struct {
	Database DatabaseConfiguration
}

func NewHamster(configuration DatabaseConfiguration) *Hamster {
	hamster := new(Hamster)
	hamster.Database = configuration
	return hamster
}

func (*Hamster) run() {





}