package main

import (
	"fmt"
	"github.com/joeriggs/ssa/data"
	"github.com/joeriggs/retirement/social_security"
	"github.com/joeriggs/retirement/wages"
)

func main() {
	fmt.Printf("Beginning the retirement program.\n")

	fmt.Printf("1985 MaxEarnings %d.\n", social_security.MaxEarnings(1985))
	fmt.Printf("1995 MaxEarnings %d.\n", social_security.MaxEarnings(1995))
	fmt.Printf("2005 MaxEarnings %d.\n", social_security.MaxEarnings(2005))
	fmt.Printf("2015 MaxEarnings %d.\n", social_security.MaxEarnings(2015))
	fmt.Printf("2025 MaxEarnings %d.\n", social_security.MaxEarnings(2025))
	fmt.Printf("====================================\n")

	fmt.Printf("1985 AverageWageIndex %f.\n", social_security.AverageWageIndex(1985))
	fmt.Printf("1995 AverageWageIndex %f.\n", social_security.AverageWageIndex(1995))
	fmt.Printf("2005 AverageWageIndex %f.\n", social_security.AverageWageIndex(2005))
	fmt.Printf("2015 AverageWageIndex %f.\n", social_security.AverageWageIndex(2015))
	fmt.Printf("2025 AverageWageIndex %f.\n", social_security.AverageWageIndex(2025))
	fmt.Printf("====================================\n")

	a, b, c, d, e := social_security.BendPoints(1985)
	fmt.Printf("1985 BendPoints %d %d %d %d %d.\n", a, b, c, d, e)
	a, b, c, d, e  = social_security.BendPoints(1995)
	fmt.Printf("1995 BendPoints %d %d %d %d %d.\n", a, b, c, d, e)
	a, b, c, d, e  = social_security.BendPoints(2005)
	fmt.Printf("2005 BendPoints %d %d %d %d %d.\n", a, b, c, d, e)
	a, b, c, d, e  = social_security.BendPoints(2015)
	fmt.Printf("2015 BendPoints %d %d %d %d %d.\n", a, b, c, d, e)
	a, b, c, d, e  = social_security.BendPoints(2025)
	fmt.Printf("2025 BendPoints %d %d %d %d %d.\n", a, b, c, d, e)

	var person1Wages wages.WageList = make(wages.WageList)

	statement := data.New("Your_Social_Security_Statement_Data.xml")

	var numEarningsYears = statement.NumEarningsYears()

	for i := 0; i < numEarningsYears; i++ {
		var year, ficaEarnings, medicareEarnings int
		year, ficaEarnings, medicareEarnings = statement.EarningsYear(i)
		fmt.Printf("          %4v   %9v  %9v\n", year, ficaEarnings, medicareEarnings)

		ficaEarningsFloat := float32(ficaEarnings)
		wages.WageListAdd(year, ficaEarningsFloat, person1Wages)
	}

	social_security.Benefit(1962, person1Wages)
}

