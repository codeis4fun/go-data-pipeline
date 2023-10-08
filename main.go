package main

import (
	"fmt"
	"time"

	"github.com/codeis4fun/go-data-pipeline/data"
	"github.com/codeis4fun/go-data-pipeline/pipeline"
	"github.com/codeis4fun/go-data-pipeline/transformer"
)

func main() {
	// data
	person := data.NewPerson("John", "Doe", "London", time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC), 1.8, 80)

	// transformations
	nameToUppercase := transformer.NewToUpper(person.Name)
	surnameToUppercase := transformer.NewToUpper(person.Surname)
	cityToUppercase := transformer.NewToUpper(person.City)

	// Create a pipeline
	p1 := pipeline.NewPipeline("Transform name, surname and city")
	// Add stages to the pipeline
	p1.
		AddStage("transform name to uppercase", nameToUppercase).
		AddStage("transform surname to uppercase", surnameToUppercase).
		AddStage("transform city to uppercase", cityToUppercase)
	// Run the pipeline
	p1.Run()

	// transformations
	joinNameAndSurname := transformer.NewJoinStrings(nameToUppercase.Output, surnameToUppercase.Output)
	enrichCountryByCity := transformer.NewEnrichCountryByCity(cityToUppercase.Output)
	enrichAgeByDateOfBirth := transformer.NewEnrichAgeByDateOfBirth(person.DateOfBirth)
	enrichBMIByWeightAndHeight := transformer.NewEnrichBMIByHeightAndWeight(person.Height, person.Weight)

	// Create a pipeline
	p2 := pipeline.NewPipeline("Transform to complete name, enrich country, enrich age and enrich BMI")
	// Add stages to the pipeline
	p2.
		AddStage("transform to complete name", joinNameAndSurname).
		AddStage("enrich country by city", enrichCountryByCity).
		AddStage("enrich age by date of birth", enrichAgeByDateOfBirth).
		AddStage("enrich BMI by weight and height", enrichBMIByWeightAndHeight)
	// Run the pipeline
	p2.Run()

	// Assign the output of the pipeline to the data
	personTreated := data.NewPersonTreated(joinNameAndSurname.Output, enrichCountryByCity.Country, enrichAgeByDateOfBirth.Age, enrichBMIByWeightAndHeight.BMI)

	// Print the result
	fmt.Printf("Person treated: %+v\n", personTreated)
}
