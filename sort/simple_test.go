package sort

import (
	"runtime/debug"
	"testing"
)

func TestReplaces(t *testing.T) {
	cases := map[string]sortFunc{
		"stupid":    stupid,
		"bubble":    bubble,
		"shaker":    shaker,
		"selection": selection,
	}
	t.Log(cases)

	for name, f := range cases {
		t.Run(name, func(t *testing.T) {
			f := f
			t.Parallel()
			src := randomSlice(10)
			checkSort(t, src, f)
		})
	}
}

func BenchmarkSimple(b *testing.B) {
	m := 15000

	// benchSort(b, "stupid", stupid, m) // no-no-no, not running this one
	debug.SetGCPercent(-1)
	benchSort(b, "bubble", bubble, m)
	// benchSort(b, "shaker", shaker, m)
	// benchSort(b, "selection", selection, m)
	// benchSort(b, "insertion", insertion, m)
}
