package week02_test

import "testing"

func TestArray(t *testing.T) {
	exp := [4]string{"Golang", "Python", "JavaScript", "Ruby"}
	var act [4]string

	// copying the value
	for i, v := range exp {
		act[i] = v
	}

	for i, v := range act {
		if exp[i] != v {
			t.Errorf("Expected Value : %s \n Got value: %s", v, exp[i])
		}
	}
}

// test slicee

func TestSlice(t *testing.T) {
	exp := []string{"Docker", "Kubernetes", "Terraform", "Github Actions"}
	var act []string

	// copying the value
	for _, v := range exp {
		act = append(act, v)
	}

	for i, v := range act {

		// checking length
		if len(exp) != len(act) {
			t.Errorf("Expected length: %d Got length %d", len(exp), len(act))

		} else {
			if exp[i] != v {
				t.Errorf("Expected Value : %s \n Got value: %s", v, exp[i])
			}
		}
	}
}

// test maps

func TestMaps(t *testing.T) {
	exp := map[string]string{
		"joshua@example.com":  "Joshua",
		"jebaraj@example.com": "Jebaraj",
		"jo@example.com":      "Jo",
	}
	act := make(map[string]string)

	// copying the value
	for k, v := range exp {
		act[k] = v
	}

	// checking length
	if len(exp) != len(act) {
		t.Errorf("Expected length: %d Got length %d", len(exp), len(act))

	}

	for k, v := range act {
		value, ok := exp[k]
		if ok {
			if v != value {

				t.Errorf("Expected value %s \n Got value %s \n", value, v)
			}
		}
	}

}
