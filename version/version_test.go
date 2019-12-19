package version

import (
	"testing"

	"github.com/Masterminds/semver"
)

// TestInt64 ...
func TestInt64(t *testing.T) {
	type tb struct {
		verstr string
		ret    int64
	}

	for _, tv := range []tb{
		{"v1.2.3", 1020003},
		{"v1.22.333", 1220333},
		{"v1.22.3", 1220003},
		{"v1.22.9999", 1229999},
		{"v1.99.9999", 1999999},
		{"v999.99.9999", 999999999},
		{"v1.22.33333", 0},
		{"v1.122.3", 0},
	} {
		sv := semver.MustParse(tv.verstr)
		n := getInt64(*sv)
		if n != tv.ret {
			t.Errorf("failed, %v, got %v", tv, n)
		}
	}
}
