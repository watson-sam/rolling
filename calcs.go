package rolling

import "math"

// Sum - find the sum of all values
func Sum(values []float64) (result float64) {
	for _, v := range values {
		result += v
	}
	return result
}

// Avg - find the mean of all values
func Avg(values []float64) (result float64) {
	return Sum(values) / float64(len(values))
}

// Count - find the number of values
func Count(values []float64) (result float64) {
	return float64(len(values))
}

// Min - find the minimum value
func Min(values []float64) (min float64) {
	min = math.Inf(1)
	for _, vi := range values {
		if min > vi {
			min = vi
		}
	}
	return min
}

// Max - find the maximum value
func Max(values []float64) (max float64) {
	max = math.Inf(-1)
	for _, vi := range values {
		if max < vi {
			max = vi
		}
	}
	return max
}

// NUnique - find the number of distinct values
func NUnique(values []float64) float64 {
	dist := make(map[float64]struct{})
	for _, vi := range values {
		dist[vi] = struct{}{}
	}
	return float64(len(dist))
}

// Var - find the variance of the values
func Var(values []float64) (result float64) {
	mean := Avg(values)
	for _, v := range values {
		result += math.Pow(v-mean, 2)
	}
	return result / float64(len(values))
}

// Std - find the standard deviation of the values
func Std(values []float64) (result float64) {
	return math.Sqrt(Var(values))
}

// Prod - find the product of all values
func Prod(values []float64) (result float64) {
	result = 1
	for _, v := range values {
		result *= v
	}
	return result
}

func runCalc(calc string, values []float64) float64 {
	if calc == "sum" {
		return Sum(values)
	} else if calc == "avg" {
		return Avg(values)
	} else if calc == "count" {
		return Count(values)
	} else if calc == "max" {
		return Max(values)
	} else if calc == "min" {
		return Min(values)
	} else if calc == "nunique" {
		return NUnique(values)
	} else if calc == "std" {
		return Std(values)
	} else if calc == "var" {
		return Var(values)
	} else if calc == "prod" {
		return Prod(values)
	}
	panic(
		"Supplied `calc` argument is not valid - must be one of: 'sum', 'avg', 'min', 'max', 'count', 'nunique', 'std', 'var' or 'prod', received value: " + calc,
	)
}
