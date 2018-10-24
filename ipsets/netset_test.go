package ipsets

import (
	"os"
	"testing"
)

func TestNetset_CIDRS(t *testing.T) {
	f, err := os.Open("./fixtures/fullbogons.netset")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
}
