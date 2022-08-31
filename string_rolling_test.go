package rolling

import "testing"

var strSamples = [10]string{
	"94", "61", "64", "38", "87", "93", "98", "51", "60", "41",
}

func Test5Window(t *testing.T) {
	ro := NewRollingStringObject(5)
	for _, f := range strSamples {
		ro.Add(f)
	}
	result := ro.Join(",", false)
	expected := "93,98,51,60,41"
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}

func Test5WindowRev(t *testing.T) {
	ro := NewRollingStringObject(5)
	for _, f := range strSamples {
		ro.Add(f)
	}
	result := ro.Join(":", true)
	expected := "41:60:51:98:93"
	if result != expected {
		t.Errorf("e.Value() is %v, wanted %v", result, expected)
	}
}
