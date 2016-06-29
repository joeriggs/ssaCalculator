package main

import (
	"fmt"
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

	var joe_wages wages.WageList = make(wages.WageList)
	wages.WageListAdd(1978,    191, joe_wages)
	
	social_security.Benefit(joe_wages)
}

