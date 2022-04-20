package rolling

import "time"

func (ro *RollingTimeObject) WeightValues(date time.Time) {
	weighted := make(valueAndTimes, len(ro.values))
	for _, vad := range ro.values {
		days := date.Sub(vad.date).Hours() / 24
		vad.weighted = vad.value * (1 / (ro.weight * days))
		weighted = append(weighted, vad)
	}
	ro.values = weighted
}

// NewRollingObject - set up a new rolling object with a supplied window with the default settings
func NewRollingTimeWeightedObject(weight float64) *RollingTimeObject {
	return &RollingTimeObject{
		weight:           weight,
		values:           []valueAndTime{},
		ignoreNanValues:  ignoreNanValuesDefault,
		ignoreInfValues:  ignoreInfValuesDefault,
		ignoreZeroValues: ignoreZeroValuesDefault,
	}
}
