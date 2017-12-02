package percent

// Percent calculate what is [percent]% of [number]
// For Example 25% of 200 is 50
// It returns result as float64
func Percent(pcent int, all int) float64{
  percent := ((float64(all) * float64(pcent)) / float64(100))
  return percent
}

// PercentOf calculate [number1] is what percent of [number2]
// For example 300 is 12.5% of 2400
// It returns result as float64
func PercentOf(current int, all int) float64 {
  percent := ( float64(current) * float64(100) ) / float64(all)
  return percent
}

// Change calculate what is the percentage increase/decrease from [number1] to [number2]
// For example 60 is 200% increase from 20
// It returns result as float64
func Change(before int, after int) float64{
  diff := float64(after) - float64(before)
  realDiff := diff / float64(before)
  percentDiff := 100 * realDiff

  return percentDiff
}
