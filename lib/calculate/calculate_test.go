package calculate

import "testing"

func TestAddition(t *testing.T) {
	if Addition(1, 2) != 3 {
		t.Error("Expected is 3")
	}
}
