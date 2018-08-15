package main

import (
	"fmt"
	"time"
	"github.com/joeriggs/ssa/data"
	"github.com/joeriggs/ssa/social_security"
	"github.com/joeriggs/ssa/wages"
)

func main() {
	fmt.Printf("Beginning the retirement program.\n")

	var person1Wages wages.List = make(wages.List)

	statement := data.New("Your_Social_Security_Statement_Data.xml")

	var numEarningsYears = statement.NumEarningsYears()

	// Process the user's statement.
	for i := 0; i < numEarningsYears; i++ {
		var year, ficaEarnings, medicareEarnings int
		year, ficaEarnings, medicareEarnings = statement.EarningsYear(i)
		fmt.Printf("          %4v   %9v  %9v\n", year, ficaEarnings, medicareEarnings)

		ficaEarningsFloat := float32(ficaEarnings)
		wages.Add(year, ficaEarningsFloat, person1Wages)
	}

	/* Get the last income year and wage from the user's history.  We will
	 * use that information to build a "future" earnings picture. */
	_, mostRecentWage := wages.MostRecentYear(person1Wages)

	/* Get their birth year.  Used to calculate future ages. */
	var yearOfBirth = statement.DateOfBirthYear()

	/* Get the current date into t. */
	var currentDate = time.Now()
	var t, err = time.Parse("2006-01-02", currentDate.Format("2006-01-02"))
	if err != nil {
		fmt.Println(err)
	} else {
		var currentYear = t.Year()
		var currentAge = currentYear - yearOfBirth

		/* Get the user's early, full, and delayed retirement ages and
		 * benefit amounts. */
		var earlyAgeYears,                  earlyBenefit   = statement.EarlyRetirement()
		var fullAgeYears,    fullAgeMonths, fullBenefit    = statement.FullRetirement()
		var delayedAgeYears,                delayedBenefit = statement.DelayedRetirement()

		fmt.Printf("\n")
		fmt.Printf("Early Benefit Age %d: Early Benefit Amount $%d\n",
		           earlyAgeYears, earlyBenefit)
		fmt.Printf("Full Benefit Age %d Years %d Months: Full Benefit Amount $%d\n",
		           fullAgeYears, fullAgeMonths, fullBenefit)
		fmt.Printf("Delayed Benefit Age %d: Delayed Benefit Amount $%d\n",
		           delayedAgeYears, delayedBenefit)

		/* If you file early, or if you delay your filing, it will
		 * affect the size of your benefit.  These numbers represent the
		 * penalty or bonus for filing early or delaying.  They're based
		 * on a table found at https://www.ssa.gov/oact/ProgData/ar_drc.html
		 * My array only covers people born on or after 1960 at this time.
		 * We can improve it later.
		 */
		var multipliers = []float32 { .7000, .7500, .8000, .8667, .9333, 1.00, 1.08, 1.16, 1.24 }

		/* We'll calculate these in the loop below.  They need
		 * to be done at the appropriate time. */
		var calculatedEarlyBenefit, calculatedFullBenefit, calculatedDelayedBenefit int

		fmt.Printf("\n")
		fmt.Printf("Retirement    -------- Monthly Benefit If Filing At Age: ---------\n")
		fmt.Printf("Year & Age     62    63    64    65    66    67    68    69    70 \n")
		fmt.Printf("----   ---    ----  ----  ----  ----  ----  ----  ----  ----  ----\n")

		var year = currentYear
		var age = currentAge
		for ; age <= 70; {
			/* Calculate the full benefit for the year.  All early
			 * and delayed benefits are calculated from this value. */
			var fullBenefit = social_security.Benefit(yearOfBirth, person1Wages)
			var benefits [9]string

			var i int
			for i = 0; i < len(benefits); i++ {
				if (age - 62) <= i {
					var tmp float32 = float32(fullBenefit) * multipliers[i]
					benefits[i] = fmt.Sprintf("%4d", int(tmp))
				} else {
					benefits[i] = "----"
				}
			}

			fmt.Printf("%4d   %3d    %4s  %4s  %4s  %4s  %4s  %4s  %4s  %4s  %4s\n",
			           year, age, benefits[0], benefits[1], benefits[2], benefits[3],
			           benefits[4], benefits[5], benefits[6], benefits[7], benefits[8])

			/* Do a few extra calculations along the way, just to
			 * see how closely our calculations match the user's
			 * statement. */
			if age == 62 {
				var tmpFullBenefit = social_security.Benefit(yearOfBirth, person1Wages)
				var earlyBenefit = float32(tmpFullBenefit) * multipliers[0]
				calculatedEarlyBenefit = int(earlyBenefit)
			} else if age == 67 {
				calculatedFullBenefit = social_security.Benefit(yearOfBirth, person1Wages)
			} else if age == 70 {
				var tmpFullBenefit = social_security.Benefit(yearOfBirth, person1Wages)
				var delayedBenefit = float32(tmpFullBenefit) * multipliers[8]
				calculatedDelayedBenefit = int(delayedBenefit)
			}

			/* Add another simulated year to their future earnings. */
			if age < 70 {
				wages.Add(year, mostRecentWage, person1Wages)
			}

			year++
			age++
		}
		fmt.Printf("\n")

		/* Print the early, full, and delayed benefit values that we
		 * calculated in the loop above.  Just for the sake of testing
		 * our algorithms. */
		fmt.Printf("Calculated   early benefit: $%d\n", calculatedEarlyBenefit)
		fmt.Printf("Calculated    full benefit: $%d\n", calculatedFullBenefit)
		fmt.Printf("Calculated delayed benefit: $%d\n", calculatedDelayedBenefit)
	}
}

