package asiatz

import "testing"

type testConversion struct {
	time     string
	expected string
}

var tests = map[string][]testConversion{
	"Shanghai": {
		{"08:00", "00:00"},
		{"12:00", "04:00"},
		{"23:59", "15:59"},
	},
	"Tokyo": {
		{"09:00", "00:00"},
		{"12:00", "03:00"},
		{"23:59", "14:59"},
	},
	"Hong Kong": {
		{"08:00", "00:00"},
		{"12:00", "04:00"},
		{"23:59", "15:59"},
	},
	"Singapore": {
		{"08:00", "00:00"},
		{"12:00", "04:00"},
		{"23:59", "15:59"},
	},
	"Taipei": {
		{"08:00", "00:00"},
		{"12:00", "04:00"},
		{"23:59", "15:59"},
	},
}

func runConversionTests(t *testing.T, tests []testConversion, conversionFunc func(string) (string, error)) {
	for _, test := range tests {
		actual, err := conversionFunc(test.time)
		if err != nil {
			t.Errorf("Unexpected error for %s: %v", test.time, err)
			continue
		}
		if actual != test.expected {
			t.Errorf("Expected %s for %s but got %s", test.expected, test.time, actual)
		}
	}
}

func TestAllConversions(t *testing.T) {
	for timezone, tests := range tests {
		t.Run(timezone, func(t *testing.T) {
			switch timezone {
			case "Shanghai":
				runConversionTests(t, tests, ShanghaiToUTC)
			case "Tokyo":
				runConversionTests(t, tests, TokyoToUTC)
			case "Hong Kong":
				runConversionTests(t, tests, HongKongToUTC)
			case "Singapore":
				runConversionTests(t, tests, SingaporeToUTC)
			case "Taipei":
				runConversionTests(t, tests, TaipeiToUTC)
			default:
				t.Errorf("Unexpected timezone %s", timezone)
			}
		})
	}
}
