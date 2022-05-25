package rolling

func Covariance(ro1 *RollingObject, ro2 *RollingObject) (cov float64) {
	ro1Values := ro1.Values()
	ro2Values := ro2.Values()

	if len(ro1Values) != len(ro2Values) {
		panic("Lengths of values are not equal, cannot calculate covariance")
	}
	mean1 := ro1.Calc("avg")
	mean2 := ro2.Calc("avg")

	var total float64
	for i, ro1V := range ro1Values {
		ro2V := ro2Values[i]
		total += ((ro1V - mean1) * (ro2V - mean2))
	}
	return total / (float64(len(ro1Values)) - 1)
}
