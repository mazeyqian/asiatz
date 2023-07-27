package asiatz

import "testing"

func TestConvertShanghaiToUTC(t *testing.T) {
	shanghaiTime := "08:00"
	expected := "00:00"
	actual, err := ConvertShanghaiToUTC(shanghaiTime)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if actual != expected {
		t.Errorf("Expected %s but got %s", expected, actual)
	}
}
