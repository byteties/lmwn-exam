package apps

import (
	"fmt"
	"testing"
)

func TestGetCountAges(t *testing.T) {
	var cases = []struct {
		ageGroup string
		input int
		want int
	}{
			{"0-30", 15, 1},
			{"N/A", -1, 1},
			{"31-60", 45, 1},
			{"61+", 99, 1},
			{"61+", 15, 0},
	}

	for _, tt := range cases {
		testname := fmt.Sprintf("%s,input %d, expect %d", tt.ageGroup, tt.input,tt.want)
		t.Run(testname, func(t *testing.T) {
				ages := Ages{tt.ageGroup:0}
				GetCountAges(ages,tt.input)
				expected := tt.want

				if ages[tt.ageGroup] != expected {
					t.Errorf("expected '%d' but got '%d'", expected, ages[tt.ageGroup])
				}
		})
	}
}

func TestGetCountProvince(t *testing.T) {
	var cases = []struct {
		provinceGroup string
		key string
		input int
		want int
	}{
			{"Bangkok","Bangkok", 0, 1},
			{"Buriram","Bangkok", 0, 0},
			{"N/A","", 0, 1},
			{"N/A","N/A", 0, 1},
	}

	for _, tt := range cases {
		testname := fmt.Sprintf("%s,input %d, expect %d", tt.provinceGroup, tt.input,tt.want)
		t.Run(testname, func(t *testing.T) {
				provinces := Provinces{tt.provinceGroup:tt.input}
				GetCountProvince(provinces,tt.key)
				expected := tt.want
				
				if provinces[tt.provinceGroup] != expected {
					t.Errorf("expected '%d' but got '%d'", expected, provinces[tt.key])
				}
		})
	}
}


