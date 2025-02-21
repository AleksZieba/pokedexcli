package main

import( 
	"testing"
	"fmt"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{ 
			input:		"pikachu BULBASAUR charizard", 
			expected: 	[]string{"pikachu", "bulbasaur", "charizard"}, 
		}, 
		{
			input: 		"I  am  Ironman", 
			expected: 	[]string{"i", "am", "ironman"},
		},
	} 
	for _, c := range cases {
		actual := cleanInput(c.input) 
		if len(c.expected) != len(actual) {
			t.Errorf("Length of Expected and Actual are different")
			fmt.Printf("Test Failed")
		} else {
			fmt.Println("Passed length test")
		}
	
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i] 
			if word != expectedWord {
				t.Errorf("Expected and Actual words do not match") 
				fmt.Printf("Test Failed")
			} else {
				fmt.Println("Passed both tests")
			}
		}
	}
}