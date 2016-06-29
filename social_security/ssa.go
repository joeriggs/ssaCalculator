package social_security

/*******************************************************************************
 * The ssa package an implementation of the algorithms that are used to
 * calculate or estimate the Social Security benefits for a retiree.
 *
 * Here's a simple description of how it all works:
 * - The Average Wage Index (AWI) Table contains the AWI for each year, going
 *   back to 1951.
 * - The AWI is used to calculate the "indexing factor" for each year that a
 *   person worked and earned wages.
 * - The "indexing factor" for year xxxx is calculated by dividing the AWI for
 *   the year the worker turned 60 by the AWI for year xxxx.
 * - There is a maximum "Nominal Earnings" each year.  You can see the maximum
 *   values in the maxWages table (max_earnings.go).  If the person earned more
 *   than the max in a given year, then the max is used.
 * - The "indexing factor" is multiplied by their nominal earnings for year xxxx
 *   to calculate their "indexed earnings" for that year.
 * - The highest 35 "indexed earnings" are added together to produce the
 *   "highest-35 total".
 * - The "highest-35 total" is divided by 420 to produce their "Average
 *   Indexed Monthly Earnings" (AIME).  FYI, there are 420 months in 35 years.
 *
 * Here is a simple example:
 * - The worker was born in 1954.  We need to calculate the "indexing factor"
 *   for this person for the year 1976.
 * - They were born in 1954, which means they turn 60 in 2014.
 * - AWI(2014) = 46,481.52
 * - AWI(1976) =  9,226.48
 * - Indexing Factor = 46,481.52 / 9,226.48 = 5.0378.
 * - Indexed Earnings for 1976 for this worker = Nominal_Earnings * 5.0378
 *
 ******************************************************************************/

import ( "github.com/joeriggs/retirement/wages" )

/* This function calculates the "indexing factor" for the specified year. */
func indexingFactor(dob, year int) float32 {
	var indexing_factor float32 = 0.0

	/* If the person is age 60+, then their indexing factor is always 1. */
	var age_60 int = dob + 60
	if year >= age_60 {
		indexing_factor = 1.0
	} else {
		/* If the person is less than age 60, calculate their indexing
		 * factor. */
		var awi_60 float32 = AverageWageIndex(age_60)
		var awi    float32 = AverageWageIndex(year)
		indexing_factor = awi_60 / awi
	}

	return indexing_factor
}

func Benefit(career_earnings wages.WageList) {
	var highest_earnings WageList = WageListCreate()

	WageListHighestIndexedEarnings(career_earnings, highest_earnings)
}

