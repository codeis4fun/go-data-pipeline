package data

import "time"

type Person struct {
	Name        string
	Surname     string
	City        string
	DateOfBirth time.Time
	Height      float64
	Weight      float64
}

type PersonTreated struct {
	completeName string
	country      string
	age          int
	bmi          float64
}

func NewPerson(name, surname, city string, dateOfBirth time.Time, height, weight float64) Person {
	return Person{
		Name:        name,
		Surname:     surname,
		City:        city,
		DateOfBirth: dateOfBirth,
		Height:      height,
		Weight:      weight,
	}
}

func NewPersonTreated(completeName, country string, age int, bmi float64) PersonTreated {
	return PersonTreated{
		completeName: completeName,
		country:      country,
		age:          age,
		bmi:          bmi,
	}
}
