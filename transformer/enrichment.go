package transformer

import (
	"math"
	"time"
)

type EnrichCountryByCity struct {
	City    string
	Country string
}

type EnrichAgeByDateOfBirth struct {
	DateOfBirth time.Time
	Age         int
}

type EnrichBMIByHeightAndWeight struct {
	Height float64
	Weight float64
	BMI    float64
}

func NewEnrichCountryByCity(city string) *EnrichCountryByCity {
	return &EnrichCountryByCity{City: city}
}

func NewEnrichAgeByDateOfBirth(dateOfBirth time.Time) *EnrichAgeByDateOfBirth {
	return &EnrichAgeByDateOfBirth{DateOfBirth: dateOfBirth}
}

func NewEnrichBMIByHeightAndWeight(height, weight float64) *EnrichBMIByHeightAndWeight {
	return &EnrichBMIByHeightAndWeight{Height: height, Weight: weight}
}

func (t *EnrichCountryByCity) Transform() {
	var countries map[string]string
	countries = make(map[string]string)

	countries["LONDON"] = "UK"

	t.Country = countries[t.City]
}

func (t *EnrichAgeByDateOfBirth) Transform() {
	currentDate := time.Now()
	birthDate := t.DateOfBirth

	// Calculate age based on the difference in years
	age := currentDate.Year() - birthDate.Year()

	// Check if the birthdate for this year has occurred or not
	if birthDate.YearDay() > currentDate.YearDay() {
		age--
	}

	t.Age = age
}

func (t *EnrichBMIByHeightAndWeight) Transform() {
	t.BMI = math.Round(t.Weight / (t.Height * t.Height))
}
