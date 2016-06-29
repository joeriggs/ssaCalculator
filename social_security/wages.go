package social_security

import ( "fmt"
         "github.com/joeriggs/retirement/wages"
)

type WageListEntry struct {
	year int
	wage float32
}
type WageList [] WageListEntry

func WageListCreate() WageList {
	var list = make(WageList, 35)
	return list
}

func WageListHighestIndexedEarnings(earnings_list wages.WageList, highest_earnings WageList) {
	for wageYear, wageWage := range earnings_list {
		fmt.Printf("Processing %d (%f).\n", wageYear, wageWage)
		var y int
		for i := 0; i < len(highest_earnings); i++ {
			if highest_earnings[i].wage == 0.0 {
				y = i
				break
			} else if highest_earnings[i].wage < highest_earnings[y].wage {
				y = i
			}
		}

		if highest_earnings[y].year != 0 {
			fmt.Printf("Lowest index is %d (%d %f).\n", y, highest_earnings[y].year, highest_earnings[y].wage)
		}
		if highest_earnings[y].wage > wageWage {
			fmt.Printf("Throwing away %d : %f.\n", wageYear, wageWage)
			continue
		}

		fmt.Printf("Replacing %d (%f) with %d (%f).\n", highest_earnings[y].year, highest_earnings[y].wage, wageYear, wageWage)
		highest_earnings[y].year = wageYear
		highest_earnings[y].wage = wageWage
	}

	/* Take a look at the result. */
	for i := 0; i < len(highest_earnings); i++ {
		if highest_earnings[i].year > 0 {
			fmt.Printf("%2d: %d: %11.2f\n", i, highest_earnings[i].year, highest_earnings[i].wage)
		}
	}
}

