package social_security

/* Average Wage Index (a.k.a. AWI) Table, taken from:
 *   https://www.ssa.gov/oact/COLA/AWI.html
 *
 * Dates after 2014 can be estimated using data from this page, but we aren't
 * doing that estimation.  It looks like the ssa.gov website doesn't, so we will
 * do the same as them.
 *   https://www.ssa.gov/oact/TR/TRassum.html
 *
 * A description of how to use this table can be found at:
 *   https://www.ssa.gov/oact/ProgData/retirebenefit1.html
 */
var averageWageIndexes = map[int]float32{
  1951:   2799.16,
  1952:   2973.32,
  1953:   3139.44,
  1954:   3155.64,
  1955:   3301.44,
  1956:   3532.36,
  1957:   3641.72,
  1958:   3673.80,
  1959:   3855.80,
  1960:   4007.12,
  1961:   4086.76,
  1962:   4291.40,
  1963:   4396.64,
  1964:   4576.32,
  1965:   4658.72,
  1966:   4938.36,
  1967:   5213.44,
  1968:   5571.76,
  1969:   5893.76,
  1970:   6186.24,
  1971:   6497.08,
  1972:   7133.80,
  1973:   7580.16,
  1974:   8030.76,
  1975:   8630.92,
  1976:   9226.48,
  1977:   9779.44,
  1978:  10556.03,
  1979:  11479.46,
  1980:  12513.46,
  1981:  13773.10,
  1982:  14531.34,
  1983:  15239.24,
  1984:  16135.07,
  1985:  16822.51,
  1986:  17321.82,
  1987:  18426.51,
  1988:  19334.04,
  1989:  20099.55,
  1990:  21027.98,
  1991:  21811.60,
  1992:  22935.42,
  1993:  23132.67,
  1994:  23753.53,
  1995:  24705.66,
  1996:  25913.90,
  1997:  27426.00,
  1998:  28861.44,
  1999:  30469.84,
  2000:  32154.82,
  2001:  32921.92,
  2002:  33252.09,
  2003:  34064.95,
  2004:  35648.55,
  2005:  36952.94,
  2006:  38651.41,
  2007:  40405.48,
  2008:  41334.97,
  2009:  40711.61,
  2010:  41673.83,
  2011:  42979.61,
  2012:  44321.67,
  2013:  44888.16,
  2014:  46481.52,
}

var awiMostRecentYear int = 0
func mostRecentAverageWageIndexYear() int {
  if awiMostRecentYear == 0 {
    for key, _ := range averageWageIndexes {
      if key > awiMostRecentYear {
        awiMostRecentYear = key
      }
    }
  }
  return awiMostRecentYear
}

func AverageWageIndex(year int) float32 {
  if year > mostRecentAverageWageIndexYear() {
    return averageWageIndexes[mostRecentAverageWageIndexYear()]
  } else {
    return averageWageIndexes[year]
  }
}

