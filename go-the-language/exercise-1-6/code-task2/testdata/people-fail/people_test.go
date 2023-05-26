package people_test

import (
	"testing"
	"walkintoabar/people"
)

func TestMoodString(t *testing.T) {
	//begin expected
	tests := []struct {
		mood         people.Mood
		expectedStr  string
	}{
		{people.Neutral, "neural"},
		{people.Mood(100), "unknown"}, // Test case for unknown mood
	}
	//end expected
	
	//begin test
	for _, test := range tests {
		result := test.mood.String()
		if result != test.expectedStr {
			t.Errorf("Mood %d: Expected %s, but got %s", test.mood, test.expectedStr, result)
		}
	}
	//end test
}
