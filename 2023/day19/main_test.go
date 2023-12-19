package main

import (
	"reflect"
	"testing"
)

var mockData = []string{
	"px{a<2006:qkq,m>2090:A,rfg}",
	"pv{a>1716:R,A}",
	"lnx{m>1548:A,A}",
	"rfg{s<537:gd,x>2440:R,A}",
	"qs{s>3448:A,lnx}",
	"qkq{x<1416:A,crn}",
	"crn{x>2662:A,R}",
	"in{s<1351:px,qqz}",
	"qqz{s>2770:qs,m<1801:hdj,R}",
	"gd{a>3333:R,R}",
	"hdj{m>838:A,pv}",
	"",
	"{x=787,m=2655,a=1222,s=2876}",
	"{x=1679,m=44,a=2067,s=496}",
	"{x=2036,m=264,a=79,s=2244}",
	"{x=2461,m=1339,a=466,s=291}",
	"{x=2127,m=1623,a=2188,s=1013}",
}

func TestNewAction(t *testing.T) {
	tests := []struct {
		input    string
		expected action
	}{
		{
			input: "x>2:5",
			expected: action{
				part:        "x",
				operation:   ">",
				value:       2,
				destination: "5",
			},
		},
		{
			input: "abc<360:def",
			expected: action{
				part:        "abc",
				operation:   "<",
				value:       360,
				destination: "def",
			},
		},
		{
			input: "123>456:789",
			expected: action{
				part:        "123",
				operation:   ">",
				value:       456,
				destination: "789",
			},
		},
		{
			input: "abcdefg",
			expected: action{
				destination: "abcdefg",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			actual, err := newAction(test.input)
			if err != nil {
				t.Errorf("NewAction: %s", err)
			}
			if reflect.DeepEqual(actual, test.expected) == false {
				t.Errorf("NewAction: Expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestNewParts(t *testing.T) {
	testCases := []struct {
		input    string
		expected partsMap
	}{
		{
			input: "{x=1,m=2,a=3,s=4}",
			expected: partsMap(map[string]int{
				"x": 1, "m": 2, "a": 3, "s": 4,
			}),
		},
		{
			input: "{x=10,m=20,a=30,s=40}",
			expected: partsMap(map[string]int{
				"x": 10, "m": 20, "a": 30, "s": 40,
			}),
		},
		{
			input: "{s=100,a=200,m=300,x=400}",
			expected: partsMap(map[string]int{
				"s": 100, "a": 200, "m": 300, "x": 400,
			}),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			actual, err := newParts(testCase.input)
			if err != nil {
				t.Errorf("NewParts: %s", err)
			}
			if reflect.DeepEqual(actual, testCase.expected) == false {
				t.Errorf("NewParts: Expected %v, got %v", testCase.expected, actual)
			}
		})
	}
}

func TestParse(t *testing.T) {
	rules, parts, err := parse(mockData)

	if err != nil {
		t.Errorf("Parse: %s", err)
	}

	if len(rules) != 11 {
		t.Errorf("Parse: Expected 10 rules, got %d", len(rules))
	}

	if len(parts) != 5 {
		t.Errorf("Parse: Expected 4 parts, got %d", len(parts))
	}

	if want := partsMap(map[string]int{"x": 787, "m": 2655, "a": 1222, "s": 2876}); !reflect.DeepEqual(parts[0], want) {
		t.Errorf("Parse: Expected %v, got %v", want, parts[0])
	}

	if _, ok := rules["px"]; !ok {
		t.Errorf("Parse: Expected rule px")
	}

	if len(rules["px"]) != 3 {
		t.Errorf("Parse: Expected 3 actions, got %d", len(rules["px"]))
	}

	if rules["px"][0].part != "a" {
		t.Errorf("Parse: Expected a, got %s", rules["px"][0].part)
	}
}

func TestAction_Evaluate(t *testing.T) {
	p := partsMap(map[string]int{"x": 787, "m": 2655, "a": 1222, "s": 2876})

	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "x>2:5",
			expected: true,
		},
		{
			input:    "x>800:5",
			expected: false,
		},
		{
			input:    "x<800:5",
			expected: true,
		},
		{
			input:    "m>2:5",
			expected: true,
		},
		{
			input:    "m<800:5",
			expected: false,
		},
		{
			input:    "a>3000:5",
			expected: false,
		},
		{
			input:    "a<2000:5",
			expected: true,
		},
		{
			input:    "s>3000:5",
			expected: false,
		},
		{
			input:    "s<3000:5",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			a, err := newAction(test.input)
			if err != nil {
				t.Errorf("NewAction: %s", err)
			}
			actual := a.evaluate(p)
			if actual != test.expected {
				t.Errorf("Action.Evaluate: Expected %v, got %v", test.expected, actual)
			}
		})
	}
}

func TestNumPartsAccepted(t *testing.T) {
	want := 19114
	p, r, err := parse(mockData)
	if err != nil {
		t.Errorf("Parse: %s", err)
	}
	actual := numPartsAccepted(p, r)

	if actual != want {
		t.Errorf("NumPartsAccepted: Expected %d, got %d", want, actual)
	}
}

func TestPart1(t *testing.T) {
	want := 495298
	actual := part1()

	if actual != want {
		t.Errorf("Part1: Expected %d, got %d", want, actual)
	}
}

func TestRanges_update(t *testing.T) {
	r := ranges{min: 0, max: 10}

	testCases := []struct {
		input    string
		expected ranges
	}{
		{
			input:    "x>2:5",
			expected: ranges{min: 3, max: 10},
		},
		{
			input:    "x>800:5",
			expected: ranges{min: 801, max: 10},
		},
		{
			input:    "x<800:5",
			expected: ranges{min: 0, max: 10},
		},
		{
			input:    "x<2:5",
			expected: ranges{min: 0, max: 1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.input, func(t *testing.T) {
			a, err := newAction(testCase.input)
			if err != nil {
				t.Errorf("NewAction: %s", err)
			}
			got := r.update(a)
			if got != testCase.expected {
				t.Errorf("Ranges.update: Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestRanges_Valid(t *testing.T) {
	testCases := []struct {
		id       string
		input    ranges
		expected bool
	}{
		{
			id:       "min < max",
			input:    ranges{min: 0, max: 10},
			expected: true,
		},
		{
			id:       "min == max",
			input:    ranges{min: 10, max: 0},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			got := testCase.input.isValid()
			if got != testCase.expected {
				t.Errorf("Ranges.isValid: Expected %v, got %v", testCase.expected, got)
			}
		})
	}
}

func TestRanges_Len(t *testing.T) {
	testCases := []struct {
		id       string
		input    ranges
		expected int
	}{
		{
			id:       "min < max",
			input:    ranges{min: 0, max: 10},
			expected: 11,
		},
		{
			id:       "min == max",
			input:    ranges{min: 10, max: 10},
			expected: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.id, func(t *testing.T) {
			got := testCase.input.len()
			if got != testCase.expected {
				t.Errorf("Ranges.len: Expected %d, got %d", testCase.expected, got)
			}
		})
	}
}

func TestNewPartsMapRanges(t *testing.T) {
	want := partsMapRange{
		"x": {min: 0, max: 10},
		"m": {min: 0, max: 10},
		"a": {min: 0, max: 10},
		"s": {min: 0, max: 10},
	}
	actual := newPartsMapRanges(0, 10)

	if !reflect.DeepEqual(actual, want) {
		t.Errorf("NewPartsMapRanges: Expected %v, got %v", want, actual)
	}
}

func TestRanges_Split(t *testing.T) {
	r := ranges{min: 5, max: 10}
	testCases := []struct {
		action             string
		wantSatisfies      ranges
		wantDoesNotSatisfy ranges
	}{
		{
			action:             "x>1:5",
			wantSatisfies:      ranges{min: 5, max: 10},
			wantDoesNotSatisfy: ranges{},
		},
		{
			action:             "x>7:5",
			wantSatisfies:      ranges{min: 8, max: 10},
			wantDoesNotSatisfy: ranges{min: 5, max: 7},
		},
		{
			action:             "x>15:5",
			wantSatisfies:      ranges{},
			wantDoesNotSatisfy: ranges{min: 5, max: 10},
		},

		{
			action:             "x<1:5",
			wantSatisfies:      ranges{},
			wantDoesNotSatisfy: ranges{min: 5, max: 10},
		},
		{
			action:             "x<7:5",
			wantSatisfies:      ranges{min: 5, max: 6},
			wantDoesNotSatisfy: ranges{min: 7, max: 10},
		},
		{
			action:             "x<15:5",
			wantSatisfies:      ranges{min: 5, max: 10},
			wantDoesNotSatisfy: ranges{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.action, func(t *testing.T) {
			a, err := newAction(testCase.action)
			if err != nil {
				t.Errorf("NewAction: %s", err)
			}
			gotSatisfies, gotDoesNotSatisfy := r.split(a)
			if !reflect.DeepEqual(gotSatisfies, testCase.wantSatisfies) {
				t.Errorf("Ranges.split: Expected %v, got %v", testCase.wantSatisfies, gotSatisfies)
			}
			if !reflect.DeepEqual(gotDoesNotSatisfy, testCase.wantDoesNotSatisfy) {
				t.Errorf("Ranges.split: Expected %v, got %v", testCase.wantDoesNotSatisfy, gotDoesNotSatisfy)
			}
		})
	}
}

func TestRanges_IsEmpty(t *testing.T) {
	testCases := []struct {
		description string
		input       ranges
		want        bool
	}{
		{
			description: "ranges{}",
			input:       ranges{},
			want:        true,
		},
		{
			description: "ranges{min: 0, max: 1}",
			input:       ranges{min: 0, max: 1},
			want:        false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := tc.input.isEmpty()
			if got != tc.want {
				t.Errorf("Ranges.isEmpty: Expected %v, got %v", tc.want, got)
			}
		})
	}
}

func TestPartsMapRanges_Update(t *testing.T) {
	p := newPartsMapRanges(0, 10)

	testCases := []struct {
		key   string
		value ranges
	}{
		{
			key:   "x",
			value: ranges{min: 5, max: 10},
		},
		{
			key:   "m",
			value: ranges{min: 5, max: 10},
		},
		{
			key:   "a",
			value: ranges{min: 5, max: 10},
		},
		{
			key:   "s",
			value: ranges{min: 5, max: 10},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.key, func(t *testing.T) {
			got := p.update(tc.key, tc.value)
			if !reflect.DeepEqual(got[tc.key], tc.value) {
				t.Errorf("PartsMapRanges.update: Expected %v, got %v", tc.value, got[tc.key])
			}
		})
	}
}

func TestSumPossibleAcceptablePartsSimple(t *testing.T) {
	testCases := []struct {
		action []string
		want   int
	}{
		{
			action: []string{"in{x>9:A,R}"},
			want:   1 * 10 * 10 * 10,
		},
		{
			action: []string{"in{x<10:R,m<10:R,a<10:R,s<10:R,A}"},
			want:   1,
		},
		{
			action: []string{"in{x>100:A,R"},
			want:   0,
		},
		{
			action: []string{"in{x<100:A,R"},
			want:   10 * 10 * 10 * 10,
		},
		{
			action: []string{
				"in{x>9:2,R}",
				"2{a>9:3,R}",
				"3{m<2:A,R}",
			},
			want: 1 * 1 * 1 * 10,
		},
		{
			action: []string{
				"in{x>9:2,R}",
				"2{a>9:3,4}",
				"3{m<2:A,R}",
				"4{m>9:A,R}",
			},
			want: 1*1*1*10 + 1*9*1*10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.action[0], func(t *testing.T) {
			r, _, err := parse(tc.action)
			if err != nil {
				t.Errorf("Parse: %s", err)
			}
			p := newPartsMapRanges(1, 10)
			got := sumPossibleAcceptableParts("in", 0, r, p)
			if got != tc.want {
				t.Errorf("SumPossibleAcceptableParts: Expected %d, got %d", tc.want, got)
			}
		})
	}
}
func TestSumPossibleAcceptableParts(t *testing.T) {
	r, _, err := parse(mockData)

	if err != nil {
		t.Errorf("Parse: %s", err)
	}

	want := 167409079868000

	p := newPartsMapRanges(1, 4000)

	start := "in"
	actual := sumPossibleAcceptableParts(start, 0, r, p)

	if actual != want {
		t.Errorf("SumPossibleAcceptableParts: Expected %d, got %d", want, actual)
	}

}

func TestPart2(t *testing.T) {
	want := 132186256794011
	actual := part2()

	if actual != want {
		t.Errorf("Part2: Expected %d, got %d", want, actual)
	}
}
