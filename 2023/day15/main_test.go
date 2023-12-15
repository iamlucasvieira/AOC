package main

import "testing"

var mockData = []string{
	"rn=1,cm-,qp=3,cm=2,qp-,pc",
	"=4,ot=9,ab=5,pc-,pc=6,ot=7",
}

func TestParse(t *testing.T) {
	data := parse(mockData)

	if len(data) != 11 {
		t.Errorf("Expected 11, got %d", len(data))
	}

	if data[0] != "rn=1" {
		t.Errorf("Expected rn=1, got %s", data[0])
	}
}

func TestHash(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"rn=1", 30},
		{"cm-", 253},
		{"qp=3", 97},
		{"cm=2", 47},
		{"qp-", 14},
		{"pc=4", 180},
		{"ot=9", 9},
		{"ab=5", 197},
		{"pc-", 48},
		{"pc=6", 214},
		{"ot=7", 231},
		{"HASH", 52},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			actual := hash(tc.input)
			if actual != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func TestHashSum(t *testing.T) {
	data := parse(mockData)
	actual := hashSum(data)
	expected := 1320

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestParseLenses(t *testing.T) {
	lenses := parseLenses(mockData)

	if len(lenses) != 11 {
		t.Errorf("Expected 11, got %d", len(lenses))
	}

	if lenses[0].label != "rn" {
		t.Errorf("Expected rn, got %s", lenses[0].label)
	}

	if lenses[0].focal != 1 {
		t.Errorf("Expected 1, got %d", lenses[0].focal)
	}

	if lenses[0].operation != "=" {
		t.Errorf("Expected =, got %s", lenses[0].operation)
	}
}

func TestAddLens(t *testing.T) {
	b := make(boxes)

	l1 := lens{"rn", 1, "="}
	l2 := lens{"cm", 2, "="}
	l3 := lens{"qp", 3, "="}

	addLens(b, l1)

	if b[0][0] != l1 {
		t.Errorf("Expected %v, got %v", l1, b[0][0])
	}

	addLens(b, l2)

	if b[0][1] != l2 {
		t.Errorf("Expected %v, got %v", l2, b[0][1])
	}

	addLens(b, l3)

	if b[1][0] != l3 {
		t.Errorf("Expected %v, got %v", l3, b[0][2])
	}
}

func TestRemoveLens(t *testing.T) {

	b := make(boxes)
	l1 := lens{"rn", 1, "="}
	l2 := lens{"cm", 2, "="}
	l3 := lens{"qp", 3, "="}

	addLens(b, l1)
	addLens(b, l2)
	addLens(b, l3)

	l4 := lens{"rn", -1, "-"}

	removeLens(b, l4)

	if len(b[0]) != 1 {
		t.Errorf("Expected 1, got %d", len(b[0]))
	}

	if b[0][0] != l2 {
		t.Errorf("Expected %v, got %v", l2, b[0][0])
	}
}

func TestProcessLenses(t *testing.T) {
	data := parseLenses(mockData)

	b := make(boxes)

	processLenses(b, data)
	t.Run("Length of box 0", func(t *testing.T) {
		if len(b[0]) != 2 {
			t.Errorf("Expected 2, got %d", len(b[0]))
		}
	})

	t.Run("Length of box 3", func(t *testing.T) {
		if len(b[3]) != 3 {
			t.Errorf("Expected 3, got %d", len(b[3]))
		}
	})

	t.Run("Items box 0", func(t *testing.T) {
		if b[0][0].focal != 1 {
			t.Errorf("Expected %v, got %v", 1, b[0][0].focal)
		}
		if b[0][1].focal != 2 {
			t.Errorf("Expected %v, got %v", 2, b[0][1].focal)
		}
	})

	t.Run("Items box 3", func(t *testing.T) {
		if b[3][0].focal != 7 {
			t.Errorf("Expected %v, got %v", 7, b[3][0].focal)
		}
		if b[3][1].focal != 5 {
			t.Errorf("Expected %v, got %v", 5, b[3][1].focal)
		}
		if b[3][2].focal != 6 {
			t.Errorf("Expected %v, got %v", 6, b[3][2].focal)
		}
	})
}

func TestScore(t *testing.T) {
	data := parseLenses(mockData)

	b := make(boxes)

	processLenses(b, data)

	actual := score(b)
	expected := 145

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
