package rolling

import "time"

func (ro *RollingTimeObject) WeightValues(date time.Time) {
	ro.weighted = []float64{}
	for _, vad := range ro.values {
		days := date.Sub(vad.date).Hours() / 24
		if days == 0 {
			days = 1
		}
		ro.weighted = append(ro.weighted, vad.value*(1/(ro.weight*days)))
	}
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
