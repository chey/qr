package code

import "testing"

func TestVersion(t *testing.T) {
	if Version() == "" {
		t.Errorf("missing version")
	}
}
