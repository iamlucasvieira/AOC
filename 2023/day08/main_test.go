package main

import (
	"slices"
	"testing"
)

// ZZZ = (ZZZ, ZZZ)
var mockData = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

var mockData2 = []string{
	"LLR",
	"",
	"AAA = (BBB, BBB)",
	"BBB = (AAA, ZZZ)",
	"ZZZ = (ZZZ, ZZZ)",
}

var mockData3 = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestParseInstructions(t *testing.T) {
	tests := []struct {
		input string
		want  []int
	}{
		{"LRLRL", []int{0, 1, 0, 1, 0}},
		{"RLRLRRR", []int{1, 0, 1, 0, 1, 1, 1}},
		{"L", []int{0}},
		{"R", []int{1}},
		{"", []int{}},
	}

	for _, tt := range tests {
		got, err := parseInstructions(tt.input)
		if err != nil {
			t.Errorf("parseInstructions(%q) = %v, want %v", tt.input, err, tt.want)
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("parseInstructions(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestParseOption(t *testing.T) {
	testCases := []struct {
		input    string
		wantName string
		wantLR   lrPair
	}{
		{"AAA = (BBB, CCC)", "AAA", [2]string{"BBB", "CCC"}},
		{"BBB = (DDD, EEE)", "BBB", [2]string{"DDD", "EEE"}},
		{"CCC = (ZZZ, GGG)", "CCC", [2]string{"ZZZ", "GGG"}},
		{"DDD = (DDD, DDD)", "DDD", [2]string{"DDD", "DDD"}},
		{"EEE = (EEE, EEE)", "EEE", [2]string{"EEE", "EEE"}},
		{"GGG = (GGG, GGG)", "GGG", [2]string{"GGG", "GGG"}},
		{"ZZZ = (ZZZ, ZZZ)", "ZZZ", [2]string{"ZZZ", "ZZZ"}},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			gotName, gotLR, err := parseOption(tc.input)
			if err != nil {
				t.Errorf("parseOption(%q) = %v, want %v", tc.input, err, tc.wantName)
			}
			if gotName != tc.wantName {
				t.Errorf("parseOption(%q) = %v, want %v", tc.input, gotName, tc.wantName)
			}
			if gotLR != tc.wantLR {
				t.Errorf("parseOption(%q) = %v, want %v", tc.input, gotLR, tc.wantLR)
			}
		})
	}
}

func TestParseData(t *testing.T) {
	gotInstruction, gotMap, err := parseData(mockData)
	if err != nil {
		t.Errorf("parseData(%q) = %v, want %v", mockData, err, nil)
	}
	wantInstruction := []int{1, 0}

	if !slices.Equal(gotInstruction, wantInstruction) {
		t.Errorf("parseData(%q) = %v, want %v", mockData, gotInstruction, wantInstruction)
	}

	wantMap := map[string]lrPair{
		"AAA": {"BBB", "CCC"},
		"BBB": {"DDD", "EEE"},
		"CCC": {"ZZZ", "GGG"},
		"DDD": {"DDD", "DDD"},
		"EEE": {"EEE", "EEE"},
		"GGG": {"GGG", "GGG"},
		"ZZZ": {"ZZZ", "ZZZ"},
	}

	for k, v := range wantMap {
		t.Run(k, func(t *testing.T) {
			if gotMap[k][0] != v[0] || gotMap[k][1] != v[1] {
				t.Errorf("parseData(%q) = %v, want %v", mockData, gotMap, wantMap)
			}
		})
	}
}

func TestStepsToZZZ(t *testing.T) {
	testCases := []struct {
		data []string
		want int
	}{
		{mockData, 2},
		{mockData2, 6},
	}

	for _, tc := range testCases {
		t.Run(tc.data[0], func(t *testing.T) {
			instructions, options, err := parseData(tc.data)
			if err != nil {
				t.Errorf("parseData(%q) = %v, want %v", tc.data[0], err, nil)
			}
			got, err := stepsToZZZ(instructions, options)
			if err != nil {
				t.Errorf("stepsToZ(%q) = %v, want %v", tc.data[0], err, tc.want)
			}
			if got != tc.want {
				t.Errorf("stepsToZ(%q) = %v, want %v", tc.data[0], got, tc.want)
			}
		})
	}
}

func TestStepsToZ(t *testing.T) {
	instructions, options, err := parseData(mockData3)
	if err != nil {
		t.Errorf("parseData() = %v, want %v", err, nil)
	}

	var got int
	got, err = stepsToZ(instructions, options)

	if err != nil {
		t.Errorf("stepsToZ() = %v, want %v", err, 6)
	}

	if got != 6 {
		t.Errorf("stepsToZ() = %v, want %v", got, 6)
	}

}
