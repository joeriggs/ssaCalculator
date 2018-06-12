package main

import (
	"fmt"
	"github.com/joeriggs/ssa/data"
	"github.com/joeriggs/ssa/social_security"
	"github.com/joeriggs/ssa/wages"
)

func main() {
	fmt.Printf("Beginning the retirement program.\n")

	var person1Wages wages.List = make(wages.List)

	statement := data.New("Your_Social_Security_Statement_Data.xml")

	var numEarningsYears = statement.NumEarningsYears()

	for i := 0; i < numEarningsYears; i++ {
		var year, ficaEarnings, medicareEarnings int
		year, ficaEarnings, medicareEarnings = statement.EarningsYear(i)
		fmt.Printf("          %4v   %9v  %9v\n", year, ficaEarnings, medicareEarnings)

		ficaEarningsFloat := float32(ficaEarnings)
		wages.Add(year, ficaEarningsFloat, person1Wages)
	}

	social_security.Benefit(1962, person1Wages)
}

