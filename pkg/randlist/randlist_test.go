package randlist

import (
	"testing"
)

func TestRandomizeStrings(t *testing.T) {
	var olist = []string{
		"Australia",
		"Belgium",
		"Canada",
		"Denmark",
		"Egypt",
		"France",
		"Germany",
		"India",
		"Japan",
		"Malaysia",
		"Netherlands",
		"Pholippines",
		"Russia",
		"Singapore",
		"Turkey",
		"USA",
		"Venezueka",
	}

	var rlist = RandomizeStrings(olist)

	if len(olist) != len(rlist) {
		t.Error("Mismatch between ordered list and randon list sizes")
		return
	}
}
