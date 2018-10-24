package ipsets

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNetset_CIDRS(t *testing.T) {
	f, err := os.Open("./fixtures/fullbogons.tiny.netset")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	netset, err := NewNetset(f)
	assert.NoError(t, err)

	cidrs, err := netset.CIDRS()
	assert.NoError(t, err)

	assert.Equal(t, []string{
		"0.0.0.0/8",
		"2.56.0.0/14",
		"5.133.64.0/18",
		"5.180.0.0/14",
		"5.252.0.0/15",
		"10.0.0.0/8",
		"31.40.192.0/18",
		"37.44.192.0/18",
		"37.221.64.0/18",
		"41.62.0.0/16",
	}, cidrs)
}
